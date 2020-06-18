package game

import (
	"fmt"
	"scrabble"
	"scrabble/bonus"
	"scrabble/common"

	"github.com/sirupsen/logrus"
)

type Engine struct {
	log         logrus.FieldLogger
	Dictionary  []scrabble.Word
	Board       *scrabble.Board
	Multipliers map[scrabble.Cord]bonus.Bonus
}

// newEmptyBoard initialize new, clean board
func newEmptyBoard() *scrabble.Board {
	var board scrabble.Board
	letters := make(map[scrabble.Cord]rune)
	bonusOccupied := make(map[scrabble.Cord]bool)
	for i := -scrabble.Dimension; i <= scrabble.Dimension; i += 1 {
		for j := -scrabble.Dimension; j <= scrabble.Dimension; j += 1 {
			letters[scrabble.Cord{X: j, Y: i}] = 0
			bonusOccupied[scrabble.Cord{X: j, Y: i}] = false
		}
	}
	board.Letters = letters
	board.BonusOccupied = bonusOccupied
	return &board
}

// NewGameEngine is constructor function
func NewGameEngine(log logrus.FieldLogger) (*Engine, error) {
	engine := Engine{
		log:         log.WithField("service", "GameEngine"),
		Board:       newEmptyBoard(),
		Multipliers: bonus.Cords,
	}

	if dict, err := common.LoadWordsDirectory(); err != nil {
		return nil, fmt.Errorf("failed to load word directory")
	} else {
		engine.Dictionary = dict
	}
	return &engine, nil
}

// Put method checks if plates can be put, if placed plates create valid words
// at last it counts points and returns result
func (e *Engine) Put(wordReq scrabble.PutRequest) (int, error) {
	var word []scrabble.PlacedPlate
	for _, letter := range wordReq.Plates {
		word = append(word, scrabble.PlacedPlate{
			Letter: rune(letter.Letter[0]),
			Cord:   letter.Cord,
		})
	}
	fmt.Println(word)

	defer PrettyPrintBoard(e.Board.Letters)

	if !e.canBePut(word) {
		e.log.Error(scrabble.ErrPlateOccupied)
		return 0, scrabble.ErrPlateOccupied
	}

	minX, maxX, minY, maxY := findCordsRange(word)

	var pointSum int

	points, err := e.calcRows(word, minX, minY, maxY)
	if err != nil {
		e.log.WithError(err).Error("failed to calculate row")
		return 0, fmt.Errorf("failed to calculate row, %w", err)
	}
	pointSum += points

	points, err = e.calcCols(word, minY, minX, maxX)
	if err != nil {
		e.log.WithError(err).Error("failed to calculate column")
		return 0, fmt.Errorf("failed to calculate column, %w", err)
	}
	pointSum += points

	// put plates for real
	for _, w := range word {
		e.Board.SetLetter(w)
		e.Board.SetBonusOccupied(w.Cord)
	}

	return pointSum, nil
}

func (e Engine) canBePut(word []scrabble.PlacedPlate) bool {
	cordsInPlacedPlates := map[scrabble.Cord]struct{}{}

	for _, w := range word {
		if _, ok := cordsInPlacedPlates[w.Cord]; ok {
			return false
		}
		cordsInPlacedPlates[w.Cord] = struct{}{}
		if v, ok := e.Board.Letters[w.Cord]; ok {
			if v != 0 {
				return false
			}
		}
	}
	return true
}

func (e Engine) isValidWord(word string) bool {
	for _, dictWord := range e.Dictionary {
		if word == dictWord.Meaning {
			return true
		}
	}
	return false
}

func findCordsRange(word []scrabble.PlacedPlate) (minX, maxX, minY, maxY int) {
	minX, maxX, minY, maxY = 7, -7, 7, -7
	for _, w := range word {
		if minX > w.X {
			minX = w.X
		}
		if maxX < w.X {
			maxX = w.X
		}
		if minY > w.Y {
			minY = w.Y
		}
		if maxY < w.Y {
			maxY = w.Y
		}
	}
	return
}

const (
	InfoColor    = "\033[1;34m%s  \033[0m"
	NoticeColor  = "\033[1;36m%s  \033[0m"
	WarningColor = "\033[1;33m%s  \033[0m"
	ErrorColor   = "\033[1;31m%s  \033[0m"
	DebugColor   = "\033[0;36m%s  \033[0m"
)

func PrettyPrintBoard(board map[scrabble.Cord]rune) {
	fmt.Println("-7 -6 -5 -4 -3 -2 -1  0  1  2  3  4  5  6  7")
	for i := -scrabble.Dimension; i <= scrabble.Dimension; i += 1 {
		row := ""
		for j := -scrabble.Dimension; j <= scrabble.Dimension; j += 1 {
			cord := scrabble.Cord{j, i}
			if v, _ := board[cord]; v == 0 {
				row += fmt.Sprintf(InfoColor, "0")
			} else {
				row += fmt.Sprintf(ErrorColor, string(v))
			}
		}
		fmt.Println(row)
	}
}
