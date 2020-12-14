package graph

import (
	"errors"
	"fmt"
	"github.com/lmx-Hexagram/data_structure/queue"
)

type ADT interface {
	LocateVertex(data interface{}) (index int, err error)
	FirstAdjVertex(index int) (resIndex int, err error)
}

const (
	Infinity = 99999
)

var (
	ErrorNotFound      = errors.New("not found")
	ErrInvalidData     = errors.New("invalid data")
	ErrNoMoreAdjVertex = errors.New("no more adj vertex")
)

//const (
//	DG  = iota // 有向图 DiGraph
//	UDG        // 无向图 Undigraph
//	DN         // 有向网
//	UDN        // 无向网
//)

type Arc struct {
	Data interface{}
	//Data   interface{}
}

type Vertex struct {
	Data interface{}
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

func CreateGraph(vertexes []interface{}, arcMatrix [][]interface{}) (g *MGraph, err error) {
	if len(vertexes) != len(arcMatrix) {
		return nil, ErrInvalidData
	}

	g = New(len(vertexes))
	for _, vertex := range vertexes {
		err = g.AddVertex(vertex)
		if err != nil {
			return nil, err
		}
	}
	for i, arcs := range arcMatrix {
		for j, arc := range arcs {
			err = g.AddArcByIndex(i, j, arc)
			if err != nil {
				return nil, err
			}
		}
	}
	return g, nil
}

func (g *MGraph) LocateVertex(data interface{}) (index int, err error) {
	for i := 0; i < g.VertexNumber; i++ {
		if g.Vertexes[i].Data == data {
			return i, nil
		}
	}
	return -1, ErrorNotFound
}

func (g *MGraph) FirstAdjVertex(index int) (resIndex int, err error) {
	resIndex, err = g.NextAdjVertex(index, 0)
	return
}

func (g *MGraph) NextAdjVertex(index int, start int) (resIndex int, err error) {
	for i := start + 1; i < g.VertexNumber; i++ {
		if g.ArcMatrix[index][i].Data != 0 {
			return i, nil
		}
	}
	return -1, ErrNoMoreAdjVertex
}

func (g *MGraph) AddVertex(data interface{}) (err error) {
	g.Vertexes[g.VertexNumber].Data = data
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

func (g *MGraph) AddArcByData(fromData interface{}, toData interface{}, data interface{}) (err error) {
	fromIndex, err := g.LocateVertex(fromData)
	if err != nil {
		return err
	}
	toIndex, err := g.LocateVertex(toData)
	if err != nil {
		return err
	}
	err = g.AddArcByIndex(fromIndex, toIndex, data)
	if err != nil {
		return err
	}
	return nil
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
func (g *MGraph) DFSTraverse(visit func(vertex Vertex) interface{}) (res []interface{}) {
	res = make([]interface{}, 0, 100)
	var visited []bool = make([]bool, g.VertexNumber)
	for i := 0; i < g.VertexNumber; i++ {
		if !visited[i] {
			g.DFS(i, visit, &visited, &res)
		}
	}
	return res
}

func (g *MGraph) DFS(index int, visit func(vertex Vertex) interface{}, visited *[]bool, res *[]interface{}) {
	(*visited)[index] = true
	*res = append(*res, visit(g.Vertexes[index]))

	for w, _ := g.FirstAdjVertex(index); w >= 0; w, _ = g.NextAdjVertex(index, w) {
		if !(*visited)[w] {
			g.DFS(w, visit, visited, res)
		}
	}
}

// 广度优先搜索
func (g *MGraph) BFSTraverse(visit func(vertex Vertex) interface{}) (res []interface{}) {
	res = make([]interface{}, 0, 100)
	var visited []bool = make([]bool, g.VertexNumber)
	q := queue.Init(301)
	for i := 0; i < g.VertexNumber; i++ {
		if !visited[i] {
			visited[i] = true
			res = append(res, visit(g.Vertexes[i]))
			q.Put(i)
			for !q.IsEmpty() {
				interfaceData, _ := q.Poll()
				index := interfaceData.(int)
				for w, _ := g.FirstAdjVertex(index); w >= 0; w, _ = g.NextAdjVertex(index, w) {
					if !visited[w] {
						visited[w] = true
						res = append(res, visit(g.Vertexes[w]))
						q.Put(w)
					}
				}
			}
		}
	}
	return res
}

// Dijkstra算法
func (g *MGraph) Dijkstra() []int {
	var dis []int = make([]int, g.VertexNumber)
	var res []bool = make([]bool, g.VertexNumber)
	var min = Infinity
	var index int

	for i := 0; i < g.VertexNumber; i++ {
		//hack
		if g.ArcMatrix[0][i].Data == nil || g.ArcMatrix[0][i].Data.(int) == 0 {
			dis[i] = Infinity
			continue
		}
		dis[i] = g.ArcMatrix[0][i].Data.(int)
	}
	//hack
	dis[0] = 0
	res[0] = true

	for i := 1; i < g.VertexNumber; i++ {
		min = Infinity
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
					fmt.Println(dis)
				}
			}
		}
	}

	return dis
}
