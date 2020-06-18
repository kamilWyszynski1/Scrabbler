package bonus

import (
	"fmt"
	"scrabble"
)

type Multiplier int

const (
	DoubleMultiplier = Multiplier(2)
	TripleMultiplier = Multiplier(3)
)

func (m Multiplier) Int() int {
	return int(m)
}

type Type string

const (
	WordBonusType   = Type("word")
	LetterBonusType = Type("letter")
)

type Bonus struct {
	Multiplier Multiplier
	BonusType  Type
	Used       bool
}

func (b Bonus) ToStringColor() string {
	switch b.Multiplier {
	case DoubleMultiplier:
		switch b.BonusType {
		case WordBonusType:
			return fmt.Sprintf("\033[1;31m%s  \033[0m", "DW")
		case LetterBonusType:
			return fmt.Sprintf("\u001B[1;34m%s  \u001B[0m", "DL")
		}
	case TripleMultiplier:
		switch b.BonusType {
		case WordBonusType:
			return fmt.Sprintf("\033[1;33m%s  \033[0m", "TW")
		case LetterBonusType:
			return fmt.Sprintf("\u001B[1;36m%s  \u001B[0m", "TL")
		}
	}
	return ""
}

func (b Bonus) ToString() string {
	switch b.Multiplier {
	case DoubleMultiplier:
		switch b.BonusType {
		case WordBonusType:
			return fmt.Sprintf("%s  ", "DW")
		case LetterBonusType:
			return fmt.Sprintf("%s  ", "DL")
		}
	case TripleMultiplier:
		switch b.BonusType {
		case WordBonusType:
			return fmt.Sprintf("%s  ", "TW")
		case LetterBonusType:
			return fmt.Sprintf("%s  ", "TL")
		}
	}
	return ""
}

var Cords = map[scrabble.Cord]Bonus{
	// TripleWordScores
	scrabble.Cord{-7, 7}:  {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{0, 7}:   {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{7, 7}:   {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{7, 0}:   {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{7, -7}:  {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{0, -7}:  {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{-7, -7}: {TripleMultiplier, WordBonusType, false},
	scrabble.Cord{-7, 0}:  {TripleMultiplier, WordBonusType, false},

	// DoubleWordScores
	scrabble.Cord{-6, 6}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{6, 6}:   {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{6, -6}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{-6, -6}: {DoubleMultiplier, WordBonusType, false},

	scrabble.Cord{-5, 5}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{5, 5}:   {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{5, -5}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{-5, -5}: {DoubleMultiplier, WordBonusType, false},

	scrabble.Cord{-4, 4}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{4, 4}:   {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{4, -4}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{-4, -4}: {DoubleMultiplier, WordBonusType, false},

	scrabble.Cord{-3, 3}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{3, 3}:   {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{3, -3}:  {DoubleMultiplier, WordBonusType, false},
	scrabble.Cord{-3, -3}: {DoubleMultiplier, WordBonusType, false},

	//TripleLetterScores
	scrabble.Cord{-2, 6}:  {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{2, 6}:   {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{2, -6}:  {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{-2, -6}: {TripleMultiplier, LetterBonusType, false},

	scrabble.Cord{-6, 2}:  {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{-6, -2}: {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{6, 2}:   {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{6, -2}:  {TripleMultiplier, LetterBonusType, false},

	scrabble.Cord{2, 2}:   {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{2, -2}:  {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{-2, 2}:  {TripleMultiplier, LetterBonusType, false},
	scrabble.Cord{-2, -2}: {TripleMultiplier, LetterBonusType, false},

	//DoubleLetterScores
	scrabble.Cord{4, 7}:   {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{4, -7}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-4, 7}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-4, -7}: {DoubleMultiplier, LetterBonusType, false},

	scrabble.Cord{7, 4}:   {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{7, -4}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-7, 4}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-7, -4}: {DoubleMultiplier, LetterBonusType, false},

	scrabble.Cord{1, 5}:   {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{1, -5}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-1, 5}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-1, -5}: {DoubleMultiplier, LetterBonusType, false},

	scrabble.Cord{5, 1}:   {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{5, -1}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-5, 1}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-5, -1}: {DoubleMultiplier, LetterBonusType, false},

	scrabble.Cord{0, 4}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{0, -4}: {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{4, 0}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-4, 0}: {DoubleMultiplier, LetterBonusType, false},

	scrabble.Cord{1, 1}:   {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{1, -1}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-1, 1}:  {DoubleMultiplier, LetterBonusType, false},
	scrabble.Cord{-1, -1}: {DoubleMultiplier, LetterBonusType, false},
}
