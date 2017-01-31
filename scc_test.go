package scc

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"io/ioutil"

	"github.com/chasestarr/scc/graph"
)

func TestSimple(t *testing.T) {
	g := graph.New()
	g.AddNode(1)
	g.AddNode(2)
	g.AddNode(3)
	g.AddNode(4)
	g.AddNode(5)
	g.AddEdge(2, 1)
	g.AddEdge(1, 3)
	g.AddEdge(3, 2)
	g.AddEdge(1, 4)
	g.AddEdge(4, 5)
	sccs := Search(g)
	correct := 3
	if len(sccs) != correct {
		t.Fatalf("incorrect scc count, expected: %d, received: %d", sccs, correct)
	}
}

func readFile(name string) [][]int {
	input, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal("error reading file")
	}
	s := strings.Split(string(input), "\n")
	output := make([][]int, len(s))
	for i, line := range s {
		items := strings.Split(string(line), " ")
		lineItems := make([]int, len(items))
		for j, item := range items {
			x, _ := strconv.Atoi(item)
			lineItems[j] = x
		}
		output[i] = lineItems
	}
	return output
}

func buildGraph(input [][]int) *graph.Graph {
	g := graph.New()

	for _, line := range input {
		head := line[0]
		tail := line[1]
		g.AddNode(head)
		g.AddNode(tail)
		g.AddEdge(head, tail)
	}

	return g
}

func TestLarge(t *testing.T) {
	input := readFile("test_input.txt")
	g := buildGraph(input)

	count := Search(g)
	fmt.Println(len(count))
}
