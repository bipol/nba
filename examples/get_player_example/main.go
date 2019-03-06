package main

import (
	"net/http"

	"github.com/bipol/nba/pkg/nba"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	httpClient := http.Client{}
	client := nba.NewClient(httpClient)
	res, err := client.GetPlayerStats(
		nba.DuringSeason("2018-19"),
		nba.ForPlayerPosition(nba.Forward),
		nba.WithTeamID(nba.AtlantaHawks),
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(res)
}
