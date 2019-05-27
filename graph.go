package graph

import (
  "errors"
  "fmt"
)

type Graph struct {
  name string
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
    for end, edge:= range g.edges[start] {
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
  end Node
}

func (e Edge) String() string {
  return fmt.Sprintf("%v -> %v\n", e.start, e.end)
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
        edges = append(edges, Edge{ start: g.nodes[startIndex], end: g.nodes[endIndex] })
      }
    }
  }
  return edges
}
