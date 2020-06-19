package game

import (
	"scrabble"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_arePlatesInOneLine(t *testing.T) {
	type args struct {
		word []scrabble.PlacedPlate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "In single X line",
			args: args{
				[]scrabble.PlacedPlate{
					{
						Letter: 'a', Cord: scrabble.Cord{0, 1},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{0, 2},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{0, 3},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{0, 4},
					},
				},
			},
			want: true,
		},
		{
			name: "In single Y line",
			args: args{
				[]scrabble.PlacedPlate{
					{
						Letter: 'a', Cord: scrabble.Cord{1, 0},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{2, 0},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{3, 0},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{4, 0},
					},
				},
			},
			want: true,
		},
		{
			name: "Basic",
			args: args{
				[]scrabble.PlacedPlate{
					{
						Letter: 'a', Cord: scrabble.Cord{0, 1},
					},
					{
						Letter: 'a', Cord: scrabble.Cord{2, 2},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arePlatesInOneLine(tt.args.word, nil); got != tt.want {
				t.Errorf("arePlatesInOneLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adjoinsWithOtherPlates(t *testing.T) {
	type args struct {
		word  []scrabble.PlacedPlate
		board *scrabble.Board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Basic",
			args: args{
				word: []scrabble.PlacedPlate{
					{Letter: 'a', Cord: scrabble.Cord{0, 1}},
					{Letter: 'a', Cord: scrabble.Cord{0, 2}},
					{Letter: 'a', Cord: scrabble.Cord{0, 3}},
					{Letter: 'a', Cord: scrabble.Cord{0, 4}},
				},
				board: &scrabble.Board{
					Letters: nil,
					IsEmpty: true,
				},
			},
			want: true,
		},
		{
			name: "Basic",
			args: args{
				word: []scrabble.PlacedPlate{
					{Letter: 'a', Cord: scrabble.Cord{0, 1}},
					{Letter: 'a', Cord: scrabble.Cord{0, 2}},
					{Letter: 'a', Cord: scrabble.Cord{0, 3}},
					{Letter: 'a', Cord: scrabble.Cord{0, 4}},
				},
				board: &scrabble.Board{
					Letters: map[scrabble.Cord]rune{
						scrabble.Cord{-1, 1}: 'b',
					},
				},
			},
			want: true,
		},
		{
			name: "Basic",
			args: args{
				word: []scrabble.PlacedPlate{
					{Letter: 'a', Cord: scrabble.Cord{0, 1}},
					{Letter: 'a', Cord: scrabble.Cord{0, 2}},
					{Letter: 'a', Cord: scrabble.Cord{0, 3}},
					{Letter: 'a', Cord: scrabble.Cord{0, 4}},
				},
				board: &scrabble.Board{
					Letters: map[scrabble.Cord]rune{
						scrabble.Cord{0, 5}: 'b',
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adjoinsWithOtherPlates(tt.args.word, tt.args.board, nil); got != tt.want {
				t.Errorf("adjoinsWithOtherPlates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEngine_canBePut(t *testing.T) {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	board := &scrabble.Board{
		Letters: map[scrabble.Cord]rune{
			{0, 0}: 'a',
			{0, 1}: 'a',
			{0, 2}: 'a',
			{0, 3}: 'a',
			{0, 4}: 'a',
		},
		IsEmpty: false,
	}
	type args struct {
		word []scrabble.PlacedPlate
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "single plate",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{1, 2}},
				},
			},
			want: true,
		},
		{
			name: "single plate, don't adjoin",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{5, 2}},
				},
			},
			want: false,
		},
		{
			name: "single plate, overlap",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{0, 2}},
				},
			},
			want: false,
		},
		{
			name: "multiple plates, adjoin",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{1, 2}},
					{'a', scrabble.Cord{2, 2}},
					{'a', scrabble.Cord{3, 2}},
				},
			},
			want: true,
		},
		{
			name: "multiple plates, don't adjoin",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{2, 3}},
					{'a', scrabble.Cord{3, 3}},
					{'a', scrabble.Cord{4, 3}},
				},
			},
			want: false,
		},
		{
			name: "multiple plates, divided by placed plates",
			args: args{
				word: []scrabble.PlacedPlate{
					{'a', scrabble.Cord{-1, 0}},
					{'a', scrabble.Cord{1, 0}},
					{'a', scrabble.Cord{2, 0}},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Engine{
				log:   log,
				Board: board,
			}
			if got := e.canBePut(tt.args.word); got != tt.want {
				t.Errorf("canBePut() = %v, want %v", got, tt.want)
			}
		})
	}
}
