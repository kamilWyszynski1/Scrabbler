package game

import (
	"fmt"
	"scrabble"
	"scrabble/bonus"
	"scrabble/common"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_findWordCreated(t *testing.T) {
	engine, _ := NewGameEngine(logrus.New())
	letters := make(map[scrabble.Cord]rune)
	for i := -scrabble.Dimension; i <= scrabble.Dimension; i += 1 {
		for j := -scrabble.Dimension; j <= scrabble.Dimension; j += 1 {
			letters[scrabble.Cord{X: i, Y: j}] = 0
		}
	}

	letters[scrabble.Cord{1, 1}] = 'a'
	letters[scrabble.Cord{2, 1}] = 'b'
	letters[scrabble.Cord{3, 1}] = 'b'
	letters[scrabble.Cord{4, 1}] = 'a'
	letters[scrabble.Cord{6, 1}] = 'a'

	if word, points := engine.findAndCalculateCreatedWord(letters, 1, 1, "", 0, 1, scrabble.DirRight); word != "abba" {
		t.Error("found word isn't abba")
	} else {
		fmt.Println(points)
	}

	letters[scrabble.Cord{1, 1}] = 'a'
	letters[scrabble.Cord{1, 2}] = 'b'
	letters[scrabble.Cord{2, 3}] = 'b'
	letters[scrabble.Cord{3, 4}] = 'a'
	letters[scrabble.Cord{4, 1}] = 'a'
}

func TestEngine_calcRows(t *testing.T) {
	engine, _ := NewGameEngine(logrus.New())
	engine.Dictionary = append(engine.Dictionary, scrabble.Word{
		Meaning:   "cxb",
		Histogram: nil,
	})
	engine.Dictionary = append(engine.Dictionary, scrabble.Word{
		Meaning:   "axx",
		Histogram: nil,
	})

	/*

	  aa
	 cxb
	 cxb
	   b
	*/

	engine.Board.Letters[scrabble.Cord{2, 0}] = 'a'
	engine.Board.Letters[scrabble.Cord{3, 0}] = 'a'
	engine.Board.Letters[scrabble.Cord{3, 1}] = 'b'
	engine.Board.Letters[scrabble.Cord{3, 2}] = 'b'
	engine.Board.Letters[scrabble.Cord{3, 3}] = 'b'
	engine.Board.Letters[scrabble.Cord{1, 1}] = 'c'
	engine.Board.Letters[scrabble.Cord{1, 2}] = 'c'

	points, err := engine.calcRows(
		[]scrabble.PlacedPlate{
			{'x', scrabble.Cord{2, 1}},
			{'x', scrabble.Cord{2, 2}},
		}, 2, 1, 2)

	fmt.Println(points, err)
}

func TestEngine_Put(t *testing.T) {
	dictionary := []scrabble.Word{
		{Meaning: "aaa", Histogram: nil},
		{Meaning: "abbb", Histogram: nil},
		{Meaning: "aaay", Histogram: nil},
		{Meaning: "accc", Histogram: nil},
		{Meaning: "cxby", Histogram: nil},
		{Meaning: "cxbg", Histogram: nil},
		{Meaning: "yyg", Histogram: nil},
		{Meaning: "cxb", Histogram: nil},
		{Meaning: "axx", Histogram: nil},
		{Meaning: "elo", Histogram: nil},
		{Meaning: "aaayo", Histogram: nil},
	}
	type args struct {
		word []scrabble.PlacedPlateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "aaa",
			args: args{word: []scrabble.PlacedPlateRequest{
				{Letter: "a", Cord: scrabble.Cord{0, 0}},
				{Letter: "a", Cord: scrabble.Cord{1, 0}},
				{Letter: "a", Cord: scrabble.Cord{2, 0}},
			}},
			want: 3,
		},
		{
			name: "ccc",
			args: args{
				[]scrabble.PlacedPlateRequest{
					{Letter: "c", Cord: scrabble.Cord{0, 1}},
					{Letter: "c", Cord: scrabble.Cord{0, 2}},
					{Letter: "c", Cord: scrabble.Cord{0, 3}},
				},
			},
			want: 10,
		},
		{
			name: "bbb",
			args: args{
				[]scrabble.PlacedPlateRequest{
					{Letter: "b", Cord: scrabble.Cord{2, 1}},
					{Letter: "b", Cord: scrabble.Cord{2, 2}},
					{Letter: "b", Cord: scrabble.Cord{2, 3}},
				},
			},
			want: 16,
		},
		{
			name: "xx",
			args: args{
				[]scrabble.PlacedPlateRequest{
					{Letter: "x", Cord: scrabble.Cord{1, 1}},
					{Letter: "x", Cord: scrabble.Cord{1, 2}},
				},
			},
			want: 61,
		},
		{
			name: "yyg",
			args: args{
				[]scrabble.PlacedPlateRequest{
					{Letter: "y", Cord: scrabble.Cord{3, 0}},
					{Letter: "y", Cord: scrabble.Cord{3, 1}},
					{Letter: "g", Cord: scrabble.Cord{3, 2}},
				},
			},
			want:    51,
			wantErr: false,
		},
		{
			name: "elo",
			args: args{
				[]scrabble.PlacedPlateRequest{
					{Letter: "e", Cord: scrabble.Cord{4, -2}},
					{Letter: "l", Cord: scrabble.Cord{4, -1}},
					{Letter: "o", Cord: scrabble.Cord{4, 0}},
				},
			},
			want:    13,
			wantErr: false,
		},
	}
	e, _ := NewGameEngine(logrus.New())
	e.Dictionary = append(e.Dictionary, dictionary...)
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			got, err := e.Put(scrabble.PutRequest{tt.args.word})
			if (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			common.PrettyPrintBoard(e.Board.Letters)
			if got != tt.want {
				t.Errorf("Put() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func prettyPrintBonuses(bonuses map[scrabble.Cord]bool) {
	fmt.Println("-7  -6 -5 -4 -3 -2 -1  0  1  2  3  4  5  6  7")
	for i := -scrabble.Dimension; i <= scrabble.Dimension; i += 1 {
		row := ""
		for j := -scrabble.Dimension; j <= scrabble.Dimension; j += 1 {
			cord := scrabble.Cord{j, i}
			if v, ok := bonus.Cords[cord]; ok {
				if vb, _ := bonuses[cord]; vb {
					row += v.ToStringColor()
				} else {
					row += v.ToStringColor()
				}
			}
			row += fmt.Sprintf("%s   ", "0")
		}
		fmt.Println(row)
	}
}
