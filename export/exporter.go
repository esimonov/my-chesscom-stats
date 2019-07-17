package export

// Exporter is those who exports chess stats.
type Exporter interface {
	AllGames(filename string) error
}
