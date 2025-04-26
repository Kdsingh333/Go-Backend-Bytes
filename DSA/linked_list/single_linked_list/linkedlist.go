
// Implemented Single linked list and their respective function as well  
package single_linked_list

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int64
}


func InitList[T any](data T) *List[T] {
	TempNode := &Node[T]{Data: data, Next: nil}
	return &List[T]{Head: TempNode, Size: 1, Tail: TempNode}
}

func (l *List[T]) InsertNodeAtBeg(data T) {
	//  var Temp *Node
	if l == nil || l.Head == nil {
		l.Head = &Node[T]{Data: data, Next: nil}
		l.Tail = l.Head
		l.Size = 1
		return
	}
	Head := l.Head
	tempNode := &Node[T]{Data: data, Next: nil}
	l.Head = tempNode
	l.Head.Next = Head
	l.Size++
}

func (l *List[T]) InsertNodeAtlast(data T) {
	//  var Temp *Node
	if l == nil || l.Head == nil {
		l.Head = &Node[T]{Data: data, Next: nil}
		l.Tail = l.Head
		l.Size = 1
		return
	}
	l.Tail.Next = &Node[T]{Data: data, Next: nil}
	l.Tail = l.Tail.Next
	l.Size++
}

func (l *List[T]) DeleteNodeBeg() {
	if l == nil {
		fmt.Println("linked list is under flow")
		return
	}
	current := l.Head.Next

	// l.Head = nil
	l.Head = current
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
	current := l.Head
	for current.Next != l.Tail {
		current = current.Next
	}
	current.Next = nil
	l.Tail = current
	l.Size--
}

func (l *List[T]) PrintList() {
	Head := l.Head
	for Head.Next != nil {
		fmt.Print(Head.Data, "<-")
		Head = Head.Next
	}
	fmt.Print(Head.Data)
	fmt.Println()
}

func (l *List[T]) GetSize() int64 {
	return l.Size
}

func DemoSlinkedList() {
	// LL := &List[string]{}  // Both are the same this is using direct struct
	LL := InitList("Kuldeep") // this some how we can say that mimicing the constructor func
	LL.InsertNodeAtBeg("Pushp")
	fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertNodeAtlast("Ram")
	fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertNodeAtBeg("Shyam")
	fmt.Println("Current size of list : ", LL.GetSize())
	LL.InsertNodeAtlast("Gagan")
	fmt.Println("Current size of list : ", LL.GetSize())
	LL.PrintList()
	LL.DeleteNodeBeg()
	LL.PrintList()
	LL.DeleteNodeLast()
	LL.PrintList()
	fmt.Println("Current size of list : ", LL.GetSize())
}
