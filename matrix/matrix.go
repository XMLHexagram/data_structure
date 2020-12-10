package matrix

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMatrixOnMultiply = errors.New("invalid matrix on multiply")
	ErrInvalidMatrixOnPlus     = errors.New("invalid matrix on plus")
)

type ADT interface {
	Transpose() *Matrix
	QuickTranspose() *Matrix
	Normalize() [][]interface{}
	NormalizePrint()
	Print()
	Multiply(N *Matrix) (res *Matrix, err error)
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
	Data int
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

func (m *Matrix) QuickTranspose() *Matrix {
	r := &Matrix{
		Points: make([]Point, m.Count),
		Row:    m.Col,
		Col:    m.Row,
		Count:  m.Count,
	}

	// 每一列有几个
	nums := make([]int, m.Col)
	// 每一列的第一个 在转换之后应该在的位置
	// 每一列第一个非零元所在的位置
	colPos := make([]int, m.Col)
	for _, point := range m.Points {
		nums[point.J-1]++
	}
	colPos[0] = 1
	for i := 1; i < len(colPos); i++ {
		colPos[i] = colPos[i-1] + nums[i-1]
	}

	for _, point := range m.Points {
		r.Points[colPos[point.J-1]-1] = Point{
			I:    point.J,
			J:    point.I,
			Data: point.Data,
		}
		colPos[point.J-1]++
	}

	return r
}

func (m *Matrix) Print() {
	//fmt.Println(m.Points)

	count := 0
	for i := 1; i <= m.Row; i++ {
		for j := 1; j <= m.Col; j++ {
			if count < len(m.Points) && m.Points[count].I == i && m.Points[count].J == j {
				fmt.Print(m.Points[count].Data, " ")
				count++
			} else {
				fmt.Print("0", " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *Matrix) Normalize() [][]int {
	res := make([][]int, m.Row)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, m.Col)
	}

	for _, v := range m.Points {
		res[v.I-1][v.J-1] = v.Data
	}

	return res
}

func (m *Matrix) NormalizePrint() {
	matrix := m.Normalize()
	for _, raw := range matrix {
		for _, point := range raw {
			if point != nil {
				fmt.Print(point, " ")
				continue
			}
			fmt.Print("0", " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func (M *Matrix) Multiply(N *Matrix) (res *Matrix, err error) {
	if M.Col != N.Row {
		return nil, ErrInvalidMatrixOnMultiply
	}

	res = &Matrix{
		Points: make([]Point, 0, 100),
		Row:    M.Row,
		Col:    N.Col,
		Count:  0,
	}

	for i := 0; i < M.Row; i++ {
		// count 可能有坑
		for j := 0; j < N.Col; j++ {
			var temp int
			for m := 0; m < M.Count; m++ {
				for n := 0; n < N.Count; n++ {
					if M.Points[m].I-1 == i && N.Points[n].J-1 == j && M.Points[m].J == N.Points[n].I {
						//fmt.Println(M.Points[m], " ", N.Points[n])
						temp = temp + M.Points[m].Data*N.Points[n].Data
						//fmt.Println(temp)
					}
					//fmt.Println(1)
				}
			}
			if temp != 0 {
				res.Points = append(res.Points, Point{
					I:    i + 1,
					J:    j + 1,
					Data: temp,
				})
			}
		}
	}
	//fmt.Println(res.Points)
	return res, nil
}

func (M *Matrix) Plus(N *Matrix) (res *Matrix, err error) {
	if M.Col != N.Col || M.Row != N.Row {
		return nil, ErrInvalidMatrixOnPlus
	}

	res = &Matrix{
		Points: make([]Point, 0, 100),
		Row:    M.Row,
		Col:    M.Col,
		Count:  0,
	}

	//for i := 0; i < M.Row; i++ {
	//	for j := 0; j < M.Col; j++ {
	//
	//	}
	//}
	for _, mPoint := range M.Points {
		for _, nPoint := range N.Points {
			if mPoint.I == nPoint.I && mPoint.J == nPoint.J {
				mPoint.Data = mPoint.Data + nPoint.Data
				break
			}
		}
		res.Points = append(res.Points, mPoint)
	}
	return res, nil
}

func (M *Matrix) Minus(N *Matrix) (res *Matrix, err error) {
	if M.Col != N.Col || M.Row != N.Row {
		return nil, ErrInvalidMatrixOnPlus
	}

	res = &Matrix{
		Points: make([]Point, 0, 100),
		Row:    M.Row,
		Col:    M.Col,
		Count:  0,
	}

	for _, mPoint := range M.Points {
		for _, nPoint := range N.Points {
			if mPoint.I == nPoint.I && mPoint.J == nPoint.J {
				mPoint.Data = mPoint.Data - nPoint.Data
				break
			}
		}
		res.Points = append(res.Points, mPoint)
	}
	return res, nil
}
