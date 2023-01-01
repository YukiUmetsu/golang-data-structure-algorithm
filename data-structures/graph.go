package dsalgo

import "fmt"

type Graph struct {
	length    int
	adjacents map[string][]string
}

func (g *Graph) AddVertex(vertexName string) {
	if g.adjacents == nil {
		g.adjacents = make(map[string][]string)
	}
	g.adjacents[vertexName] = []string{}
	g.length++
}

func (g *Graph) AddEdge(vertexName1, vertexName2 string) {
	g.adjacents[vertexName1] = append(g.adjacents[vertexName1], vertexName2)
	g.adjacents[vertexName2] = append(g.adjacents[vertexName2], vertexName1)
}

func (g *Graph) ShowConnections() {
	for k, v := range g.adjacents {
		adjacentVertexString := ""
		if len(v) < 1 {
			continue
		}
		for _, vertexName := range v {
			adjacentVertexString += fmt.Sprintf(" %s", vertexName)
		}
		fmt.Println(k + "-->" + adjacentVertexString)
	}
}
