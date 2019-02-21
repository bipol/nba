package nba

import (
	"net/http"
)

//APIOption is an option that affects the request
//The idea here is that the stats.nba api has similar arguments
//for each request -- this allows us to mutate them in a
//genericish way
type APIOption = func(*http.Request)

//Players Endpoint Parameters
//NOTE:  I think that all the players endpoints will use the same as these below
//so we might just be able to dedupe and go from there?
//also, below is the default for stats.nba
/**
College:
Conference:
Country:
DateFrom:
DateTo:
Division:
DraftPick:
DraftYear:
GameScope:
GameSegment:
Height:
LastNGames: 0
LeagueID: 00
Location:
MeasureType: Base
Month: 0
OpponentTeamID: 0
Outcome:
PORound: 0
PaceAdjust: N
PerMode: PerGame
Period: 0
PlayerExperience:
PlayerPosition:
PlusMinus: N
Rank: N
Season: 2018-19
SeasonSegment:
SeasonType: Regular Season
ShotClockRange:
StarterBench:
TeamID: 0
TwoWay: 0
VsConference:
VsDivision:
Weight:
**/

// WithMode set the mode argument
func WithMode(mode string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("PerMode", mode)
		req.URL.RawQuery = q.Encode()
	}
}

// WithLeagueID sets the league argument
func WithLeagueID(id string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("LeagueID", id)
		req.URL.RawQuery = q.Encode()
	}
}

// WithScope sets the scope argument
func WithScope(scope string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("Scope", scope)
		req.URL.RawQuery = q.Encode()
	}

}

// DuringSeason sets the season argument
func DuringSeason(season string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("Season", season)
		req.URL.RawQuery = q.Encode()
	}
}

// WithSeasonType sets the season type argument
func WithSeasonType(seasonType string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("SeasonType", seasonType)
		req.URL.RawQuery = q.Encode()
	}
}

// WithStatCategory sets the stat type argument
func WithStatCategory(stats string) APIOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add("StatsCategory", stats)
		req.URL.RawQuery = q.Encode()
	}
}
