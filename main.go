package main

import (
	"fmt"
	"os"

	"github.com/esimonov/my-chesscom-stats/exporter"
	csv "github.com/esimonov/my-chesscom-stats/exporter/csv"
)

var username string

func init() {
	username = os.Getenv("CHESS_COM_USERNAME")
	if len(username) == 0 {
		panic("CHESS_COM_USERNAME not set")
	}
}

func main() {
	dispatcher := exporter.NewDispatcher(username)

	csvFilename := fmt.Sprintf("%s_chess_stats.csv", username)
	csvExporter, err := csv.NewExporter(csvFilename, username)
	if err != nil {
		panic(err)
	}

	exporters := []exporter.Exporter{csvExporter}

	if err = dispatcher.AllGames(exporters); err != nil {
		panic(err)
	}

	fmt.Println("Success!")
}
