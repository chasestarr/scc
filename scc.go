package scc

import (
	"math/rand"
	"time"

	"github.com/chasestarr/scc/graph"
)

// Search graph and return leader nodes for each strongly connected component
func Search(g *graph.Graph) []*graph.Node {
	o := order(g)
	seed := genSeed()
	leaders := []*graph.Node{}
	for i := len(o) - 1; i >= 0; i-- {
		node := g.Nodes[o[i]]
		if !visited(node, seed) {
			dfs(node, seed)
			leaders = append(leaders, node)
		}
	}
	return leaders
}

func dfs(node *graph.Node, seed int) {
	node.Seed = seed
	for i := 0; i < len(node.Edges); i++ {
		child := node.Edges[i].Tgt
		if !visited(child, seed) {
			dfs(child, seed)
		}
	}
}

func visited(node *graph.Node, seed int) bool {
	if seed == node.Seed {
		return true
	}
	return false
}

func genSeed() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	return rand.New(s1).Intn(9999)
}

func order(g *graph.Graph) []int {
	seed := genSeed()
	stack := []int{}
	for i := 1; i < len(g.Nodes)+1; i++ {
		node := g.Nodes[i]
		if !visited(node, seed) {
			fillOrder(node, seed, &stack)
		}
	}
	return stack
}

func fillOrder(node *graph.Node, seed int, stack *[]int) {
	node.Seed = seed
	for i := 0; i < len(node.Reversed); i++ {
		child := node.Reversed[i].Tgt
		if !visited(child, seed) {
			fillOrder(child, seed, stack)
		}
	}
	*stack = append(*stack, node.ID)
}
