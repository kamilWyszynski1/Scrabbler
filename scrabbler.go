package scrabble

import (
	"encoding/json"
	"errors"
	"fmt"
)

const Dimension = int(7)

type FinderEngine interface {
	// FindWord finds possible  words from given data
	// letters - user's letters
	// word - one of a word from board
	FindWord(letters []rune, word string) []Word
}

type GameEngine interface {
	// Put checks if word can be placed and returns value of placed word
	Put(word PutRequest) (int, error)
}

type Cord struct {
	X int
	Y int
}

// PlacedPlate is structure used to distinguished placed letters on board
type PlacedPlate struct {
	Letter rune `json:"letter"`
	Cord
}

func (p *PlacedPlate) Unmarshal(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	fmt.Println(p)
	return nil
}

type PutRequest struct {
	Plates []PlacedPlateRequest `json:"plates"`
}

// only from http communication
type PlacedPlateRequest struct {
	Letter string `json:"letter"`
	Cord
}

type Board struct {
	Letters       map[Cord]rune
	BonusOccupied map[Cord]bool
	IsEmpty       bool
}

func (b *Board) SetLetter(plate PlacedPlate) { b.Letters[plate.Cord] = plate.Letter }

func (b *Board) SetBonusOccupied(cord Cord) { b.BonusOccupied[cord] = true }

type Word struct {
	Meaning   string
	Histogram map[rune]int
}

func (w Word) IsEqaul(w2 Word) bool {
	return w.Meaning == w2.Meaning
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
	ErrWrongPlatesSetup = errors.New("cannot be placed!")
)
