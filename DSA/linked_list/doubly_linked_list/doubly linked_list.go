package doubly_linked_list

import "fmt"

type Node[T any] struct {
	Previous *Node[T]
	Data     T
	Next     *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int64
}

func InitList[T any](data T) *List[T] {
	TempNode := &Node[T]{Data: data, Next: nil, Previous: nil}
	return &List[T]{Head: TempNode, Size: 1, Tail: TempNode}
}

func (l *List[T]) InsertAtHead(data T) {
	//  var Temp *Node
	if l == nil || l.Head == nil {
		l.Head = &Node[T]{Data: data,Previous: nil, Next: nil}
		l.Tail = l.Head
		l.Size = 1
		return
	}
	Head := l.Head
	tempNode := &Node[T]{Data: data, Previous: nil,Next: Head}
	l.Head = tempNode
	l.Head.Next = Head
	l.Size++
}

func (l *List[T]) InsertAtTail(data T) {
	//  var Temp *Node
	if l == nil || l.Head == nil {
		l.Head = &Node[T]{Data: data,Previous: nil, Next: nil}
		l.Tail = l.Head
		l.Size = 1
		return
	}
	l.Tail.Next = &Node[T]{Data: data, Previous: l.Tail,Next: nil}
	l.Tail = l.Tail.Next
	l.Size++
}

func (l *List[T]) DeleteAtHead() {
	if l == nil {
		fmt.Println("linked list is under flow")
		return
	}
	current := l.Head.Next

	// l.Head = nil
	l.Head = current
	l.Head.Previous = nil
	l.Size--

	if l.Size == 0 {
		l.Tail = nil
	}
}

func (l *List[T]) DeleteNodeLast() {
	if l == nil {
		fmt.Println("linked list is under flow")
		return
	}
	
	
	current := l.Tail.Previous
	current.Next = nil
	l.Tail = current
	l.Size--
}

func (l *List[T]) PrintListForward() {
	Head := l.Head
	for Head.Next != nil {
		fmt.Print(Head.Data, "<-")
		Head = Head.Next
	}
	fmt.Print(Head.Data)
	fmt.Println()
}
func (l *List[T]) PrintListBackward() {
	Head := l.Tail
	for Head.Previous!= nil {
		fmt.Print(Head.Data, "<-")
		Head = Head.Previous
	}
	fmt.Print(Head.Data)
	fmt.Println()
}

func (l *List[T]) GetSize() int64 {
	return l.Size
}

func DemoDoublylinkedList() {
	// LL := &List[string]{}  // Both are the same this is using direct struct
	LL := InitList("Kuldeep") // this some how we can say that mimicing the constructor func
	LL.InsertAtHead("Pushp")
	// fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertAtTail("Ram")
	// fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertAtTail("Shyam")
	// fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertAtTail("Gagan")
	// fmt.Println("Current size of list : ", LL.GetSize())
	LL.PrintListForward()
	LL.DeleteAtHead()
	LL.PrintListForward()
	LL.DeleteNodeLast()
	LL.PrintListForward()
	fmt.Println("Current size of list : ", LL.GetSize())
	LL.PrintListBackward()
}
