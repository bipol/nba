package nba

import (
	"net/http"
)

//APIOption is an option that affects the request
type APIOption = func(*http.Request)

/**
LeagueID: 00
PerMode: PerGame
Scope: S
Season: 2018-19
SeasonType: Regular Season
StatCategory: PTS
**/

//WithPerGameMode will query for the request on a per game basis
func WithPerGameMode(req *http.Request) {
	q := req.URL.Query()
	q.Add("PerMode", "PerGame")
	req.URL.RawQuery = q.Encode()
}
