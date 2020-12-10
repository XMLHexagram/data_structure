package matrix

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidMatrixOnMultiply = errors.New("invalid matrix on multiply")
	ErrInvalidMatrixOnPlus     = errors.New("invalid matrix on plus")
	ErrInvalidMatrixOnMinus    = errors.New("invalid matrix on minus")
)

type ADT interface {
	Transpose() *Matrix
	QuickTranspose() *Matrix
	Normalize() [][]int
	NormalizePrint() //
	Print()          //
	Multiply(N *Matrix) (res *Matrix, err error)
	Plus(N *Matrix) (res *Matrix, err error)
	Minus(N *Matrix) (res *Matrix, err error)
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

func (M *Matrix) Transpose() *Matrix {
	r := &Matrix{
		Points: make([]Point, 0, 100),
		Row:    M.Col,
		Col:    M.Row,
		Count:  M.Count,
	}

	for col := 1; col <= M.Col; col++ {
		//fmt.Println(m.Count, "##")
		for count := 0; count < M.Count; count++ {
			if M.Points[count].J == col {
				r.Points = append(r.Points, Point{
					I:    M.Points[count].J,
					J:    M.Points[count].I,
					Data: M.Points[count].Data,
				})
			}
		}
	}

	return r
}

func (M *Matrix) QuickTranspose() *Matrix {
	r := &Matrix{
		Points: make([]Point, M.Count),
		Row:    M.Col,
		Col:    M.Row,
		Count:  M.Count,
	}

	// 每一列有几个
	nums := make([]int, M.Col)
	// 每一列的第一个 在转换之后应该在的位置
	// 每一列第一个非零元所在的位置
	colPos := make([]int, M.Col)
	for _, point := range M.Points {
		nums[point.J-1]++
	}
	colPos[0] = 1
	for i := 1; i < len(colPos); i++ {
		colPos[i] = colPos[i-1] + nums[i-1]
	}

	for _, point := range M.Points {
		r.Points[colPos[point.J-1]-1] = Point{
			I:    point.J,
			J:    point.I,
			Data: point.Data,
		}
		colPos[point.J-1]++
	}

	return r
}

func (M *Matrix) Print() {
	//fmt.Println(m.Points)

	count := 0
	for i := 1; i <= M.Row; i++ {
		for j := 1; j <= M.Col; j++ {
			if count < len(M.Points) && M.Points[count].I == i && M.Points[count].J == j {
				fmt.Print(M.Points[count].Data, " ")
				count++
			} else {
				fmt.Print("0", " ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (M *Matrix) Normalize() [][]int {
	res := make([][]int, M.Row)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, M.Col)
	}

	for _, v := range M.Points {
		res[v.I-1][v.J-1] = v.Data
	}

	return res
}

func (M *Matrix) NormalizePrint() {
	matrix := M.Normalize()
	for _, raw := range matrix {
		for _, point := range raw {
			if point != 0 {
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

func (M Matrix) Plus(N *Matrix) (res *Matrix, err error) {
	if M.Col != N.Col || M.Row != N.Row {
		return nil, ErrInvalidMatrixOnPlus
	}

	//var mNum []int = make([]int, M.Row)
	//for _, v := range M.Points {
	//	mNum[v.I-1]++
	//}
	//
	//var mRowPos []int = make([]int, M.Row)
	//mRowPos[0] = 1
	//for i := 1; i < len(mRowPos); i++ {
	//	mRowPos[i] = mRowPos[i-1] + mNum[i-1]
	//}
	//
	//var nNum []int = make([]int, N.Row)
	//for _, v := range M.Points {
	//	nNum[v.I-1]++
	//}
	//
	//var nRowPos []int = make([]int, N.Row)
	//nRowPos[0] = 1
	//for i := 1; i < len(nRowPos); i++ {
	//	nRowPos[i] = nRowPos[i-1] + nNum[i-1]
	//}

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

	//for i := 0; i < M.Row; i++ {
	//	for j := 0; j < M.Col; j++ {
	//		if  {
	//
	//		}
	//	}
	//}

	mCount := 0
	nCount := 0
	for {
		if M.Points[mCount].I == N.Points[nCount].I {

			if M.Points[mCount].J == N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    M.Points[mCount].I,
					J:    M.Points[mCount].J,
					Data: M.Points[mCount].Data + N.Points[nCount].Data,
				})
				mCount++
				nCount++
			} else if M.Points[mCount].J < N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    M.Points[mCount].I,
					J:    M.Points[mCount].J,
					Data: M.Points[mCount].Data,
				})
				mCount++
			} else if M.Points[mCount].J > N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    N.Points[mCount].I,
					J:    N.Points[mCount].J,
					Data: N.Points[mCount].Data,
				})
				nCount++
			}
		} else if M.Points[mCount].I < N.Points[nCount].I {
			res.Points = append(res.Points, Point{
				I:    M.Points[mCount].I,
				J:    M.Points[mCount].J,
				Data: M.Points[mCount].Data,
			})
			mCount++
		} else if M.Points[mCount].I > N.Points[nCount].I {
			res.Points = append(res.Points, Point{
				I:    N.Points[mCount].I,
				J:    N.Points[mCount].J,
				Data: N.Points[mCount].Data,
			})
			nCount++
		}
		if mCount == len(M.Points) {
			res.Points = append(res.Points, N.Points[nCount:]...)
			break
		}
		if nCount == len(N.Points) {
			res.Points = append(res.Points, M.Points[nCount:]...)
			break
		}
	}

	//for _, mPoint := range M.Points {
	//	for _, nPoint := range N.Points {
	//		if mPoint.I == nPoint.I && mPoint.J == nPoint.J {
	//			mPoint.Data = mPoint.Data + nPoint.Data
	//			break
	//		}
	//	}
	//	if mPoint.Data != 0 {
	//		fmt.Println(mPoint.Data)
	//		res.Points = append(res.Points, mPoint)
	//	}
	//}
	return res, nil
}

func (M *Matrix) Minus(N *Matrix) (res *Matrix, err error) {
	if M.Col != N.Col || M.Row != N.Row {
		return nil, ErrInvalidMatrixOnMinus
	}

	res = &Matrix{
		Points: make([]Point, 0, 100),
		Row:    M.Row,
		Col:    M.Col,
		Count:  0,
	}

	mCount := 0
	nCount := 0
	for {
		if M.Points[mCount].I == N.Points[nCount].I {

			if M.Points[mCount].J == N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    M.Points[mCount].I,
					J:    M.Points[mCount].J,
					Data: M.Points[mCount].Data - N.Points[nCount].Data,
				})
				mCount++
				nCount++
			} else if M.Points[mCount].J < N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    M.Points[mCount].I,
					J:    M.Points[mCount].J,
					Data: M.Points[mCount].Data,
				})
				mCount++
			} else if M.Points[mCount].J > N.Points[nCount].J {
				res.Points = append(res.Points, Point{
					I:    N.Points[mCount].I,
					J:    N.Points[mCount].J,
					Data: -N.Points[mCount].Data,
				})
				nCount++
			}
		} else if M.Points[mCount].I < N.Points[nCount].I {
			res.Points = append(res.Points, Point{
				I:    M.Points[mCount].I,
				J:    M.Points[mCount].J,
				Data: M.Points[mCount].Data,
			})
			mCount++
		} else if M.Points[mCount].I > N.Points[nCount].I {
			res.Points = append(res.Points, Point{
				I:    N.Points[mCount].I,
				J:    N.Points[mCount].J,
				Data: -N.Points[mCount].Data,
			})
			nCount++
		}
		if mCount == len(M.Points) {
			//res.Points = append(res.Points, N.Points[nCount:]...)

			for _, point := range N.Points[nCount:] {
				res.Points = append(res.Points, Point{
					I:    point.I,
					J:    point.J,
					Data: -point.Data,
				})
			}
			break
		}
		if nCount == len(N.Points) {
			res.Points = append(res.Points, M.Points[nCount:]...)
			break
		}
	}
	return res, nil
}

//func (M *Matrix)GetRowPos()  {
//
//}
