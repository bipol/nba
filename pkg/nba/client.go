package nba

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"

	log "github.com/sirupsen/logrus"
)

const (
	//LeagueLeadersEndpoint represents the stats.nba endpoint for league leaders
	LeagueLeadersEndpoint = "https://stats.nba.com/stats/leagueLeaders"
	//PlayersEndpoint exposes all of the players in the nba, with general stats
	//Note that this has a lot of filters -- see options.go
	PlayersEndpoint = "https://stats.nba.com/stats/leaguedashplayerstats"
	//PlayersDefenseDashboardEndpoint
	PlayersDefenseDashboardEndpoint = "https://stats.nba.com/stats/leaguedashptdefend"

	//js data endpoints
	//these endpoints don't use the standard wrapper that the stats endpoints do
	//and will require more work
	//AdvancedLeadersEndpoint represents the stats.nba endpoint for advanced leaders
	AdvancedLeadersEndpoint = "https://stats.nba.com/js/data/widgets/advanced_leaders.json"
	//HustleLeadersEndpoint
	HustleLeadersEndpoint = "https://stats.nba.com/js/data/widgets/hustle_leaders.json"
)

//API exposes stats.nba.com endpoints
//go:generate counterfeiter . API
type API interface {
	GetLeagueLeaders(options ...APIOption) ([]LeagueLeaderRow, error)
}

//Client contains all the needed information to query the NBA API
//TODO: Change to interfaces
type Client struct {
	client http.Client
}

//NewClient will an NBA Client
//TODO: Pass in a logger here
func NewClient(client http.Client) Client {
	return Client{
		client: client,
	}
}

//GetLeagueLeaders returns the current league leaders
func (c *Client) sendRequest(url string, required map[string]string, options ...APIOption) (ResponseEnvelope, error) {
	var set ResponseEnvelope
	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
		GotFirstResponseByte: func() {
			fmt.Printf("Got First Byte\n")
		},
		Got100Continue: func() {
			fmt.Printf("Got 100 Continue\n")
		},
		Wait100Continue: func() {
			fmt.Printf("Wait 100 Continue\n")
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return set, err
	}
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	handleOptions(req, options, required)
	addHeaders(req)
	log.Debugf("request %v", req)
	resp, err := c.client.Do(req)
	if err != nil {
		return set, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return set, err
	}

	err = json.Unmarshal(body, &set)
	if err != nil {
		return set, err
	}

	return set, err
}

//GetLeagueLeaders returns the current league leaders
func (c *Client) GetLeagueLeaders(options ...APIOption) ([]LeagueLeaderRow, error) {
	var rows []LeagueLeaderRow
	envelope, err := c.sendRequest(LeagueLeadersEndpoint, LeagueLeaderRequiredFields, options...)
	if err != nil {
		return rows, err
	}
	for _, row := range envelope.ResultSet.RowSet {
		var llRow LeagueLeaderRow
		err = llRow.UnmarshalRawMessage(row)
		if err == nil {
			rows = append(rows, llRow)
		}
	}
	return rows, err
}

//GetPlayerStats will return player stats
//TODO: Figure out the best way to return these stats...
//do we delay the actual processing and defer that to the user?
//that seems like a mostly useless thing to do
func (c *Client) GetPlayerStats(options ...APIOption) ([][]json.RawMessage, error) {
	var rows [][]json.RawMessage
	envelope, err := c.sendRequest(PlayersEndpoint, PlayerRequiredFields, options...)
	if err != nil {
		return rows, err
	}
	//TODO: How to handle multi result sets?
	if len(envelope.ResultSets) > 0 {
		for _, row := range envelope.ResultSets {
			for _, result := range row.RowSet {
				rows = append(rows, result)
			}
		}
	}
	return rows, err
}

func addHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	//Fake User Agent??
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
}

func handleOptions(req *http.Request, options []APIOption, required map[string]string) {
	for _, option := range options {
		option(req)
	}
	defaultArguments(req, required)
}

func defaultArguments(req *http.Request, required map[string]string) {
	q := req.URL.Query()
	for field, val := range required {
		if _, ok := q[field]; !ok {
			q.Add(field, val)
		}
	}
	req.URL.RawQuery = q.Encode()
}
