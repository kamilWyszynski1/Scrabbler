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

func NewEmptyBoard() *scrabble.Board {
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

func NewGameEngine() (*Engine, error) {
	engine := Engine{
		log:         logrus.New().WithField("service", "GameEngine"),
		Board:       NewEmptyBoard(),
		Multipliers: bonus.Cords,
	}

	if dict, err := common.LoadWordsDirectory(); err != nil {
		return nil, fmt.Errorf("failed to load word directory")
	} else {
		engine.Dictionary = dict
	}
	return &engine, nil
}

func (e *Engine) Put(word []scrabble.PlacedPlate) (int, error) {
	if !e.canBePut(word) {
		return 0, scrabble.ErrPlateOccupied
	}

	minX, maxX, minY, maxY := findCordsRange(word)

	var pointSum int

	points, err := e.calcRows(word, minX, minY, maxY)
	if err != nil {
		return 0, fmt.Errorf("failed to calculate row, %w", err)
	}
	pointSum += points

	points, err = e.calcCols(word, minY, minX, maxX)
	if err != nil {
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
	for _, w := range word {
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
