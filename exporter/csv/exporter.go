package exporter

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/esimonov/my-chesscom-stats/model"
)

const csvExtension = ".csv"

// CSVExporter exports chess stats to CSV file.
type CSVExporter struct {
	filename string
	file     *os.File
	writer   *csv.Writer
	username string
}

// NewExporter creates new CSV exporter.
// Filename is a name of CSV sheet. If extention is not provided, it will be added automatically.
func NewExporter(filename, username string) (*CSVExporter, error) {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		return nil, fmt.Errorf("Empty filename")
	}
	if !strings.Contains(filename, csvExtension) {
		filename = strings.Join([]string{filename, csvExtension}, "")
	}

	return &CSVExporter{
		filename: filename,
		username: username,
	}, nil
}

// Open opens exporter's writer.
func (e *CSVExporter) Open() error {
	out, err := os.OpenFile(e.filename, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("Cannot open CSV file: %s", err)
	}

	w := csv.NewWriter(out)
	if err = w.Write(model.GetGameExportCSVHeader()); err != nil {
		return fmt.Errorf("Cannot write CSV header: %s", err)
	}

	e.file = out
	e.writer = w
	return nil
}

// ExportGame writes game to CSV file.
func (e *CSVExporter) ExportGame(g model.Game, countryName string) error {
	return e.writer.Write(g.ToGameExport(e.username, countryName).ToStringSlice())
}

// Close closes CSV file.
func (e *CSVExporter) Close() {
	e.writer.Flush()
	if err := e.writer.Error(); err != nil {
		log.Println("Error occured during the flush:", err)
	}

	log.Println("Closing CSV file..")
	if err := e.file.Close(); err != nil {
		log.Println("Error closing CSV file:", err)
	} else {
		log.Println("CSV file successfully closed")
	}
}
