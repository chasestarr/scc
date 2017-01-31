package graph

// Graph keeps track of all nodes within the data structure
type Graph struct {
	Nodes map[int]*Node
}

// Node is a vertex element within the graph
type Node struct {
	ID       int
	Edges    []Edge
	Reversed []Edge
	Seed     int
}

// Edge is an edge from source to target
type Edge struct {
	Tgt *Node
}

// New returns a pointer to a new Graph struct
func New() *Graph {
	return &Graph{}
}

// AddNode adds new node struct to graph, takes in id int
func (g *Graph) AddNode(id int) *Node {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*Node)
	}

	if _, ok := g.Nodes[id]; ok {
		return nil
	}

	node := &Node{ID: id}
	g.Nodes[id] = node
	return node
}

// AddEdge adds a new edge from source to target
func (g *Graph) AddEdge(src, tgt int) *Edge {
	e := Edge{Tgt: g.Nodes[tgt]}
	g.Nodes[src].Edges = append(g.Nodes[src].Edges, e)
	// add revered edge for scc alg
	g.Nodes[tgt].Reversed = append(g.Nodes[tgt].Reversed, Edge{Tgt: g.Nodes[src]})
	return &e
}
