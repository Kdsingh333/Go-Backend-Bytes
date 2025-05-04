package queue

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

// enqueue(x)	         Add an element x to the rear (end) of the queue.
// dequeue()	         Remove and return the front element from the queue.
// peek() / front()     View (but not remove) the front element.
// isEmpty()	         Check whether the queue has no elements.
// size()	             Get the number of elements in the queue.

type Queue[T constraints.Ordered] struct {
	Array []T
}

// data types which Can be use to construct queue
// type Ordered interface {
// 	~int | ~int8 | ~int16 | ~int32 | ~int64 |
// 		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
// 		~float32 | ~float64 |
// 		~string
// }

type IQueue[T constraints.Ordered] interface {
	Enqueue(item T)
	Dequeue() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	Display()
	Clear()
	Contains(val T) bool
	Equals(other *Queue[T]) bool
}

func ConstructQueue[T constraints.Ordered]() IQueue[T] {
	return &Queue[T]{}
}

func InitQueue[T constraints.Ordered]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(data T) {
	q.Array = append(q.Array, data)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var element T
	if q.IsEmpty() {
		fmt.Printf("Queue is underflow")
		return element, errors.New("Queue is underflow")
	}
	element = q.Array[0]
	q.Array = q.Array[1:]
	return element, nil
}
func (q *Queue[T]) Peek() (T, error) {
	var element T
	if q.IsEmpty() {
		return element, errors.New("Queue is underflow")
	}
	return q.Array[0], nil
}
func (q *Queue[T]) IsEmpty() bool {
	return len(q.Array) == 0
}
func (q *Queue[T]) Size() int {
	return len(q.Array)
}
func (q *Queue[T]) Clear() {
	q.Array = nil
}

func (q *Queue[T]) Contains(val T) bool {
	for _, v := range q.Array {
		if v == val {
			return true
		}
	}
	return false
}

func (q *Queue[T]) Display() {
	if q.IsEmpty() {
		fmt.Print("queue is underflow")
	}
	for _, val := range q.Array {
		fmt.Print(val)
	}

	fmt.Println()
}

func (q *Queue[T]) Equals(other *Queue[T]) bool {
	if q.Size() != other.Size() {
		return false
	}

	for i, val := range q.Array {
		if val != other.Array[i] {
			return false
		}
	}

	return true
}

func StartDemoQueue() {
	q := ConstructQueue[string]()

	q.Enqueue("K")
	q.Enqueue("U")
	q.Enqueue("L")
	q.Dequeue()
	// val, _ := q.Peek()
	fmt.Println(q.Peek())

	fmt.Println(q.IsEmpty())
	q.Enqueue("D")
	q.Enqueue("E")
	q.Enqueue("E")
	q.Enqueue("P")
	q.Display()
	q.Clear()
	q.Display()
}
