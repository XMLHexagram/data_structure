package matrix

import "testing"

func TestMatrix(T *testing.T) {
	testInput := [][3]int{{1, 1, 3}, {1, 2, 5}, {2, 3, 7}, {2, 5, 11}, {3, 6, 8}, {5, 5, 9}, {5, 6, 10}}
	points := make([]Point, 0, 100)

	for _, ints := range testInput {
		points = append(points, Point{
			I:    ints[0],
			J:    ints[1],
			Data: ints[2],
		})
	}

	m := Init(5, 6, points)
	m.Print()
	t := m.Transpose()
	t.Print()
}
