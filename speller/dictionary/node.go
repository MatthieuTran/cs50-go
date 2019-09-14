// Node defines the node struct
package dictionary

// Maximum length for a word (e.g., pneumonoultramicroscopicsilicovolcanoconiosis)
const MAX_LENGTH = 45

// Represents a node in a hash table
type Node struct {
	Word string
	Next *Node
}

func NewNode(word string) *Node {
	return &Node{Word: word, Next: nil}
}

