package common

import (
	"bufio"
	"fmt"
	"os"
	"scrabble"
	"strings"
)

func LoadWordsDirectory() ([]scrabble.Word, error) {
	var dict []scrabble.Word
	f, err := os.Open("/home/kamil/go/src/scrabble/common/words.txt")
	if err != nil {
		return nil, fmt.Errorf("failed to open word.txt")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict = append(dict, StringToScrabbleWord(word))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dict, nil
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

func StringToScrabbleWord(str string) scrabble.Word {
	return scrabble.Word{
		Meaning:   str,
		Histogram: wordToHistogram(str),
	}
}
