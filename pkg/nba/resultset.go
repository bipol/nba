package nba

import (
	"encoding/json"
	"errors"
)

//ResultSet represents the results we get back from the nba api
//I'm stupidly assuming that all the other stat.nba endpoints have a similar envelope
type ResultSet struct {
	Name    string          `json:"name"`
	Headers []string        `json:"headers"`
	RowSet  json.RawMessage `json:"rowSet"`
}

//ResponseEnvelope is the wrapper around the result set that the nba api
//TODO: maybe i don't expose this dude to the world and just use it internally
//to help with the multi response bodies?
type ResponseEnvelope struct {
	ResultSet  ResultSet   `json:"resultSet"`
	ResultSets []ResultSet `json:"resultSets"`
}

//LeagueLeaderRow contains the info needed to destruct the league leader
//api response
type LeagueLeaderRow struct {
	PlayerID json.Number `json:"PLAYER_ID"`
	Rank     json.Number `json:"RANK"`
	Player   string      `json:"PLAYER"`
	Team     string      `json:"TEAM"`
	Gp       json.Number `json:"GP"`
	Min      json.Number `json:"MIN"`
	Fgm      json.Number `json:"FGM"`
	Fga      json.Number `json:"FGA"`
	FgPCT    json.Number `json:"FG_PCT"`
	Fg3m     json.Number `json:"FD_3M"`
	Fg3a     json.Number `json:"FG_3A"`
	Fg3PCT   json.Number `json:"FG_3PCT"`
	Ftm      json.Number `json:"FTM"`
	Fta      json.Number `json:"FTA"`
	FtPCT    json.Number `json:"FT_PCT"`
	Oreb     json.Number `json:"OREB"`
	Dreb     json.Number `json:"DREB"`
	Reb      json.Number `json:"REB"`
	Ast      json.Number `json:"AST"`
	Stl      json.Number `json:"STL"`
	Blk      json.Number `json:"BLK"`
	Tov      json.Number `json:"TOV"`
	Pts      json.Number `json:"PTS"`
	Eff      json.Number `json:"EFF"`
}

//PlayerRequiredFields outline the required fields for the query
//TODO convert this stuff to consts?
var PlayerRequiredFields = map[string]string{
	"DateFrom":         "",
	"DateTo":           "",
	"GameScope":        "",
	"GameSegment":      "",
	"LastNGames":       "0",
	"Location":         "",
	"MeasureType":      "Base",
	"Month":            "0",
	"OpponentTeamID":   "0",
	"Outcome":          "",
	"PaceAdjust":       "N",
	"PerMode":          "PerGame",
	"Period":           "0",
	"PlayerExperience": "",
	"PlayerPosition":   "",
	"PlusMinus":        "N",
	"Rank":             "N",
	"Season":           "2018-19",
	"SeasonSegment":    "",
	"SeasonType":       "Regular Season",
	"StarterBench":     "",
	"VsConference":     "",
	"VsDivision":       "",
}

//LeagueLeaderRequiredFields outline the required fields for the query
var LeagueLeaderRequiredFields = map[string]string{
	"PerMode":      "PerGame",
	"LeagueID":     "00",
	"Scope":        "S",
	"Season":       "2018-19",
	"SeasonType":   "Regular Season",
	"StatCategory": "PTS",
}

var (
	//ErrNotEnoughCols is thrown if we have a raw message from the nba
	//response that has more or less columns than we expect
	ErrNotEnoughCols = errors.New("not enough rows to process")
)

//UnmarshalJSON will take the response and convert it into a league
//leader row.
func (l *LeagueLeaderRow) UnmarshalJSON(bytes []byte) error {
	var (
		err        error
		rawMessage []json.RawMessage
	)

	err = json.Unmarshal(bytes, &rawMessage)
	if err != nil {
		return err
	}

	if len(rawMessage) != 24 {
		return ErrNotEnoughCols
	}

	err = json.Unmarshal(rawMessage[0], &l.PlayerID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[1], &l.Rank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[2], &l.Player)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[3], &l.Team)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[4], &l.Gp)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[5], &l.Min)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[6], &l.Fgm)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[7], &l.Fga)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[8], &l.FgPCT)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[9], &l.Fg3m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[10], &l.Fg3a)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[11], &l.Fg3PCT)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[12], &l.Ftm)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[13], &l.Fta)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[14], &l.FtPCT)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[15], &l.Oreb)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[16], &l.Dreb)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[17], &l.Reb)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[18], &l.Ast)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[19], &l.Stl)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[20], &l.Blk)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[21], &l.Tov)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[22], &l.Pts)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rawMessage[23], &l.Eff)
	return err
}
