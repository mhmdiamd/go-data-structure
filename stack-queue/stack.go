package stackqueue

import "fmt"

type Stack struct {
  items []int
}

func (s *Stack) Push(item int) {
  s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
  l := len(s.items) - 1

  toRemove := s.items[l]
  s.items = s.items[:l]

  fmt.Println(toRemove)
  fmt.Println(s.items)

  return toRemove
}

func ExecStack() {
  myItems := Stack{}
  myItems.Push(10)
  myItems.Push(20)
  myItems.Push(30)
  myItems.Push(40)

  fmt.Println(myItems)
  myItems.Pop()
  fmt.Println(myItems)

}


