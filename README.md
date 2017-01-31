## [Strongly Connected Components](https://en.wikipedia.org/wiki/Strongly_connected_component)

[![GoDoc](https://godoc.org/github.com/chasestarr/scc?status.svg)](https://godoc.org/github.com/chasestarr/scc)

```go
import (
  "fmt"

  "github.com/chasestarr/scc"
  "github.com/chasestarr/scc/graph"
)

func main() {
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
	sscs := scc.Search(g)
	fmt.Println(len(sccs)) // 3
}
