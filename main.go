package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := flag.String("file", "default", "the path to the file to be queried")
	letter := flag.String("letter", "default", "letter that will be searched for in the word")

	// parse all flags
	flag.Parse()

	// opens connection with file
	file, err := os.Open(*filePath)

	// error during connection
	if err != nil {
		log.Fatal(err)
	}

	// reads the file contents
	text, err := readFile(file)

	// error during reading
	if err != nil {
		log.Fatal(err)
	}

	count := countLetter(text, []rune(*letter)[0])

	fmt.Printf("Occurrences of the letter %v in the file: %v", *letter, count)

	_ = file.Close()
}

func readFile(file *os.File) (string, error) {
	bytes, err := io.ReadAll(io.Reader(file))

	if err != nil {
		return "", err
	}

	return string(bytes), nil
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
