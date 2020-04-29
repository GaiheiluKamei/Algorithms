package bfs

// Node represents a node in a graph.
type Node struct {
	Name string
	MangoSeller bool
}

func (n *Node) String() string {
	return n.Name
}

// Graph represents a graph.
type Graph struct {
	Edges map[*Node][]*Node
}

// AddEdge add an edge to the graph.
func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.Edges == nil {
		g.Edges = make(map[*Node][]*Node)
	}
	
	g.Edges[n1] = append(g.Edges[n1], n2)
}

func (g *Graph) BFS(start *Node) *Node {
	if start.MangoSeller {
		return start
	}

	// initialize queue
	var queue []*Node
	for _, n := range g.Edges[start] {
		queue = append(queue, n)
	}

	searched := make(map[*Node]struct{})

	for len(queue) != 0 {

		// pop node from queue
		var n *Node
		n, queue = queue[0], queue[1:]

		// check if node was already searched
		_, alreadyChecked := searched[n]
		if alreadyChecked {
			continue
		} else {
			searched[n] = struct{}{}
		}

		if n.MangoSeller {
			return n
		}

		for _, v := range g.Edges[n] {
			 queue = append(queue, v)
		}
	}
	
	return nil
}
