package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//word := flag.String("word", "default", "word that will be queried")
	//letter := flag.String("letter", "default", "letter that will be searched for in the word")

	filePath := flag.String("file", "default", "the path to the file to be queried")

	// parse all flags
	flag.Parse()

	// opens connection with file
	file, err := os.Open(*filePath)

	// error during connection
	if err != nil {
		log.Fatal(err)
	}

	// reads the file contents
	_, err = readFile(file)

	// error during reading
	if err != nil {
		log.Fatal(err)
	}

	_ = file.Close()
}

func readFile(file *os.File) (string, error) {

	bytes, err := io.ReadAll(io.Reader(file))

	if err != nil {
		return "", err
	}

	text := string(bytes)
	fmt.Println("File content: " + text)

	return text, nil
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
