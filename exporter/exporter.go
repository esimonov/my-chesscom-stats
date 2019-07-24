package exporter

import (
	"github.com/esimonov/my-chesscom-stats/model"
)

// Exporter is those who exports chess stats to particular destination.
type Exporter interface {
	Open() error
	ExportGame(g model.Game, countryName string) error
	Close()
}
