// speller implements a spell-checker
package main

import (
	"github.com/MatthieuTran/cs50-go/speller/cmd/speller/dictionary"
	"io"
	"log"
	"os"
	"unicode"
)

// Default dictionary
const DICTIONARY = "./dictionaries/large"

func main() {
	argc := len(os.Args)
	// Check for correct number of args
	if argc != 2 && argc != 3 {
		log.Fatalln("Usage: ./speller [dictionaryPath] text")
	}

	var dictionaryPath string
	// Determine dictionaryPath to use
	if argc == 3 {
		dictionaryPath = os.Args[1]
	} else {
		dictionaryPath = DICTIONARY
	}

	// Load Dictionary
	dict := dictionary.New()
	if err := dict.Load(dictionaryPath); err != nil {
		// Exit if dictionary is not loaded
		log.Fatalf("Could not load %s: %s\n", dictionaryPath, err)
	}

	// Determine text to use
	var text string
	if argc == 3 {
		text = os.Args[2]
	} else {
		text = os.Args[1]
	}

	// Try to open text
	file, err := os.OpenFile(text, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalf("Could not open %s: %s\n", text, err)
	}
	defer file.Close()

	// Prepare to report misspellings
	log.Print("\nMISSPELLED WORDS\n\n")

	// Prepare to spell-check
	var (
		index
		misspellings
		words int
		word string
	)

	// Another way to do this would be to use the scanner, to also capture unicode
	buf := make([]byte, 1)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error reading text: %s\n", err)
		}

		c := rune(buf[0])
		// Allow only alphabetical characters and apostrophes
		if unicode.IsLetter(c) || (c == rune('\'') && index > 0) {
			// Append character to word
			word += string(c)
			index++

			// Ignore alphabetical strings too long to be words
			if index > dictionary.MAX_LENGTH {
				// Consume remainder of alphabetical string
				for
			}
		} else if unicode.IsDigit(c) { // Ignore words with numbers (like MS Word can)

		} else if index > 0 { // We must have found a whole word

		}
	}
}