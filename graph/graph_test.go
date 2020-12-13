package graph

import (
	"fmt"
	"testing"
)

func TestMGraph(t *testing.T) {
	g := New(8)
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")
	g.AddVertex("7")
	g.AddVertex("8")

	g.AddArcBothByVertexData("1", "2")
	g.AddArcBothByVertexData("1", "3")
	g.AddArcBothByVertexData("2", "4")
	g.AddArcBothByVertexData("2", "5")
	g.AddArcBothByVertexData("4", "8")
	g.AddArcBothByVertexData("5", "8")
	g.AddArcBothByVertexData("3", "7")
	g.AddArcBothByVertexData("3", "6")
	g.AddArcBothByVertexData("6", "7")

	//index, _ := g.LocateVertexByData("a")
	//fmt.Println(g.FirstAdjVertex(index))
	//fmt.Println(g.NextAdjVertex(index, 3))
	visit := func(vertex Vertex) {
		fmt.Print(vertex.data, " ")
	}
	g.DFSTraverse(visit)
	fmt.Println()
	g.BFSTraverse(visit)
	fmt.Println()
	g.Print()
}
