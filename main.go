package main

import (
	"fmt"
	"os"

	"github.com/esimonov/my-chesscom-stats/client"
	"github.com/esimonov/my-chesscom-stats/export"
)

var username string

func init() {
	username = os.Getenv("CHESS_COM_USERNAME")
	if len(username) == 0 {
		panic("CHESS_COM_USERNAME not set")
	}
}

func main() {
	c := client.NewClient(username)
	outFilename := fmt.Sprintf("%s_chess_stats.csv", username)

	exporters := []export.Exporter{export.NewCSVExporter(c)}
	for _, e := range exporters {
		if err := e.AllGames(outFilename); err != nil {
			panic(err)
		}
	}
}
