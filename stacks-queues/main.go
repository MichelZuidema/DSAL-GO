package main

import "fmt"

type Stack struct {
	items []int
}

type Queue struct {
	items []int
}

// Enqueue - This will add a value at the end of the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue - This will remove and return a value at the front of the queue
func (q *Queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[1:]

	return remove
}

// Push - this will add a value at the end of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop - this will remove and return the last inserted value of the stack
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	remove := s.items[l]
	s.items = s.items[:l]

	return remove
}

func main() {
	testQueue()
}

func testQueue() {
	q := Queue{}
	fmt.Println(q)

	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)
	fmt.Println(q)

	q.Dequeue()
	fmt.Println(q)
}

// func testStack() {
// 	s := Stack{}
// 	fmt.Println(s)

// 	s.Push(100)
// 	s.Push(200)
// 	s.Push(300)
// 	fmt.Println(s)

// 	s.Pop()
// 	fmt.Println(s)
// }
