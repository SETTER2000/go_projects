// Большие числа
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{"  1   ", " 11   ", "  1   ", "  1   ", "  1   ", "  1   ", " 111  "},
	{" 2222 ", "2    2", "     2", "    2 ", " 2    ", "2     ", "222222"},
	{" 3333 ", "3    3", "     3", "   33 ", "     3", "3    3", " 3333 "},
	{"4    4", "4    4", "4    4", "444444", "     4", "     4", "     4"},
	//{"    4 ", "   44 ", "  4 4 ", " 4  4 ", "444444", "    4 ", "    4 "},
	{"55555 ", "5     ", "5     ", "55555 ", "     5", "     5", "55555 "},
	{"    6 ", "   6  ", "  6   ", " 6666 ", "6    6", "6    6", " 6666 "},
	{"777777", "     7", "     7", "    7 ", "   7  ", "  7   ", " 7    "},
	{" 8888 ", "8    8", "8    8", " 8888 ", "8    8", "8    8", " 8888 "},
	{" 9999 ", "9    9", "9    9", " 9999 ", "   9  ", "  9   ", " 9    "},
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}
