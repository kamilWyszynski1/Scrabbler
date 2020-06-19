package common

import (
	"fmt"
	"scrabble"
)

const (
	InfoColor    = "\033[1;34m%s  \033[0m"
	NoticeColor  = "\033[1;36m%s  \033[0m"
	WarningColor = "\033[1;33m%s  \033[0m"
	ErrorColor   = "\033[1;31m%s  \033[0m"
	DebugColor   = "\033[0;36m%s  \033[0m"
)

func PrettyPrintBoard(board map[scrabble.Cord]rune) {
	fmt.Println("-7 -6 -5 -4 -3 -2 -1  0  1  2  3  4  5  6  7")
	for i := -scrabble.Dimension; i <= scrabble.Dimension; i += 1 {
		row := ""
		for j := -scrabble.Dimension; j <= scrabble.Dimension; j += 1 {
			cord := scrabble.Cord{j, i}
			if v, _ := board[cord]; v == 0 {
				row += fmt.Sprintf(InfoColor, "0")
			} else {
				row += fmt.Sprintf(ErrorColor, string(v))
			}
		}
		fmt.Println(row)
	}
}
