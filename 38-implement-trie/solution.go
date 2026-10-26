package trie

// Trie is a prefix tree storing lowercase English words.
type Trie struct {
	// TODO: add fields (e.g. a root node with children and an end-of-word flag)
}

// NewTrie returns an empty Trie.
func NewTrie() *Trie {
	return &Trie{}
}

// Insert adds word to the trie.
func (t *Trie) Insert(word string) {
	// TODO: implement
}

// Search reports whether word was previously inserted.
func (t *Trie) Search(word string) bool {
	// TODO: implement
	return false
}

// StartsWith reports whether any inserted word has the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	// TODO: implement
	return false
}
