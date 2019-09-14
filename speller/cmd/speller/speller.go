// speller implements a spell-checker
package main

import (
	"bufio"
	"fmt"
	"github.com/MatthieuTran/cs50-go/speller/dictionary"
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
	fmt.Print("\nMISSPELLED WORDS\n\n")

	// Prepare to spell-check
	var (
		index int
		misspellings int
		words int
		word  string
	)
	reader := bufio.NewReader(file)

	// Spell-check each word in text
	for {
		c, err := reader.ReadByte()
		if err == io.EOF {
			break
		}

		if err != nil { // error reading file
			log.Fatalf("Error reading %s: %s\n", text, err)
		}

		// Allow only alphabetical characters and apostrophes
		if unicode.IsLetter(rune(c)) || (c == '\'' && index > 0) {
			// Append character to word
			word += string(c)
			index++

			// Ignore alphabetical strings too long to be words
			if index > dictionary.MAX_LENGTH {
				// Consume remainder of alphabetical string
				for {
					c, err = reader.ReadByte()

					if err == io.EOF || unicode.IsLetter(rune(c)) {
						break
					}

					if err != nil {
						log.Fatal(err)
					}

					// Prepare for new word
					index = 0
				}
			}
			} else if unicode.IsDigit(rune(c)) { // Ignore words with numbers (like MS Word)
				// Consume remainder of alphanumeric string
				for err != io.EOF && unicode.IsLetter(rune(c)) || unicode.IsNumber(rune(c)) {
					c, err = reader.ReadByte()

					if err != nil {
						log.Fatal(err)
					}
				}

				// Prepare for new word
				index = 0
			} else if index > 0 { // We must have found a whole word
				// Update counter
				words++

				// Check word's spelling
				success := dict.Check(word)
				if !success {
					fmt.Println(word)
					misspellings++
				}

				// Prepare for next word
				word = ""
				index = 0
			}
		}

	// Determine dictionary's size
	n := dict.Size()

	// Unload dictionary
	err = dict.Unload()
	if err != nil { // Abort if dictionary not unloaded
		log.Fatalf("Could not unload %s:%s\n", dictionaryPath, err)
	}

	// Report Results
	fmt.Printf("\nWORDS MISSPELLED:     %d\n", misspellings)
	fmt.Printf("WORDS IN DICTIONARY:  %d\n", n)
	fmt.Printf("WORDS IN TEXT:        %d\n", words)
}