package hashtable

import "fmt"

const ARRAY_SIZE = 7

type HashTable struct {
  array [ARRAY_SIZE]*bucket
}

// Insert
func (h *HashTable) Insert(key string) {
  index := hash(key)
  h.array[index].insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
  index := hash(key)
  return h.array[index].search(key)
}

// Delete
func (h *HashTable) Delete(key string){
  index := hash(key)
  h.array[index].delete(key)
}

type bucket struct {
  head *bucketNode
}

// insert bucket
func (b *bucket) insert(k string) {

  if b.search(k) {
    fmt.Println("Alreadt Exists")
    return
  } 

  newNode := &bucketNode{key : k}
  newNode.next = b.head
  b.head = newNode
}

// search bucket
func (b *bucket) search(k string) bool {
  currentNode := b.head
  for currentNode != nil {
    if currentNode.key == k {
      return true
    }

    currentNode = currentNode.next
  }

  return false
}

func (b *bucket) delete(k string) {

  if b.head.key == k {
    b.head = b.head.next
    return 
  }

  prevNode := b.head
  for prevNode.next != nil {
    if prevNode.next.key == k {
      prevNode.next = prevNode.next.next
    }

    prevNode = prevNode.next
  }
} 

type bucketNode struct {
  key string 
  next *bucketNode
}

func hash(key string) int{
  sum := 0
  for _, v := range key {
    sum += int(v)
  }

  return sum % ARRAY_SIZE
}

func Init() *HashTable {
  result := &HashTable{}

  for i := range result.array {
    result.array[i] = &bucket{}
  }

  return result
}

func ExecHashTable() {
  hashTable := Init()
  list := []string{
    "ERIC",
    "KENNY",
    "KYLE",
    "STAN",
    "RANDY",
    "BUTTERS",
    "TOKEN",
  }

  for _, v := range list {
    hashTable.Insert(v)
  }

  hashTable.Delete("STAN")
  fmt.Println("STAN", hashTable.Search("STAN"))
  fmt.Println("RANDY", hashTable.Search("RANDY"))
}

