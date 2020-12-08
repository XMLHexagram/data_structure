package matrix

import (
	"fmt"
)

type ADT interface {
	Transpose() *Matrix
	Print()
}

type Matrix struct {
	Points []Point
	Row    int // 行
	Col    int // 列
	Count  int // 非零元个数
}

type Point struct {
	I    int // 行标
	J    int // 列标
	Data interface{}
}

func Init(row int, col int, points []Point) *Matrix {
	m := &Matrix{
		Points: points,
		Row:    row,
		Col:    col,
		Count:  len(points),
	}
	return m
}

func (m *Matrix) Transpose() *Matrix {
	r := &Matrix{
		Points: make([]Point, 0, 100),
		Row:    m.Col,
		Col:    m.Row,
		Count:  m.Count,
	}

	for col := 1; col <= m.Col; col++ {
		//fmt.Println(m.Count, "##")
		for count := 0; count < m.Count; count++ {
			if m.Points[count].J == col {
				r.Points = append(r.Points, Point{
					I:    m.Points[count].J,
					J:    m.Points[count].I,
					Data: m.Points[count].Data,
				})
			}
		}
	}

	return r
}

func (m *Matrix) Print() {
	fmt.Println(m.Points)

	count := 0
	for i := 1; i <= m.Row; i++ {
		for j := 1; j <= m.Col; j++ {
			if m.Points[count].I == i && m.Points[count].J == j {
				fmt.Print(m.Points[count].Data, " ")
				count++
			} else {
				fmt.Print("0", " ")
			}
		}
		fmt.Println()
	}
}
