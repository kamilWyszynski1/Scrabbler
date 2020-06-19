package game

import (
	"scrabble"

	"github.com/sirupsen/logrus"
)

// TODO add returning error insted of bool
func (e Engine) canBePut(word []scrabble.PlacedPlate) bool {
	cordsInPlacedPlates := map[scrabble.Cord]struct{}{}

	for _, w := range word {
		if _, ok := cordsInPlacedPlates[w.Cord]; ok {
			return false
		}
		cordsInPlacedPlates[w.Cord] = struct{}{}
		if v, ok := e.Board.Letters[w.Cord]; ok {
			if v != 0 {
				e.log.Debug("plates overlap")
				return false
			}
		}
	}
	return arePlatesInOneLine(word, e.log) && adjoinsWithOtherPlates(word, e.Board, e.log)
}

func (e Engine) isValidWord(word string) bool {
	for _, dictWord := range e.Dictionary {
		if word == dictWord.Meaning {
			return true
		}
	}
	return false
}

func adjoinsWithOtherPlates(word []scrabble.PlacedPlate, board *scrabble.Board, log logrus.FieldLogger) bool {
	for _, w := range word {
		for i := -1; i <= 1; i += 1 {
			for j := -1; j <= 1; j += 1 {
				cord := w.Cord
				cord.X -= i
				cord.Y -= j
				if v, ok := board.Letters[cord]; ok {
					if v != 0 {
						return true
					}
				}
			}
		}
	}
	log.Debug("plates don't adjoin")

	return board.IsEmpty && beginsFromPointZero(word, log)
}

func beginsFromPointZero(word []scrabble.PlacedPlate, log logrus.FieldLogger) bool {
	for _, w := range word {
		if w.Cord.X == 0 && w.Cord.Y == 0 {
			return true
		}
	}
	log.Debug("plates don't start from begging point")
	return false
}

func arePlatesInOneLine(word []scrabble.PlacedPlate, log logrus.FieldLogger) bool {
	x, y := word[0].X, word[0].Y

	for _, w := range word {
		if w.X != x {
			goto checkY
		}
	}
	return true

checkY:
	for _, w := range word {
		if w.Y != y {
			log.Debug("plates are not in one line")
			return false
		}
	}

	return true
}
