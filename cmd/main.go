package main

import (
	"fmt"
	"scrabble"
	"scrabble/finder"
)

func main() {
	wordsFinder, _ := finder.NewWordsFinder()

	fmt.Println(wordsFinder.FindWord([]rune("asdb"), "able"))
	fmt.Println(wordsFinder.Put(map[scrabble.Cord]rune{
		{0, 0}: 'a', {1, 0}: 'b', {2, 0}: 'l', {3, 0}: 'e',
		{4, 0}: 't', {5, 0}: 'o', {6, 0}: 'n', {7, 0}: 'e',
	}))
	fmt.Println(wordsFinder.Put(map[scrabble.Cord]rune{
		{0, 0}: 'a', {1, 0}: 'b', {2, 0}: 'l', {3, 0}: 'e',
		{4, 0}: 't', {5, 0}: 'o', {6, 0}: 'n', {7, 0}: 'e',
	}))

}
