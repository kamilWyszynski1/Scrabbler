package scrabble

import "errors"

type FinderEngine interface {
	// FindWord finds possible  words from given data
	// letters - user's letters
	// word - one of a word from board
	FindWord(letters []rune, word string) []Word

	// Put checks if word can be placed and returns value of placed word
	Put(word map[Cord]rune) (int, error)
}

type GameEngine interface {
}

type Cord struct {
	X int
	Y int
}

type Board struct {
	Letters map[Cord]rune
}

func NewBoard() *Board {
	return &Board{map[Cord]rune{}}
}

type Word struct {
	Meaning   string
	Histogram map[rune]int
	Value     int
}

var LetterValue = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 1,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

type Direction int

const (
	DirDown  = Direction(0)
	DirRight = Direction(1)
)

var (
	ErrPlateOccupied    = errors.New("plate occupied")
	ErrInvalidDirection = errors.New("invalid direction")
)
