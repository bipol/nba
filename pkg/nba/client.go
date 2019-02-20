package nba

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	logger log.Logger
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
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return set, err
	}
	handleOptions(req, options, required)
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
func (c *Client) GetPlayerStats(options ...APIOption) ([]LeagueLeaderRow, error) {
	var rows []LeagueLeaderRow
	envelope, err := c.sendRequest(PlayersEndpoint, PlayerRequiredFields, options...)
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
