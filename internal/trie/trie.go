package trie

type Node struct {
	letters [26]*Node
	isKey   bool
	data    []uint8
}

type Trie struct {
	root Node
}
