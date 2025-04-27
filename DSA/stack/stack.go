package stack

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)
// In this tutorial I have tried comparable , any , constraints.ordered interface as I have to add 
// equality check  and do sum of the data element so the suitable interface ordered one

// Function	Purpose
// push(item)	     Add an item to the top of the stack
// pop()	         Remove and return the item from the top
// peek() or top()	 View the top item without removing
// isEmpty()	     Check if the stack is empty
// size()	         Get the number of items in the stack
 
type Stack[T constraints.Ordered] struct {
	Array []T
}

type IStack[T constraints.Ordered] interface {
	Push(item T)
	Pop() (T, error)
	Peek() (T, error)
	IsEmpty() bool
	Size() int
	Clear()
	Contains(item T) bool
	Sum() T
}

func InitStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.Array = append(s.Array, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	item := s.Array[len(s.Array)-1]
	s.Array = s.Array[:len(s.Array)-1]
	return item, nil
}


func (s *Stack[T]) Clear() {
	s.Array = nil
}

// Contains implements IStack.
func (s *Stack[T]) Contains(item T) bool {
	for _, val := range s.Array {
		if val == item {
			return true
		}

	}
	return false
}


func (s *Stack[T]) IsEmpty() bool {
	return len(s.Array) == 0
}


func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}
	return s.Array[len(s.Array)-1], nil
}

func (s *Stack[T]) Size() int {
	return len(s.Array)
}

func (s *Stack[T]) Sum() T {
	var sum T
	for _, v := range s.Array {
		sum += v
	}
	return sum
}

func DemoStack() {
	s := InitStack[string]()
	s.Push("k")
	s.Push("U")
	s.Push("L")
	s.Push("D")
	s.Push("E")
	s.Push("E")
	s.Push("P")
	fmt.Println("Sum of all the values :",s.Sum())
	item, err := s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item)

	item, err = s.Peek()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item)

	fmt.Println(s.IsEmpty())
	fmt.Println(s.Size())

	
	s.Clear()
	fmt.Println(s.IsEmpty())
	item, err = s.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(item)

}
