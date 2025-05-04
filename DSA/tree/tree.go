package tree

import (
	"fmt"
	"strings"

	"github.com/Kdsingh333/Go-Backend-Bytes/DSA/generic"
)

type Tree[T generic.Ordered] struct {
	Data   T
	Parent *Tree[T]
	Child  []*Tree[T]
}

type ITree[T generic.Ordered] interface {
	SetData(data T)
	GetParent() *Tree[T]
	GetChildren() []*Tree[T]
	AddChild(child *Tree[T])
	RemoveChild(data T) bool
	Find(data T) *Tree[T]
	IsLeaf() bool
	Depth() int
	Display(level int)
}

func InitTree[T generic.Ordered](Data T) *Tree[T] {
	return &Tree[T]{Data: Data}
}

func ConstructTree[T generic.Ordered]() ITree[T] {
	return &Tree[T]{}
}

func (t *Tree[T]) SetData(data T) {
	t.Data = data
}

func (t *Tree[T]) GetParent() *Tree[T] {
	return t.Parent
}

func (t *Tree[T]) GetChildren() []*Tree[T] {
	return t.Child
}

func (t *Tree[T]) AddChild(child *Tree[T]) {
	t.Child = append(t.Child, child)
}

func (t *Tree[T]) RemoveChild(data T) bool {
	for i, val := range t.Child {
		if val.Data == data {
			t.Child = append(t.Child[:i], t.Child[i+1:]...)
		}
	}
	return false
}

func (t *Tree[T]) Find(data T) *Tree[T] {
	if t.Data == data {
		return t
	}

	for _, val := range t.Child {
		found := val.Find(data)
		if found != nil {
			return found
		}
	}

	return nil
}

func (t *Tree[T]) IsLeaf() bool {
	return t.Child == nil
}

func (t *Tree[T]) Depth() int {
	if t == nil || t.IsLeaf() {
		return 0
	}

	maxDepth := 0
	for _, child := range t.Child {
		childDepth := child.Depth()
		if childDepth > maxDepth {
			maxDepth = childDepth
		}
	}

	return maxDepth + 1
}

func (t *Tree[T]) Display(level int) {
	if t == nil {
		return
	}

	fmt.Printf("%s%v\n", strings.Repeat(" ", level*3)+"|_", t.Data)

	for _, child := range t.Child {
		child.Display(level + 1)
	}
}
func StartDemoTree() {
	World := InitTree("World")
	India := InitTree("India")
	America := InitTree("America")
	Uk := InitTree("Uk")

	World.AddChild(India)
	World.AddChild(America)
	World.AddChild(Uk)

	Delhi := InitTree("Delhi")
	Mumbai := InitTree("Mumbai")
	Bangalore := InitTree("Bangalore")
	India.AddChild(Delhi)
	India.AddChild(Mumbai)
	India.AddChild(Bangalore)

	NewYork := InitTree("New York")
	LosAngeles := InitTree("Los Angeles")
	Chicago := InitTree("Chicago")
	America.AddChild(NewYork)
	America.AddChild(LosAngeles)
	America.AddChild(Chicago)

	London := InitTree("London")
	Manchester := InitTree("Manchester")
	Birmingham := InitTree("Birmingham")
	Uk.AddChild(London)
	Uk.AddChild(Manchester)
	Uk.AddChild(Birmingham)

	fmt.Println("Tree Structure:")
	World.Display(0)

	fmt.Println("\nIs Root a leaf node?", World.IsLeaf())
	fmt.Println("Depth of the tree:", World.Depth())

	fmt.Println("\nFinding 'Chicago':")
	found := World.Find("Chicago")
	if found != nil {
		fmt.Println("Found node with data:", found.Data)
	} else {
		fmt.Println("Node not found")
	}

	fmt.Println("\nRemoving 'Los Angeles':")
	America.RemoveChild("Los Angeles")
	World.Display(0)

	fmt.Println("\nRemoving 'India':")
	World.RemoveChild("India")
	World.Display(0)
}
