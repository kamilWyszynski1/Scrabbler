package game

import (
	"fmt"
	"scrabble"
	"scrabble/bonus"
)

func (e *Engine) calcRows(word []scrabble.PlacedPlate, startX, minY, maxY int) (int, error) {
	var pointSum int

	boardCopy := e.Board
	for _, w := range word {
		boardCopy.Letters[w.Cord] = w.Letter
	}

	for y := minY; y <= maxY; y += 1 {
		var (
			wordCreated string
			rowSum      int
		)
		// find first X cord of created x
		for x := startX; x >= -scrabble.Dimension; x -= 1 {
			if v, ok := boardCopy.Letters[scrabble.Cord{x, y}]; ok && v != 0 {
				startX = x
				continue
			}
			break
		}

		// recursively created word in one row
		wordCreated, rowSum = e.findAndCalculateCreatedWord(boardCopy.Letters, startX, y, wordCreated, rowSum, 1, scrabble.DirRight)
		if len(wordCreated) < 2 {
			continue
		}
		if !e.isValidWord(wordCreated) {
			return 0, fmt.Errorf("invalid word created: %s", wordCreated)
		}
		pointSum += rowSum
	}
	return pointSum, nil
}

func (e *Engine) calcCols(word []scrabble.PlacedPlate, startY int, minX int, maxX int) (int, error) {
	var pointSum int

	boardCopy := e.Board
	for _, w := range word {
		boardCopy.Letters[w.Cord] = w.Letter
	}

	for x := minX; x <= maxX; x += 1 {
		var (
			wordCreated string
			rowSum      int
		)
		// find first X cord of created x
		for y := startY; y >= -scrabble.Dimension; y -= 1 {
			cord := scrabble.Cord{x, y}
			if v, ok := boardCopy.Letters[cord]; ok && v != 0 {
				startY = y
				continue
			}
			break
		}

		// recursively created word in one row
		wordCreated, rowSum = e.findAndCalculateCreatedWord(boardCopy.Letters, x, startY, wordCreated, rowSum, 1, scrabble.DirDown)
		if len(wordCreated) < 2 {
			continue
		}
		if !e.isValidWord(wordCreated) {
			return 0, fmt.Errorf("invalid word created: %s", wordCreated)
		}
		pointSum += rowSum
	}
	return pointSum, nil
}

func (e *Engine) findAndCalculateCreatedWord(boardCopy map[scrabble.Cord]rune, x int, y int, word string, sum int, multiplier int, direction scrabble.Direction) (string, int) {
	if v, ok := boardCopy[scrabble.Cord{X: x, Y: y}]; !ok || v == 0 {
		return word, sum
	}
	cord := scrabble.Cord{X: x, Y: y}
	letter := boardCopy[cord]
	word += string(letter)

	if v, ok := bonus.Cords[cord]; ok {
		if !e.Board.BonusOccupied[cord] {
			switch v.BonusType {
			case bonus.LetterBonusType:
				sum += v.Multiplier.Int() * scrabble.LetterValue[letter]
			case bonus.WordBonusType:
				sum += scrabble.LetterValue[letter]
				multiplier *= v.Multiplier.Int()
			}
		} else {
			sum += scrabble.LetterValue[letter]
		}
	} else {
		sum += scrabble.LetterValue[letter]
	}

	switch direction {
	case scrabble.DirRight:
		return e.findAndCalculateCreatedWord(boardCopy, x+1, y, word, sum, multiplier, direction)
	case scrabble.DirDown:
		return e.findAndCalculateCreatedWord(boardCopy, x, y+1, word, sum, multiplier, direction)
	default:
		return word, sum
	}
}
