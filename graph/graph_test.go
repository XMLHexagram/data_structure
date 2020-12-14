package graph

import (
	"fmt"
	"testing"
)

func visit(vertex Vertex) interface{} {
	return vertex.Data
}

func TestMGraph(t *testing.T) {
	var vertexes1 []interface{} = []interface{}{"v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8"}
	var arcMatrix1 [][]interface{} = [][]interface{}{
		{0, 1, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0, 0}}

	g, err := CreateGraph(vertexes1, arcMatrix1)
	if err != nil {
		t.Error(err)
	}
	g.Print()
	res := g.BFSTraverse(visit)
	fmt.Println(res)
	res = g.DFSTraverse(visit)

	fmt.Println(res)
	fmt.Println(g.Dijkstra())
}
