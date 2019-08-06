# my-chesscom-stats

### Description and motivation
This projects demonstrates fetching data from chess.com public API.
Currently, it can only export played games to a CSV sheet.  
I think that owning the data you produce is a healthy idea. It can also be a lot of fun to analyze it and gain insights.
After having some experience of doing data analysis with Python, I guess that CSV is an appropriate choice for exporting format, that's why the support for it was implemented primarily.
However, the `Exporter` is hidden under the interface, so any kind of storage can be supported with no pain.

Chess.com do not currently provide the API for solved tactics, which is a pity.

### Example
Exporting all games played by the `USERNAME` to `FILENAME.csv` only requires this setup:

```go
import (
    "github.com/esimonov/my-chesscom-stats/exporter"
    csv "github.com/esimonov/my-chesscom-stats/exporter/csv"
)

func main() {
    dispatcher := exporter.NewDispatcher(USERNAME)

    csvExporter, err := csv.NewExporter(FILENAME, USERNAME)
    if err != nil {
        panic(err)
	  }

    exporters := []exporter.Exporter{csvExporter}

    if err = dispatcher.AllGames(exporters); err != nil {
        panic(err)
    }
}
```
