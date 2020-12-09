package matrix

import "testing"

func TestMatrix(T *testing.T) {
	//testInput := [][3]int{{1, 1, 3}, {1, 2, 5}, {2, 3, 7}, {2, 5, 11}, {3, 3, 2}, {3, 6, 8}, {5, 5, 9}, {5, 6, 10}}
	//points := make([]Point, 0, 100)
	//
	//for _, ints := range testInput {
	//	points = append(points, Point{
	//		I:    ints[0],
	//		J:    ints[1],
	//		Data: ints[2],
	//	})
	//}
	//
	//m := Init(5, 6, points)
	//m.Print()
	//m.NormalizePrint()
	//t := m.QuickTranspose()
	//t.NormalizePrint()

	testInput1 := [][3]int{{1, 1, 3}, {1, 4, 5}, {2, 2, -1}, {3, 1, 2}}
	points1 := make([]Point, 0, 100)

	for _, ints := range testInput1 {
		points1 = append(points1, Point{
			I:    ints[0],
			J:    ints[1],
			Data: ints[2],
		})
	}
	m := Init(3, 4, points1)
	m.Print()

	testInput2 := [][3]int{{1, 2, 2}, {2, 1, 1}, {3, 1, -2}, {3, 2, -4}}
	points2 := make([]Point, 0, 100)

	for _, ints := range testInput2 {
		points2 = append(points2, Point{
			I:    ints[0],
			J:    ints[1],
			Data: ints[2],
		})
	}
	n := Init(4, 2, points2)
	n.Print()

	r, err := m.Multiply(n)
	if err != nil {
		T.Error(err)
	}
	r.Print()
}
