package finder

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"scrabble"
)

type Words struct {
	log        logrus.FieldLogger
	Dictionary []scrabble.Word
}

func NewWordsFinder() (scrabble.Finder, error) {
	finder := Words{
		log: logrus.New().WithField("service", "WordsFinder"),
	}

	if err := finder.LoadWordsDirectory(); err != nil {
		return nil, fmt.Errorf("failed to load word directory")
	}
	return finder, nil
}

func (w Words) FindWord(letters []rune) string {
	lettersHistogram := make(map[rune]int)

	for _, letter := range letters {
		lettersHistogram[letter] += 1
	}

	var bestWord scrabble.Word

	for _, word := range w.Dictionary {
		if compareHistograms(lettersHistogram, word.Histogram) {
			if bestWord.Value < word.Value {
				bestWord = word
			}
		}
	}

	return bestWord.Meaning
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
