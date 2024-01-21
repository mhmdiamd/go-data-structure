package main

import "fmt"

const AllphabetSize = 26

// Node represent each node in the trie
type Node struct {
  children [AllphabetSize]*Node
  isEnd bool
}

type Trie struct {
  root *Node
}

func InitTrie() *Trie {
  result := &Trie{
    root : &Node{},
  }

  return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
  wordLength := len(w)

  // fmt.Println(wordLength)
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

// Search will take in a word and RETURN true is that word is included in the Node
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
  myTrie := InitTrie()
  alphabetSlice := []string{
    "aragorn",
    "aragon",
    "argon",
    "eragon",
    "oregon",
    "oregano",
    "oreo",
  }

  for _, v := range alphabetSlice {
    myTrie.Insert(v)
  }

  fmt.Println(myTrie)
  fmt.Println(myTrie.Search("orc"))
}








