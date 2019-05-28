package graph

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	baseGraph := buildRandomGraph(5, "Node", 10)
	tree := setupTree()
	shortTree := setupShortTree()
	treeNodes := tree.GetNodes()
	treeEdges := tree.GetEdges()
	shortTreeNodes := shortTree.GetNodes()
	shortTreeEdges := shortTree.GetEdges()
	t.Run("Base Graph", testGraph(baseGraph, baseGraph.GetNodes(), baseGraph.GetEdges()))
	t.Run("Base Graph", testGraph(tree, treeNodes, treeEdges))
	t.Run("Tree", testGraph(tree.BuildTree(treeNodes[0]), treeNodes, []Edge{
		treeEdges[0],
		treeEdges[1],
		treeEdges[4],
	}))
	t.Run(
		"Short Tree",
		testGraph(
			shortTree.BuildTree(shortTreeNodes[0]),
			[]Node{
				shortTreeNodes[0],
				shortTreeNodes[1],
				shortTreeNodes[2],
			},
			[]Edge{
				shortTreeEdges[0],
				shortTreeEdges[1],
			},
		),
	)
}

func testGraph(graph Graph, expectedNodes []Node, expectedEdges []Edge) func(*testing.T) {
	return func(t *testing.T) {
		edges := graph.GetEdges()
		nodes := graph.GetNodes()
		graphNumEdges := len(edges)
		graphNumNodes := len(nodes)
		expectedNumEdges := len(expectedEdges)
		expectedNumNodes := len(expectedNodes)

		if graphNumEdges != expectedNumEdges {
			t.Error(fmt.Sprintf("Expected %v edges not %v.", expectedNumEdges, graphNumEdges))
		}
		for _, edge := range edges {
			if !graph.HasEdge(edge) {
				t.Error(fmt.Sprintf("Edge expected but not found %v.", edge))
			}
		}

		if graphNumNodes != expectedNumNodes {
			t.Error(fmt.Sprintf("Expected %v nodes not %v.", expectedNumNodes, graphNumNodes))
		}
		for _, node := range nodes {
			if !graph.HasNode(node) {
				t.Error(fmt.Sprintf("Node expected but not found %v.", node))
			}
		}
	}
}

func setupTree() Graph {
	nodes := makeNodes("Node", 4)
	graph := Graph{name: "test", nodes: make([]Node, 0)}
	pairs := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
		[]int{2, 0},
		[]int{2, 3},
	}
	edges := buildEdges(nodes, pairs)
	graph.AddNodes(nodes)
	graph.AddEdges(edges)
	return graph
}

func setupShortTree() Graph {
	nodes := makeNodes("Node", 4)
	graph := Graph{name: "test", nodes: make([]Node, 0)}
	pairs := [][]int{
		[]int{0, 1},
		[]int{0, 2},
		[]int{1, 2},
		[]int{2, 0},
	}
	edges := buildEdges(nodes, pairs)
	graph.AddNodes(nodes)
	graph.AddEdges(edges)
	return graph
}

func buildRandomGraph(numNodes int, baseNodeName string, numEdges int) Graph {
	nodes := makeNodes("Node", numNodes)
	graph := Graph{name: "test", nodes: make([]Node, 0)}
	pairs := buildRandPairs(numNodes, numEdges)
	edges := buildEdges(nodes, pairs)
	graph.AddNodes(nodes)
	graph.AddEdges(edges)
	return graph
}

func makeNodes(baseNodeName string, numNodes int) []Node {
	nodes := make([]Node, 0)
	for i := 0; i < numNodes; i++ {
		nodes = append(nodes, Node{name: fmt.Sprintf("%v %v", baseNodeName, i)})
	}
	return nodes
}

func buildRandPairs(numNodes int, numEdges int) [][]int {
	rand.Seed(time.Now().UnixNano())
	edgePairs := make([][]int, 0)
	for i := 0; i < numEdges; i++ {
		edgePairs = append(edgePairs, []int{rand.Intn(numNodes), rand.Intn(numNodes)})
	}
	return edgePairs
}

func buildEdges(nodes []Node, pairs [][]int) []Edge {
	edges := make([]Edge, 0)
	for _, pair := range pairs {
		edges = append(edges, Edge{start: nodes[pair[0]], end: nodes[pair[1]]})
	}
	return edges
}
