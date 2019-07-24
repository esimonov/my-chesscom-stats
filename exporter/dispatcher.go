package exporter

import (
	"fmt"
	"log"

	"github.com/esimonov/my-chesscom-stats/client"
)

// Dispatcher is those who dispatches chess stats exporters.
type Dispatcher struct {
	Client *client.Client
}

// NewDispatcher creates and returns new dispatcher.
func NewDispatcher(username string) *Dispatcher {
	client := client.NewClient(username)
	return &Dispatcher{
		Client: client,
	}
}

// AllGames feeds all games played by provided username to exporters.
func (d *Dispatcher) AllGames(exporters []Exporter) error {
	if exporters == nil {
		return fmt.Errorf("No exporters provided")
	}

	// Open exporters.
	for _, e := range exporters {
		if err := e.Open(); err != nil {
			return err
		}
		defer e.Close()
	}

	gameArchivesURLList, err := d.Client.GetMonthlyGameArchivesURLList()
	if err != nil {
		return err
	}

	countryURLsToNames := make(map[string]string, 1)

	// Fetch games and feed them to exporters.
	for _, gamesByMonthURL := range gameArchivesURLList {
		urlLen := len(gamesByMonthURL)
		log.Printf("Requesting games for %s\n", gamesByMonthURL[urlLen-7:])

		games, err := d.Client.GetMonthlyGameArchive(gamesByMonthURL)
		if err != nil {
			log.Println(err)
			continue
		}

		for _, g := range games {
			opponent, err := d.Client.GetPlayerByUsername(g.GetOpponentUsername(d.Client.Username))
			if err != nil {
				log.Println(err)
				continue
			}

			countryName, ok := countryURLsToNames[opponent.CountryURL]
			if !ok {
				country, err := d.Client.GetCountryByURL(opponent.CountryURL)
				if err != nil {
					log.Println(err)
					continue
				}
				countryName = country.Name
				countryURLsToNames[opponent.CountryURL] = countryName
			}

			for _, e := range exporters {
				if err = e.ExportGame(g, countryName); err != nil {
					log.Println(err)
				}
			}
		}
	}
	return nil
}
