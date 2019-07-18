package model

// Game is a played game.
type Game struct {
	URL         string `json:"url"`
	PGN         string `json:"pgn"`
	TimeControl string `json:"time_control"`
	StartTime   int    `json:"start_time"`
	EndTime     int    `json:"end_time"`
	Rated       bool   `json:"rated"`
	FEN         string `json:"fen"`
	TimeClass   string `json:"time_class"`
	Rules       string `json:"rules"`
	White       player `json:"white"`
	Black       player `json:"black"`
	Tournament  string `json:"tournament"`
	Match       string `json:"match"`
}

type player struct {
	Rating   int    `json:"rating"`
	Result   string `json:"result"`
	APIID    string `json:"@id"`
	Username string `json:"username"`
}

// ToGameExport transorms the game into exported form.
func (g *Game) ToGameExport(myUsername, countryName string) *GameExport {
	ge := new(GameExport)
	ge.parsePGNTags(g.PGN)
	ge.parsePlayers(g.White, g.Black, myUsername)
	ge.TimeControl = g.TimeControl
	ge.OpponentCountry = countryName
	return ge
}

// GetOpponentUsername returns username of your opponent in the game.
func (g *Game) GetOpponentUsername(myUsername string) string {
	if g.White.Username == myUsername {
		return g.Black.Username
	}
	return g.White.Username
}
