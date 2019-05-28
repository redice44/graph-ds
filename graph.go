package graph

import (
	"errors"
	"fmt"
)

type Graph struct {
	name  string
	nodes []Node
	edges [][]int
}

func (g Graph) String() string {
	s := fmt.Sprintf("Graph: %v\n\n", g.name)
	s += "Nodes: \n"
	for i := range g.nodes {
		s += fmt.Sprintf("%v\n", g.nodes[i].name)
	}

	s += "\nEdges: \n"

	for start := range g.edges {
		for end, edge := range g.edges[start] {
			if edge > 0 {
				s += fmt.Sprintf("[%v]: %v -> %v\n", edge, g.nodes[start].name, g.nodes[end].name)
			}
		}
	}

	return s
}

type Node struct {
	name string
}

func (n Node) String() string {
	return n.name
}

type Edge struct {
	start Node
	end   Node
}

func (e Edge) String() string {
	return fmt.Sprintf("%v -> %v\n", e.start, e.end)
}

func contains(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func (g *Graph) BuildTree(start Node) Graph {
	edges, nodes := g.BuildCompositeTree(start, []int{})
	graph := Graph{name: start.name}
	nodes = append([]Node{start}, nodes...)
	graph.AddNodes(nodes)
	graph.AddEdges(edges)
	return graph
}

func (g *Graph) BuildCompositeTree(start Node, exclusionList []int) ([]Edge, []Node) {
	neighborEdges := g.GetNeighborsEdges(start)
	startIndex, _ := g.FindNodeIndex(start)
	excludedNeighbors := make([]int, 1)
	excludedNeighbors[0] = startIndex
	edges := make([]Edge, 0)
	nodes := make([]Node, 0)

	for _, edge := range neighborEdges {
		if index, err := g.FindNodeIndex(edge.end); err == nil {
			excludedNeighbors = append(excludedNeighbors, index)
		}
	}

	for _, neighborEdge := range neighborEdges {
		if index, err := g.FindNodeIndex(neighborEdge.end); err == nil {
			if !contains(exclusionList, index) {
				nodes = append(nodes, g.nodes[index])
				treeEdges, treeNodes := g.BuildCompositeTree(
					neighborEdge.end,
					append(exclusionList, excludedNeighbors...),
				)
				edges = append(
					edges,
					neighborEdge,
				)
				edges = append(
					edges,
					treeEdges...,
				)
				nodes = append(
					nodes,
					treeNodes...,
				)
			}
		}
	}
	return edges, nodes
}

func (g *Graph) GetNeighborsEdges(start Node) []Edge {
	neighbors := make([]Edge, 0)
	if index, err := g.FindNodeIndex(start); err == nil {
		for i, v := range g.edges[index] {
			if v > 0 {
				neighbors = append(neighbors, Edge{start: start, end: g.nodes[i]})
			}
		}
	}
	return neighbors
}

func (g *Graph) FindNodeIndex(target Node) (int, error) {
	for i, n := range g.nodes {
		if n == target {
			return i, nil
		}
	}
	return 0, errors.New("Node not found.")
}

func (g *Graph) AddNode(n Node) {
	g.nodes = append(g.nodes, n)
	g.edges = append(g.edges, make([]int, len(g.edges)))
	for i := range g.edges {
		g.edges[i] = append(g.edges[i], 0)
	}
}

func (g *Graph) AddNodes(nodes []Node) {
	for _, n := range nodes {
		g.AddNode(n)
	}
}

func (g *Graph) HasNode(target Node) bool {
	for _, node := range g.nodes {
		if node == target {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(e Edge) {
	startIndex, _ := g.FindNodeIndex(e.start)
	endIndex, _ := g.FindNodeIndex(e.end)
	g.edges[startIndex][endIndex] += 1
}

func (g *Graph) AddEdges(edges []Edge) {
	for _, e := range edges {
		g.AddEdge(e)
	}
}

func (g *Graph) HasEdge(edge Edge) bool {
	startIndex, _ := g.FindNodeIndex(edge.start)
	endIndex, _ := g.FindNodeIndex(edge.end)
	return g.edges[startIndex][endIndex] > 0
}

func (g *Graph) GetEdges() []Edge {
	edges := make([]Edge, 0)
	for startIndex, ends := range g.edges {
		for endIndex, v := range ends {
			for i := 0; i < v; i++ {
				edges = append(edges, Edge{start: g.nodes[startIndex], end: g.nodes[endIndex]})
			}
		}
	}
	return edges
}

func (g *Graph) GetNodes() []Node {
	return g.nodes
}

func New(name string) Graph {
	return Graph{name: name, nodes: make([]Node, 0)}
}
