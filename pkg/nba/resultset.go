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
	Results    json.RawMessage `json:"results"`
	ResultSet  ResultSet       `json:"resultSet"`
	ResultSets []ResultSet     `json:"resultSets"`
}

//TODO: Expand struct
type DailyGameLineupRow struct {
	Game   string `json:"Game"`
	GameID string `json:"GameID"`
}

//PlayerRow contains the info needed to destruct the player
//api response
type PlayerRow struct {
	PlayerID          json.Number `json:"PLAYER_ID"`
	PlayerName        string      `json:"PLAYER_NAME"`
	TeamID            json.Number `json:"TEAM_ID"`
	TeamAbbreviation  string      `json:"TEAM_ABBREVIATION"`
	Age               json.Number `json:"AGE"`
	Gp                json.Number `json:"GP"`
	W                 json.Number `json:"W"`
	L                 json.Number `json:"L"`
	WPct              json.Number `json:"W_PCT"`
	Min               json.Number `json:"MIN"`
	Fgm               json.Number `json:"FGM"`
	Fga               json.Number `json:"FGA"`
	FgPct             json.Number `json:"FG_PCT"`
	Fg3m              json.Number `json:"FG3M"`
	Fg3a              json.Number `json:"FG3A"`
	Fg3Pct            json.Number `json:"FG3_PCT"`
	Ftm               json.Number `json:"FTM"`
	Fta               json.Number `json:"FTA"`
	FtPct             json.Number `json:"FT_PCT"`
	Oreb              json.Number `json:"OREB"`
	Dreb              json.Number `json:"DREB"`
	Reb               json.Number `json:"REB"`
	Ast               json.Number `json:"AST"`
	Tov               json.Number `json:"TOV"`
	Stl               json.Number `json:"STL"`
	Blk               json.Number `json:"BLK"`
	Blka              json.Number `json:"BLKA"`
	Pf                json.Number `json:"PF"`
	Pfd               json.Number `json:"PFD"`
	Pts               json.Number `json:"PTS"`
	PlusMinus         json.Number `json:"PLUS_MINUS"`
	NbaFantasyPts     json.Number `json:"NBA_FANTASY_PTS"`
	Dd2               json.Number `json:"DD2"`
	Td3               json.Number `json:"TD3"`
	GpRank            json.Number `json:"GP_RANK"`
	WRank             json.Number `json:"W_RANK"`
	LRank             json.Number `json:"L_RANK"`
	WPctRank          json.Number `json:"W_PCT_RANK"`
	MinRank           json.Number `json:"MIN_RANK"`
	FgmRank           json.Number `json:"FGM_RANK"`
	FgaRank           json.Number `json:"FGA_RANK"`
	FgPctRank         json.Number `json:"FG_PCT_RANK"`
	Fg3mRank          json.Number `json:"FG3M_RANK"`
	Fg3aRank          json.Number `json:"FG3A_RANK"`
	Fg3PctRank        json.Number `json:"FG3_PCT_RANK"`
	FtmRank           json.Number `json:"FTM_RANK"`
	FtaRank           json.Number `json:"FTA_RANK"`
	FtPctRank         json.Number `json:"FT_PCT_RANK"`
	OrebRank          json.Number `json:"OREB_RANK"`
	DrebRank          json.Number `json:"DREB_RANK"`
	RebRank           json.Number `json:"REB_RANK"`
	AstRank           json.Number `json:"AST_RANK"`
	TovRank           json.Number `json:"TOV_RANK"`
	StlRank           json.Number `json:"STL_RANK"`
	BlkRank           json.Number `json:"BLK_RANK"`
	BlkaRank          json.Number `json:"BLKA_RANK"`
	PfRank            json.Number `json:"PF_RANK"`
	PfdRank           json.Number `json:"PFD_RANK"`
	PtsRank           json.Number `json:"PTS_RANK"`
	PlusMinusRank     json.Number `json:"PLUS_MINUS_RANK"`
	NbaFantasyPtsRank json.Number `json:"NBA_FANTASY_PTS_RANK"`
	Dd2Rank           json.Number `json:"DD2_RANK"`
	Td3Rank           json.Number `json:"TD3_RANK"`
	Cfid              json.Number `json:"CFID"`
	Cfparams          string      `json:"CFPARAMS"`
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

//UnmarshalJSON will take the response and convert it into a player
//row.
func (l *PlayerRow) UnmarshalJSON(bytes []byte) error {
	var (
		err        error
		rawMessage []json.RawMessage
	)

	err = json.Unmarshal(bytes, &rawMessage)
	if err != nil {
		return err
	}

	if len(rawMessage) != 65 {
		return ErrNotEnoughCols
	}

	err = json.Unmarshal(rawMessage[0], &l.PlayerID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[1], &l.PlayerName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[2], &l.TeamID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[3], &l.TeamAbbreviation)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[4], &l.Age)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[5], &l.Gp)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[6], &l.W)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[7], &l.L)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[8], &l.WPct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[9], &l.Min)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[10], &l.Fgm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[11], &l.Fga)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[12], &l.FgPct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[13], &l.Fg3m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[14], &l.Fg3a)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[15], &l.Fg3Pct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[16], &l.Ftm)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[17], &l.Fta)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[18], &l.FtPct)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[19], &l.Oreb)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[20], &l.Dreb)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[21], &l.Reb)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[22], &l.Ast)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[23], &l.Tov)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[24], &l.Stl)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[25], &l.Blk)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[26], &l.Blka)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[27], &l.Pf)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[28], &l.Pfd)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[29], &l.Pts)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[30], &l.PlusMinus)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[31], &l.NbaFantasyPts)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[32], &l.Dd2)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[33], &l.Td3)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[34], &l.GpRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[35], &l.WRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[36], &l.LRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[37], &l.WPctRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[38], &l.MinRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[39], &l.FgmRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[40], &l.FgaRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[41], &l.FgPctRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[42], &l.Fg3mRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[43], &l.Fg3aRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[44], &l.Fg3PctRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[45], &l.FtmRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[46], &l.FtaRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[47], &l.FtPctRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[48], &l.OrebRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[49], &l.DrebRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[50], &l.RebRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[51], &l.AstRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[52], &l.TovRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[53], &l.StlRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[54], &l.BlkRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[55], &l.BlkaRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[56], &l.PfRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[57], &l.PfdRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[58], &l.PtsRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[59], &l.PlusMinusRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[60], &l.NbaFantasyPtsRank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[61], &l.Dd2Rank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[62], &l.Td3Rank)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[63], &l.Cfid)
	if err != nil {
		return err
	}
	err = json.Unmarshal(rawMessage[64], &l.Cfparams)
	return err
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

//NoRequiredFields outlines no required fields
var NoRequiredFields = map[string]string{}

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
