package main

import (
	"net/http"

	"github.com/bipol/nba/pkg/nba"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	httpClient := http.Client{}
	client := nba.NewClient(httpClient)
	set, err := client.GetLeagueLeaders()
	if err != nil {
		panic(err)
	}
	spew.Dump(set)
}
