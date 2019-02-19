package nba

import (
	"net/http"
	"log"
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
func (c *Client) GetLeagueLeaders(options []APIOption) error {
	req, err := http.NewRequest("GET", "https://stats.nba.com/stats/leagueLeaders", nil)
	if err != nil {
		return err
	}
	handleOptions(req, options)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
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
