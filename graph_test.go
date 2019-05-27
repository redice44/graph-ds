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
    graph, _, expectedEdges := setupGraph(numNodes, nodeNames, numEdges)
    edges := graph.GetEdges()
    nodes := graph.GetNodes()
    graphNumEdges := len(edges)
    graphNumNodes := len(nodes)
    if graphNumEdges != numEdges {
      t.Error(fmt.Sprintf("Expected %v edges not %v.", numEdges, graphNumEdges))
    }
    for _, edge := range edges {
      if !graph.HasEdge(edge) {
        t.Error(fmt.Sprintf("Edge expected but not found %v.", edge))
      }
    }
    for _, edge := range expectedEdges {
      if !graph.HasEdge(edge) {
        t.Error(fmt.Sprintf("Edge expected but not found %v.", edge))
      }
    }

    if graphNumNodes != numNodes {
      t.Error(fmt.Sprintf("Expected %v nodes not %v.", numNodes, graphNumNodes))
    }
    for _, node := range nodes {
      if !graph.HasNode(node) {
        t.Error(fmt.Sprintf("Node expected but not found %v.", node))
      }
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
