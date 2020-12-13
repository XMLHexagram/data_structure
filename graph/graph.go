package graph

import (
	"errors"
	"fmt"
	"github.com/lmx-Hexagram/data_structure/queue"
)

type ADT interface {
	LocateVertexByData(data interface{}) (index int, err error)
	AddVertex(data interface{}) (err error)
	AddArcByIndex(fromIndex int, toIndex int) (err error)
	AddArcByVertexData(fromData interface{}, toData interface{}) (err error)
	FirstAdjVertex(index int) (int, error)
	NextAdjVertex(index int, start int) int
	Print()
}

var (
	ErrorNotFound     = errors.New("not found")
	ErrorAlreadyExist = errors.New("already exist")
)

const (
	DG  = iota // 有向图 DiGraph
	UDG        // 无向图 Undigraph
	DN         // 有向网
	UDN        // 无向网
)

type Arc struct {
	Data interface{}
	//Data   interface{}
}

type Vertex struct {
	data interface{}
}

// Matrix Graph
type MGraph struct {
	Vertexes     []Vertex
	ArcMatrix    [][]Arc
	ArcNumber    int
	VertexNumber int
	GraphKind    int
}

func New(maxVertexNum int) *MGraph {
	ArcMatrix := make([][]Arc, maxVertexNum)
	for i := 0; i < maxVertexNum; i++ {
		ArcMatrix[i] = make([]Arc, maxVertexNum)
	}
	return &MGraph{
		Vertexes:     make([]Vertex, maxVertexNum),
		ArcMatrix:    ArcMatrix,
		ArcNumber:    0,
		VertexNumber: 0,
		GraphKind:    0,
	}
}

func (g *MGraph) LocateVertexByData(data interface{}) (index int, err error) {
	for i := 0; i < g.VertexNumber; i++ {
		if g.Vertexes[i].data == data {
			return i, nil
		}
	}
	return -1, ErrorNotFound
}

func (g *MGraph) AddVertex(data interface{}) (err error) {
	g.Vertexes[g.VertexNumber].data = data
	g.VertexNumber++
	return nil
}

func (g *MGraph) AddArcByIndex(fromIndex int, toIndex int, data interface{}) (err error) {
	g.ArcMatrix[fromIndex][toIndex] = Arc{Data: data}
	g.ArcNumber++
	return nil
}

//func (g *MGraph) AddArcBothByVertexData(fromData interface{}, toData interface{}) (err error) {
//	err = g.AddArcByVertexData(fromData, toData)
//	if err != nil {
//		return err
//	}
//	err = g.AddArcByVertexData(toData, fromData)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func (g *MGraph) AddArcByVertexData(fromData interface{}, toData interface{}, data interface{}) (err error) {
	fromIndex, err := g.LocateVertexByData(fromData)
	if err != nil {
		return err
	}
	toIndex, err := g.LocateVertexByData(toData)
	if err != nil {
		return err
	}
	err = g.AddArcByIndex(fromIndex, toIndex, data)
	if err != nil {
		return err
	}
	return nil
}

func (g *MGraph) FirstAdjVertex(index int) (resIndex int, hasAdj bool) {
	for i := 0; i < g.VertexNumber; i++ {
		if g.ArcMatrix[index][i].Data != 0 {
			return i, true
		}
	}
	return -1, false
}

func (g *MGraph) NextAdjVertex(index int, start int) (resIndex int, hasNext bool) {
	for i := start + 1; i < g.VertexNumber; i++ {
		if g.ArcMatrix[index][i].Data != 0 {
			return i, true
		}
	}
	return -1, false
}

func (g *MGraph) Print() {
	fmt.Println(g.Vertexes)
	for _, matrix := range g.ArcMatrix {
		for _, arc := range matrix {
			fmt.Print(arc, " ")
		}
		fmt.Println()
	}
}

// 深度优先搜索
func (g *MGraph) DFSTraverse(visit func(vertex Vertex)) {
	var visited []bool = make([]bool, g.VertexNumber)
	for i := 0; i < g.VertexNumber; i++ {
		if !visited[i] {

			g.DFS(i, visit, &visited)
		}
	}

}

func (g *MGraph) DFS(index int, visit func(vertex Vertex), visited *[]bool) {
	(*visited)[index] = true
	visit(g.Vertexes[index])
	for w, _ := g.FirstAdjVertex(index); w >= 0; w, _ = g.NextAdjVertex(index, w) {
		if !(*visited)[w] {
			g.DFS(w, visit, visited)
		}
	}
}

func (g *MGraph) BFSTraverse(visit func(vertex Vertex)) {
	var visited []bool = make([]bool, g.VertexNumber)
	q := queue.Init(301)
	for i := 0; i < g.VertexNumber; i++ {
		if !visited[i] {
			visited[i] = true
			visit(g.Vertexes[i])
			q.Put(i)
			for !q.IsEmpty() {
				interfaceData, _ := q.Poll()
				index := interfaceData.(int)
				for w, _ := g.FirstAdjVertex(index); w >= 0; w, _ = g.NextAdjVertex(index, w) {
					if !visited[w] {
						visited[w] = true
						visit(g.Vertexes[w])
						q.Put(w)
					}
				}
			}
		}
	}
}

// Dijkstra算法
func (g *MGraph) Dijkstra() {
	var dis []int = make([]int, g.VertexNumber)
	var res []bool = make([]bool, g.VertexNumber)
	var min = 999
	var index int

	for i := 0; i < g.VertexNumber; i++ {
		if g.ArcMatrix[0][i].Data == nil {
			dis[i] = 999
			continue
		}
		dis[i] = g.ArcMatrix[0][i].Data.(int)
	}
	res[0] = true

	for i := 1; i < g.VertexNumber; i++ {
		min = 999
		for j := 0; j < g.VertexNumber; j++ {
			if res[j] == false && dis[j] < min {
				min = dis[j]
				index = j
			}
		}
		res[index] = true
		for k := 0; k < g.VertexNumber; k++ {
			if g.ArcMatrix[index][k].Data != 0 && g.ArcMatrix[index][k].Data != nil {
				if dis[k] > dis[index]+g.ArcMatrix[index][k].Data.(int) {
					dis[k] = dis[index] + g.ArcMatrix[index][k].Data.(int)
				}
			}
		}
	}

	for i, re := range dis {
		fmt.Println(i, ":", re)
	}
}
