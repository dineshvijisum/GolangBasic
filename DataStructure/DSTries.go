package main

import "fmt"

//Alphabet is the number of possible characters in the tries
const AlphabetSize = 26

//Node represents each node in the tries
type Node struct {
	children [26]*Node
	isEnd    bool
}

//Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

//InitTrie will create new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

//Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

//search will take in a word and return true is that wort is included in the
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	fmt.Println(testTrie.root)
	myTrie := InitTrie()
	toAdd := []string{
		"aragon",
		"Dinesh",
		"kumar",
		"fake",
		"yes",
		"similar",
	}
	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.Search("Dinesh"))
}
