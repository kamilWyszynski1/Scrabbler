package bonus

import "scrabble"

type Multiplier int

const (
	DoubleMultiplier = Multiplier(2)
	TripleMultiplier = Multiplier(3)
)

type Type string

const (
	WordBonusType   = Type("word")
	LetterBonusType = Type("letter")
)

type Bonus struct {
	multiplier Multiplier
	bonusType  Type
	used       bool
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
