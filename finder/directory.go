package finder

import (
	"bufio"
	"fmt"
	"os"
	"scrabble"
)

func (w *Words) LoadWordsDirectory() error {
	logger := w.log.WithField("method", "LoadWordsDirectory")
	f, err := os.Open("/home/kamil/go/src/scrabble/finder/words.txt")
	if err != nil {
		return fmt.Errorf("failed to open word.txt")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := scanner.Text()
		histogram := wordToHistogram(word)
		if histogram == nil {
			continue
		}
		w.Dictionary = append(w.Dictionary, scrabble.Word{
			Meaning:   word,
			Histogram: histogram,
			Value:     wordToValue(word),
		})
	}

	if err := scanner.Err(); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func wordToHistogram(word string) map[rune]int {
	histogram := make(map[rune]int)
	for _, letter := range word {
		if letter < 'a' || letter > 'z' {
			return nil
		}
		histogram[letter] += 1
	}
	return histogram
}

func wordToValue(word string) int {
	var value int
	for _, letter := range word {
		value += scrabble.LetterValue[letter]
	}
	return value
}
