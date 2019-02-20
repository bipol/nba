package nba

import (
	"net/http"
)

//APIOption is an option that affects the request
//The idea here is that the stats.nba api has similar arguments
//for each request -- this allows us to mutate them in a
//genericish way
type APIOption = func(*http.Request)

//League Leaders Parameters
//also, below is the default for stats.nba
/**
LeagueID: 00
PerMode: PerGame
Scope: S
Season: 2018-19
SeasonType: Regular Season
StatCategory: PTS
**/

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

//WithPerGameMode will query for the request on a per game basis
func WithPerGameMode(req *http.Request) {
	q := req.URL.Query()
	q.Add("PerMode", "PerGame")
	req.URL.RawQuery = q.Encode()
}
