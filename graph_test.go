package graph

import (
  "fmt"
  "math/rand"
  "testing"
  "time"
)

func TestMain(t *testing.T) {
  t.Run("Base Graph", testBaseGraph(5, "Node", 10))
}

func testBaseGraph(numNodes int, nodeNames string, numEdges int) func(*testing.T) {
  return func(t *testing.T) {
    graph, _, _:= setupGraph(numNodes, nodeNames, numEdges)
    graphNumEdges := len(graph.GetEdges())
    if graphNumEdges != numEdges {
      t.Error(fmt.Sprintf("Expected %v edges not %v.", numEdges, graphNumEdges))
    }
  }
}

func setupGraph(numNodes int, baseNodeName string, numEdges int) (Graph, []Node, []Edge) {
  nodes := make([]Node, 0)
  graph := Graph{ name: "test", nodes: nodes }
  for i := 0; i < numNodes; i++ {
    nodes = append(nodes, Node{ name: fmt.Sprintf("%v %v", baseNodeName, i) })
  }
  pairs := buildRandPairs(numNodes, numEdges)
  edges := make([]Edge, 0)
  for _, pair := range pairs {
    edges = append(edges, Edge{ start: nodes[pair[0]], end: nodes[pair[1]] })
  }
  graph.AddNodes(nodes)
  graph.AddEdges(edges)
  return graph, nodes, edges
}

func buildRandPairs(numNodes int, numEdges int) [][]int {
  rand.Seed(time.Now().UnixNano())
  edgePairs := make([][]int, 0)
  for i := 0; i < numEdges; i++ {
    edgePairs = append(edgePairs, []int{rand.Intn(numNodes), rand.Intn(numNodes)})
  }
  return edgePairs
}
