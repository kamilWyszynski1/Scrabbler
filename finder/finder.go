package finder

import (
	"fmt"
	"scrabble"
	"scrabble/bonus"
	"strings"

	"github.com/sirupsen/logrus"
)

type Engine struct {
	log        logrus.FieldLogger
	Dictionary []scrabble.Word
	Board      *scrabble.Board
}

func NewWordsFinder() (scrabble.FinderEngine, error) {
	finder := Engine{
		log:   logrus.New().WithField("service", "WordsFinder"),
		Board: scrabble.NewBoard(),
	}

	if err := finder.LoadWordsDirectory(); err != nil {
		return nil, fmt.Errorf("failed to load word directory")
	}
	return finder, nil
}

func (e Engine) FindWord(letters []rune, boardWord string) []scrabble.Word {
	histogram := make(map[rune]int)

	for _, letter := range letters {
		histogram[letter] += 1
	}

	for _, letter := range boardWord {
		histogram[letter] += 1
	}

	var possibleWords []scrabble.Word

	for _, word := range e.Dictionary {
		if compareHistograms(histogram, word.Histogram) {
			if boardWord != "" && strings.Contains(word.Meaning, boardWord) {
				possibleWords = append(possibleWords, word)
			} else if boardWord == "" {
				possibleWords = append(possibleWords, word)
			}
		}
	}

	return possibleWords
}

func (e Engine) Put(word map[scrabble.Cord]rune) (int, error) {
	// check if word can be placed
	for v := range word {
		if _, ok := e.Board.Letters[v]; ok {
			return 0, scrabble.ErrPlateOccupied
		}
	}

	var (
		value          = 0
		wordMultiplier = 1
	)
	for v, k := range word {
		e.Board.Letters[v] = k
		if premium, ok := bonus.Cords[v]; ok {
			if !premium.Used {
				premium.Used = true

				switch premium.BonusType {
				case bonus.LetterBonusType:
					value += scrabble.LetterValue[k] * premium.Multiplier.Int()
					continue
				case bonus.WordBonusType:
					wordMultiplier *= premium.Multiplier.Int()
				}
			}
		}
		value += scrabble.LetterValue[k]
	}

	//value += e.checkStartEnd(word)
	value *= wordMultiplier

	value += e.checkAdjoining(word)

	return value, nil
}

func (e Engine) checkAdjoining(word map[scrabble.Cord]rune) int {
	return 0
}

// checkStartEnd checks if there's any letter adjoining on the end and start of word
func (e *Engine) checkStartEnd(word map[scrabble.Cord]rune) (int, error) {
	var (
		dir     scrabble.Direction
		xOffset int
		yOffset int
		counter int

		first, second scrabble.Cord
	)

	for k := range word {
		if counter == 0 {
			first = k
		} else if counter == 1 {
			second = k
		} else {
			break
		}
	}

	xOffset = second.X - first.X
	yOffset = second.Y - first.Y

	if yOffset < 0 && xOffset == 0 {
		dir = scrabble.DirDown
	} else if xOffset < 0 && yOffset == 0 {
		dir = scrabble.DirRight
	} else {
		return 0, scrabble.ErrInvalidDirection
	}

	switch dir {
	case scrabble.DirRight:

	}

	return 0, nil
}

func (e Engine) findAdjoining() {

}

// word {'a':1, 'b':1, 'e':1, 'l':1 }
// letters {'a':4, }
func compareHistograms(lettersHistogram, wordHistogram map[rune]int) bool {
	for k, v := range wordHistogram {
		if value, ok := lettersHistogram[k]; !ok {
			return false
		} else if ok && value < v {
			return false
		}
	}
	return true
}
