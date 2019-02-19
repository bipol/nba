package nba

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//LeagueLeadersEndpoint represents the stats nba endpoint for league leaders
	LeagueLeadersEndpoint = "https://stats.nba.com/stats/leagueLeaders"
)

//API exposes stats.nba.com endpoints
type API interface {
	GetLeagueLeaders()
}

type Doer interface {
}

type Logger interface {
}

//Client contains all the needed information to query the NBA API
//TODO: Change to interfaces
type Client struct {
	client http.Client
	logger log.Logger
}

//NewClient will an NBA Client
func NewClient(client http.Client) Client {
	return Client{
		client: client,
	}
}

//GetLeagueLeaders returns the current league leaders
func (c *Client) GetLeagueLeaders(options ...APIOption) (ResultSet, error) {
	var set ResultSet
	req, err := http.NewRequest("GET", LeagueLeadersEndpoint, nil)
	if err != nil {
		return set, err
	}
	handleOptions(req, options)
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

func handleOptions(req *http.Request, options []APIOption) {
	for _, option := range options {
		option(req)
	}
	defaultArguments(req)
}

func defaultArguments(req *http.Request) {
	q := req.URL.Query()
	if _, ok := q["PerMode"]; !ok {
		q.Add("PerMode", "PerGame")
	}
	q.Add("LeagueID", "00")
	q.Add("Scope", "S")
	q.Add("Season", "2018-19")
	q.Add("SeasonType", "Regular Season")
	q.Add("StatsCategory", "PTS")
	req.URL.RawQuery = q.Encode()
}
