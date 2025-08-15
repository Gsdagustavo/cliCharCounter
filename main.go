package main

import (
	"flag"
	"fmt"
)

func main() {
	word := flag.String("word", "default", "word that will be queried")
	letter := flag.String("letter", "default", "letter that will be searched for in the word")

	flag.Parse()

	count := countLetter(*word, []rune(*letter)[0])

	fmt.Println(count)
}

func countLetter(str string, letter rune) int {
	count := 0

	for _, c := range str {
		if c == letter {
			count++
		}
	}

	return count
}
