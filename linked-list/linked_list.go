package linkedlist

import "fmt"

type Node struct {
  data int
  next *Node
}

type LinkedList struct {
  Head *Node
  Length int
}

func (l *LinkedList) prepend(n *Node) {
  second := l.Head
  l.Head = n
  l.Head.next = second
  
  l.Length++
}

func (l LinkedList) printListData() {
  toPrint := l.Head
  for l.Length != 0 {
    fmt.Printf("%d ", toPrint.data)
    toPrint = toPrint.next
    l.Length--
  }

  fmt.Printf("\n")
}

func (l *LinkedList) deleteWithValue(value int) {

  // Handle if there is no list on linked list 
  if l.Length == 0 {
    return
  }

  // Handle if the value is current Head of linked list itself
  if l.Head.data == value {
    l.Head = l.Head.next
    l.Length--
    return
  }

  previousToDelete := l.Head
  for previousToDelete.next.data != value {
    if previousToDelete.next.next == nil {
      return
    }
    previousToDelete = previousToDelete.next
  }

  previousToDelete.next = previousToDelete.next.next
  l.Length--
}

func ExecLinkedList() {

  myList := LinkedList{}
  node1 := &Node{data : 48}
  node2 := &Node{data : 31}
  node3 := &Node{data : 23}
  node4 := &Node{data : 15}
  node5 := &Node{data : 7}

  myList.prepend(node1)
  myList.prepend(node2)
  myList.prepend(node3)
  myList.prepend(node4)
  myList.prepend(node5)

  myList.printListData()
  myList.deleteWithValue(7)
  myList.printListData()

}
