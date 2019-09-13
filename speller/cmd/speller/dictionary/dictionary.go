// Dictionary defines and implements the dictionary struct
package dictionary

import (
	"fmt"
	"os"
	"strings"
)

// Represents number of buckets in a hash table
const N = 26

type Dictionary struct {
	nodes *[N]Node
}

// Create a new dictionary object
func New() *Dictionary {
	return &Dictionary{}
}

// Hashes word to a number between 0 and 25, inclusive, based on its first letter
func (d *Dictionary) Hash(word string) byte {
	// Note: We are making the assumption that we will only get inputs in ASCII and in the alphabet
	return strings.ToLower(word)[0] - 'a'
}

// Loads dictionary into memory, returning true if successful else false
func (d *Dictionary) Load(dictionary string) error {
	// Open Dictionary
	file, err := os.OpenFile("dictionary", os.O_RDONLY, 0444)
	if err != nil {
		return err
	}
	defer file.Close() // Close dictionary at end of function

	// Buffer for a word (max size: MAX_LENGTH + 1)
	var word string
	var n int

	// Insert words into hash table
	for n > 0 {
		n, err = fmt.Fscanf(file, "%s", &word)
		if err != nil {
			return err
		}

		// TODO
	}

	// Indicate Success
	return nil
}

// Returns number of words in dictionary if loaded else 0 if not yet loaded
func (d *Dictionary) Size() int {
	// TODO
	return 0
}

// Returns true if word is in dictionary else false
func (d *Dictionary) Check(word string) error {
	// TODO
	return nil
}

// Unloads dictionary from memory, returning true if successful else false
func (d *Dictionary) Unload() error {
	// TODO
	return nil
}
