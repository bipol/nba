package nba

//ResultSet represents the results we get back from the nba api
type ResultSet struct {
	Name    string   `json:"name"`
	Headers []string `json:"headers"`
	RowSet  []string `json:"rowSet"`
}
