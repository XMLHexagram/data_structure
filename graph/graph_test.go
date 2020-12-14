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
	//fmt.Println(g.Dijkstra())
}

func TestMGraph2(t *testing.T) {
	g := New(25)

	citys := []string{
		"北京", "天津", "沈阳", "长春", "哈尔滨",
		"大连", "徐州", "上海", "南昌", "福州",
		"株洲", "广州", "深圳", "武汉", "郑州",
		"西安", "成都", "昆明", "贵阳", "柳州",
		"南宁", "呼和浩特", "兰州", "西宁", "乌鲁木齐"}
	for _, city := range citys {
		g.AddVertex(city)
	}

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			g.ArcMatrix[i][j].Data = 0
		}
	}

	links := [][2]int{
		{0, 1}, {0, 14}, {0, 21},
		{1, 2}, {1, 6},
		{2, 3}, {2, 5},
		{3, 4},
		{6, 7}, {6, 14},
		{7, 8},
		{8, 9}, {8, 10},
		{10, 11}, {10, 13}, {10, 18}, {10, 19},
		{11, 12},
		{13, 14},
		{14, 15},
		{15, 16}, {15, 22},
		{16, 17}, {16, 18},
		{17, 18},
		{18, 19},
		{19, 20},
		{21, 22},
		{22, 23}, {22, 24}}


	for _, link := range links {
		g.AddArcByIndex(link[0], link[1], 1)
		g.AddArcByIndex(link[1], link[0], 1)
	}


	res := g.DFSTraverse(visit)
	fmt.Println(res)
	a := g.BFSTraverse(visit)
	fmt.Println(a)
	//g.Print()
}
