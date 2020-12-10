package matrix

import "testing"

func TestMatrix_Multiply(t *testing.T) {
	testData1 := [][3]int{{1, 1, 3}, {1, 4, 5}, {2, 2, -1}, {3, 1, 2}}
	m := testDateToMatrix(testData1, 3, 4)

	testData2 := [][3]int{{1, 2, 2}, {2, 1, 1}, {3, 1, -2}, {3, 2, -4}}
	n := testDateToMatrix(testData2, 4, 2)

	r, err := m.Multiply(n)
	if err != nil {
		t.Error(err)
	}

	answer := [][]int{{0, 6}, {-1, 0}, {0, 4}}
	checkAnswer(r.Normalize(), answer, t)
	//todo:再测试一组数据
}

func TestMatrix_Transpose(t *testing.T) {
	testData1 := [][3]int{{1, 2, 3}, {1, 4, 5}, {2, 2, -1}, {3, 1, 2}}
	m := testDateToMatrix(testData1, 3, 4)

	r := m.Transpose()
	answer := [][]int{{0, 0, 2}, {3, -1, 0}, {0, 0, 0}, {5, 0, 0}}
	checkAnswer(r.Normalize(), answer, t)
}

func TestMatrix_QuickTranspose(t *testing.T) {
	testData1 := [][3]int{{1, 2, 3}, {1, 4, 5}, {2, 2, -1}, {3, 1, 2}}
	m := testDateToMatrix(testData1, 3, 4)

	r := m.QuickTranspose()
	answer := [][]int{{0, 0, 2}, {3, -1, 0}, {0, 0, 0}, {5, 0, 0}}
	checkAnswer(r.Normalize(), answer, t)
}

func TestMatrix_Normalize(t *testing.T) {
	testData1 := [][3]int{{1, 2, 3}, {1, 4, 5}, {2, 2, -1}, {3, 1, 2}}
	m := testDateToMatrix(testData1, 3, 4)

	r := m.Normalize()
	answer := [][]int{{0, 3, 0, 5}, {0, -1, 0, 0}, {2, 0, 0, 0}}
	checkAnswer(r, answer, t)
}

func TestMatrix_Plus(t *testing.T) {
	testData1 := [][3]int{{1, 1, 10}, {2, 3, 9}, {3, 1, -1}}
	testData2 := [][3]int{{2, 3, -1}, {3, 1, 1}, {3, 3, -3}}
	m := testDateToMatrix(testData1, 3, 3)
	n := testDateToMatrix(testData2, 3, 3)
	r, err := m.Plus(n)
	if err != nil {
		t.Error(err)
	}
	m.Print()
	n.Print()
r.Print()
	answer := [][]int{{10, 0, 0}, {0, 0, 8}, {0, 0, -3}}
	checkAnswer(r.Normalize(), answer, t)
}

func TestMatrix_Minus(t *testing.T) {
	testData1 := [][3]int{{1, 1, 10}, {2, 2, 9}, {3, 1, -1}}
	testData2 := [][3]int{{2, 3, -1}, {3, 1, 1}, {3, 2, -3}}
	m := testDateToMatrix(testData1, 3, 2)
	n := testDateToMatrix(testData2, 3, 2)
	r, err := m.Minus(n)
	if err != nil {
		t.Error(err)
	}

	answer := [][]int{{10, 0}, {0, 10}, {-2, -3}}
	checkAnswer(r.Normalize(), answer, t)
}

func testDateToMatrix(testData [][3]int, row int, col int) *Matrix {
	points := make([]Point, 0, 100)

	for _, ints := range testData {
		points = append(points, Point{
			I:    ints[0],
			J:    ints[1],
			Data: ints[2],
		})
	}
	m := Init(row, col, points)
	return m
}

func checkAnswer(res [][]int, answer [][]int, t *testing.T) {
	for i, raw := range res {
		for j, point := range raw {
			if answer[i][j] != point {
				t.Errorf("wrong result point:expect %d but get %d", answer[i][j], point)
			}
		}
	}
}
