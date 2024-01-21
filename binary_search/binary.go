package binarysearch

import "fmt"

var count = 0

// Node represents the components of a binary tree
type Node struct {
  Key int 
  Left *Node
  Right *Node
}

// Insert will add a note to the tree
// The key to add should not be already in 
func (n *Node) Insert(k int) {
  if n.Key < k {
    if n.Right == nil {
      n.Right = &Node{ Key: k }
    } else {
      n.Right.Insert(k)
    }
  }else if n.Key > k{
    if n.Left == nil {
      n.Left = &Node{Key : k}
    }else {
      n.Left.Insert(k)
    }
  }
}

// Search
func (n *Node) Search(k int) bool {

  count++

  if n == nil {
    return false
  }

  if n.Key < k {
    return n.Right.Search(k)
  }else if n.Key > k {
    return n.Left.Search(k)
  }

  return true
}

func ExecBinaryTree() {
  tree := &Node{ Key: 100 }

  fmt.Println(tree)

  tree.Insert(200)
  tree.Insert(99)
  tree.Insert(52)
  tree.Insert(203)
  tree.Insert(19)
  tree.Insert(76)
  tree.Insert(150)
  tree.Insert(310)
  tree.Insert(7)
  tree.Insert(24)
  tree.Insert(88)
  tree.Insert(276)

  fmt.Println(tree)

  fmt.Println(tree.Search(203))

  fmt.Println(count)
}






