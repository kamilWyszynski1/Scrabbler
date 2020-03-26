package main

import (
	"fmt"
	"scrabble/finder"
)

func main() {
	wordsFinder, _ := finder.NewWordsFinder()

	fmt.Println(wordsFinder.FindWord([]rune("dopexadasger")))

}
