// Dictionary defines and implements the dictionary struct
package dictionary

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Represents number of buckets in a hash table
const N = 26

type Dictionary struct {
	nodes []Node
}

// Create a new dictionary object
func New() *Dictionary {
	nodes := make([]Node, N, N)

	return &Dictionary{nodes: nodes}
}

// Hashes word to a number between 0 and 25, inclusive, based on its first letter
func (d *Dictionary) Hash(word string) byte {
	// Note: We are making the assumption that we will only get inputs in ASCII and in the alphabet
	return strings.ToLower(word)[0] - 'a'
}

// Loads dictionary into memory, returning true if successful else false
func (d *Dictionary) Load(dictionary string) error {
	// Open Dictionary
	file, err := os.OpenFile(dictionary, os.O_RDONLY, 0444)
	if err != nil {
		return err
	}
	defer file.Close() // Close dictionary at end of function

	// Buffer for a word (max size: MAX_LENGTH + 1)
	var word string
	n := 1

	// Insert words into hash table
	for n > 0 {
		n, err = fmt.Fscanf(file, "%s", &word)
		if err == io.EOF { // Return when dictionary is fully loaded
			return nil
		}
		if err != nil {
			return err
		}

		// Prepare node for appending to array
		node := NewNode(word)
		i := d.Hash(word)

		// Add word to dictionary.nodes
		var curr *Node
		curr = &(d.nodes[i])

		// If there is no existing entry, set as start
		if curr == nil {
			curr = node

			continue
		}

		// Else, chain the node onto existing ones
		for curr != nil {
			// Chain until last node
			if curr.Next == nil {
				// Set the next node to the newly made one
				curr.Next = node
				break
			}

			curr = curr.Next
		}
	}

	// Indicate Success
	return nil
}

// Returns number of words in dictionary if loaded else 0 if not yet loaded
func (d *Dictionary) Size() int {
	var size int
	for i := 0; i < N; i++ {
		if d.nodes[i] == (Node{}) {
			continue
		}

		var curr *Node
		curr = &(d.nodes[i])
		for curr != nil {
			if len(curr.Word) > 0 {
				size++
			}
			curr = curr.Next
		}
	}

	return size
}

// Returns true if word is in dictionary, else false
func (d *Dictionary) Check(word string) bool {
	// Make word to lower at start to reduce processing time
	word = strings.ToLower(word)

	// Get index by hashing word
	i := d.Hash(word)

	curr := d.nodes[i]
	for curr != (Node{}) {
		if strings.ToLower(curr.Word) == word {
			// Found match
			return true
		}

		if curr.Next == nil {
			return false
		}

		// Match not found yet, move onto next node
		curr = *curr.Next
	}

	return false

}

// Unloads dictionary from memory, returning true if successful else false
func (d *Dictionary) Unload() error {
	d.nodes = nil

	return nil
}