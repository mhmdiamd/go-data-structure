package stackqueue

import "fmt"

type Queue struct {
  items []int
}

func (q *Queue) Enqueue(val int) {
  q.items = append(q.items, val)
}

func (q *Queue) Dequeue() int {

  toRemove := q.items[0]
  q.items = q.items[1:]

  return toRemove
}

func ExecQueue() {
  myList := Queue{}

  myList.Enqueue(100)
  myList.Enqueue(200)
  myList.Enqueue(300)
  myList.Enqueue(400)

  fmt.Println(myList)
  myList.Dequeue()
  fmt.Println(myList)

}
