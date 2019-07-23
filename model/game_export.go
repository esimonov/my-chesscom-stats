package model

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

// GameExport is how game is exported.
type GameExport struct {
	Date              string `csv:"Date"`
	StartTime         string `csv:"Start time"`
	EndTime           string `csv:"End time"`
	ECO               string `csv:"ECO"`
	ECOExtended       string `csv:"ECO extended"`
	MyElo             int    `csv:"My Elo"`
	OpponentElo       int    `csv:"Opponent Elo"`
	MyColor           string `csv:"My color"`
	OpponentColor     string `csv:"Opponent color"`
	OpponentCountry   string `csv:"Opponent country"`
	MyWin             bool   `csv:"My win"`
	NumMoves          int    `csv:"Moves"`
	TerminationReason string `csv:"Termination reason"`
	TimeControl       string `csv:"Time control"`
}

// ToStringSlice transforms GameExport into string slice (i.e, CSV row).
func (ge *GameExport) ToStringSlice() (res []string) {
	if ge == nil {
		return
	}

	v := reflect.ValueOf(*ge)

	for j := 0; j < v.NumField(); j++ {
		field := v.Field(j)
		switch field.Type().Kind() {
		case reflect.String:
			res = append(res, field.String())
		case reflect.Int:
			res = append(res, strconv.Itoa(int(field.Int())))
		case reflect.Bool:
			res = append(res, strconv.FormatBool(field.Bool()))
		}
	}
	return
}

// GetGameExportCSVHeader returns csv header based on GameExport csv tags.
func GetGameExportCSVHeader() []string {
	ge := GameExport{}
	t := reflect.TypeOf(ge)
	header := make([]string, t.NumField())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		header[i] = strings.Split(field.Tag.Get("csv"), ",")[0]
	}
	return header
}

func (ge *GameExport) parsePlayers(white, black player, myUsername string) {
	if white.Username == myUsername {
		ge.MyColor = "white"
		ge.MyElo = white.Rating
		ge.OpponentColor = "black"
		ge.OpponentElo = black.Rating
		if white.Result == "win" {
			ge.MyWin = true
			ge.TerminationReason = black.Result
		} else {
			ge.MyWin = false
			ge.TerminationReason = white.Result
		}
	} else {
		ge.MyColor = "black"
		ge.MyElo = black.Rating
		ge.OpponentColor = "white"
		ge.OpponentElo = white.Rating
		if black.Result == "win" {
			ge.MyWin = true
			ge.TerminationReason = white.Result
		} else {
			ge.MyWin = false
			ge.TerminationReason = black.Result
		}
	}
}

func (ge *GameExport) parsePGN(pgn string) {
	s := new(scanner.Scanner)
	s.Init(strings.NewReader(pgn))
	ge.parsePGNTags(s)
	ge.parseMoves(s)
}

func (ge *GameExport) parsePGNTags(s *scanner.Scanner) {
	tagsMap := make(map[string]string)
	run := s.Peek()
Loop:
	for run != scanner.EOF {
		switch run {
		case '[', ']', '\n', '\r':
			run = s.Next()
		case '1':
			break Loop
		default:
			s.Scan()
			tag := s.TokenText()
			s.Scan()
			val := s.TokenText()
			tagsMap[tag] = strings.Trim(val, "\"")
		}
		run = s.Peek()
	}

	ge.Date = tagsMap["Date"]
	ge.StartTime = tagsMap["StartTime"]
	ge.EndTime = tagsMap["EndTime"]
	ge.ECO = tagsMap["ECO"]

	ecoSplit := strings.Split(tagsMap["ECOUrl"], "/")
	ge.ECOExtended = ecoSplit[len(ecoSplit)-1]
}

func (ge *GameExport) parseMoves(s *scanner.Scanner) {
	s.Mode = scanner.ScanFloats
	s.Error = processScannerErrors

	run := s.Peek()
	var numMoves int

	for run != scanner.EOF {
		s.Scan()
		token := s.TokenText()

		lenToken := len(token)
		if lenToken < 2 {
			run = s.Peek()
			continue
		}

		if token[lenToken-1] != '.' {
			run = s.Peek()
			continue
		}

		if numMovesInt, err := strconv.Atoi(token[:lenToken-1]); err == nil {
			numMoves = numMovesInt
		}
		run = s.Peek()
	}
	ge.NumMoves = numMoves
}

const errIllegalOctalNumber = "illegal octal number"

func processScannerErrors(s *scanner.Scanner, msg string) {
	if msg == errIllegalOctalNumber {
		return
	}

	pos := s.Position
	if !pos.IsValid() {
		pos = s.Pos()
	}
	fmt.Printf("%s: %s\n", pos, msg)
}
