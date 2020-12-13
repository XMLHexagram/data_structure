package graph

import (
	"testing"
)

func TestMGraph(t *testing.T) {
	g := New(6)
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")

	g.AddArcByVertexData("1","2",1)
	g.AddArcByVertexData("1","3",12)
	g.AddArcByVertexData("2","3",9)
	g.AddArcByVertexData("2","4",3)
	g.AddArcByVertexData("4","3",4)
	g.AddArcByVertexData("3","5",5)
	g.AddArcByVertexData("4","5",13)
	g.AddArcByVertexData("4","6",15)
	g.AddArcByVertexData("5","6",4)
	g.AddArcByVertexData("1","1",0)
	g.AddArcByVertexData("2","2",0)
	g.AddArcByVertexData("3","3",0)
	g.AddArcByVertexData("4","4",0)
	g.AddArcByVertexData("5","5",0)
	g.AddArcByVertexData("6","6",0)

	g.Print()
	g.Dijkstra()
}
