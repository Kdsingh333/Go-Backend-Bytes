package graph

import (
	"fmt"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a node with key k
func (g *Graph) AddVertex(k int) {
	if Contains(g.vertices, k) {
		fmt.Printf("vertex %v not added because it is an existing key\n", k)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds a directed edge from src to dest
func (g *Graph) AddEdge(src, dest int) {
	srcVertex := getVertex(g.vertices, src)
	destVertex := getVertex(g.vertices, dest)

	if srcVertex == nil || destVertex == nil {
		fmt.Printf("invalid edge from %v to %v\n", src, dest)
		return
	}

	if !isAdjacent(srcVertex, destVertex) {
		srcVertex.adjacent = append(srcVertex.adjacent, destVertex)
	}
}

// RemoveEdge removes an edge from src to dest
func (g *Graph) RemoveEdge(src, dest int) {
	srcVertex := getVertex(g.vertices, src)
	destVertex := getVertex(g.vertices, dest)

	if srcVertex == nil || destVertex == nil {
		fmt.Printf("invalid edge from %v to %v\n", src, dest)
		return
	}

	newAdj := []*Vertex{}
	for _, v := range srcVertex.adjacent {
		if v.key != dest {
			newAdj = append(newAdj, v)
		}
	}
	srcVertex.adjacent = newAdj
}

// RemoveNode deletes a vertex and all edges to/from it
func (g *Graph) RemoveNode(k int) {
	var updated []*Vertex
	for _, v := range g.vertices {
		if v.key != k {
			updated = append(updated, v)
		}
	}
	g.vertices = updated

	// Remove edges pointing to the deleted vertex
	for _, v := range g.vertices {
		var newAdj []*Vertex
		for _, adj := range v.adjacent {
			if adj.key != k {
				newAdj = append(newAdj, adj)
			}
		}
		v.adjacent = newAdj
	}
}

// Print displays the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v:", v.key)
		for _, adj := range v.adjacent {
			fmt.Printf(" %v", adj.key)
		}
	}
	fmt.Println()
}

// HasCycle uses DFS and visited sets to check for cycles in a directed graph
func (g *Graph) HasCycle() bool {
	visited := map[int]bool{}
	recStack := map[int]bool{}

	for _, v := range g.vertices {
		if hasCycleDFS(v, visited, recStack) {
			return true
		}
	}
	return false
}

func hasCycleDFS(v *Vertex, visited, recStack map[int]bool) bool {
	if recStack[v.key] {
		return true
	}
	if visited[v.key] {
		return false
	}

	visited[v.key] = true
	recStack[v.key] = true

	for _, neighbor := range v.adjacent {
		if hasCycleDFS(neighbor, visited, recStack) {
			return true
		}
	}
	recStack[v.key] = false
	return false
}

// Helpers

func Contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if v.key == k {
			return true
		}
	}
	return false
}

func getVertex(s []*Vertex, k int) *Vertex {
	for _, v := range s {
		if v.key == k {
			return v
		}
	}
	return nil
}

func isAdjacent(v *Vertex, n *Vertex) bool {
	for _, adj := range v.adjacent {
		if adj.key == n.key {
			return true
		}
	}
	return false
}

// StartDemoGraph shows usage
func StartDemoGraph() {
	test := &Graph{}

	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(0, 1)
	test.AddEdge(1, 2)
	test.AddEdge(2, 3)
	test.AddEdge(3, 1) // cycle

	test.Print()

	fmt.Println("Has Cycle:", test.HasCycle())

	test.RemoveEdge(3, 1)
	fmt.Println("\nAfter removing edge 3->1:")
	test.Print()
	fmt.Println("Has Cycle:", test.HasCycle())

	test.RemoveNode(2)
	fmt.Println("\nAfter removing node 2:")
	test.Print()
}
