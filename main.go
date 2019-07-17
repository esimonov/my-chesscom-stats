package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/esimonov/my-chesscom-stats/client"
	"github.com/esimonov/my-chesscom-stats/model"
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

	gameArchivesURLList, err := c.GetMonthlyGameArchivesURLList()
	if err != nil {
		panic(err)
	}

	outFilename := fmt.Sprintf("%s chess stats.csv", username)
	out, err := os.OpenFile(outFilename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Panicln("Cannot open CSV file: ", err)
	}

	defer func() {
		log.Println("Closing CSV file..")
		if err := out.Close(); err != nil {
			log.Println("Error closing CSV file:", err)
		} else {
			log.Println("CSV file successfully closed")
		}
	}()

	w := csv.NewWriter(out)
	if err = w.Write(model.GetGameExportCSVHeader()); err != nil {
		log.Panicln("Cannot write CSV header: ", err)
	}

	for _, gamesByMonthURL := range gameArchivesURLList {
		urlLen := len(gamesByMonthURL)
		log.Printf("Requesting games for %s\n", gamesByMonthURL[urlLen-7:])

		games, err := c.GetMonthlyGameArchive(gamesByMonthURL)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, g := range games {
			if err = w.Write(g.ToGameExport(username).ToStringSlice()); err != nil {
				log.Println(err)
			}
		}
	}

	w.Flush()
	if err = w.Error(); err != nil {
		log.Panicln(err)
	}
}
