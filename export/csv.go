package export

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/esimonov/my-chesscom-stats/client"
	"github.com/esimonov/my-chesscom-stats/model"
)

const csvExtension = ".csv"

// CSVExporter exports to CSV.
type CSVExporter struct {
	Client *client.Client
}

// NewCSVExporter creates new CSV exporter.
func NewCSVExporter(c *client.Client) *CSVExporter {
	return &CSVExporter{
		Client: c,
	}
}

// AllGames exports all games for provided username to csv sheet.
// Filename is a name of csv sheet. If extention is not provided, it will be added automatically.
func (e *CSVExporter) AllGames(filename string) error {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		return fmt.Errorf("Invalid filename %s", filename)
	}
	if !strings.Contains(filename, ".csv") {
		filename = strings.Join([]string{filename, csvExtension}, "")
	}

	gameArchivesURLList, err := e.Client.GetMonthlyGameArchivesURLList()
	if err != nil {
		panic(err)
	}

	out, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Panicln("Cannot open CSV file: ", err)
	}
	defer closeCSVFile(out)

	w := csv.NewWriter(out)
	if err = w.Write(model.GetGameExportCSVHeader()); err != nil {
		log.Panicln("Cannot write CSV header: ", err)
	}

	for _, gamesByMonthURL := range gameArchivesURLList {
		urlLen := len(gamesByMonthURL)
		log.Printf("Requesting games for %s\n", gamesByMonthURL[urlLen-7:])

		games, err := e.Client.GetMonthlyGameArchive(gamesByMonthURL)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, g := range games {
			if err = w.Write(g.ToGameExport(e.Client.Username).ToStringSlice()); err != nil {
				log.Println(err)
			}
		}
	}

	w.Flush()
	return w.Error()
}

func closeCSVFile(out *os.File) {
	log.Println("Closing CSV file..")
	if err := out.Close(); err != nil {
		log.Println("Error closing CSV file:", err)
	} else {
		log.Println("CSV file successfully closed")
	}
}