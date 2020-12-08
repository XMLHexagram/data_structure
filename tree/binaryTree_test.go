package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	testArrayInt := []int{1, 2, 3, 4, 5, 6, 7}
	testArrayInterface := make([]interface{}, 0, 100)
	for _, v := range testArrayInt {
		testArrayInterface = append(testArrayInterface, v)
	}

	tree := new(BiNode)
	t.Run("test create", func(t *testing.T) {
		var err error
		tree, err = Create(testArrayInterface)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}
	})

	t.Run("test point", func(t *testing.T) {
		node, err := tree.Point(3)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}
		if node.LeftChild.Data != 6 || node.RightChild.Data != 7 {
			t.Errorf("get wrong node")
		}

		node, err = tree.Point(1)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}
		if node.LeftChild.Data != 2 || node.RightChild.Data != 3 {
			t.Errorf("get wrong node")
		}
	})

	t.Run("test Insert", func(t *testing.T) {
		node, err := tree.Point(7)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}

		if node.RightChild != nil {
			t.Errorf("tree error")
		}
		node.InsertData(Left, 8)
		node.InsertData(Right, 9)
		_, err = tree.Point(8)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}
		_, err = tree.Point(9)
		if err != nil {
			t.Errorf("%s:%s", t.Name(), err)
		}

		var expect [][]int = [][]int{{1}, {2, 3}, {4, 5, 6, 7}, {-1, -1, -1, -1, -1, -1, 8, 9}}
		for i, datas := range ToDataArray(tree.LayerScan()) {
			for j, data := range datas {
				if data == nil {
					if expect[i][j] != -1 {
						t.Errorf("nil doesn't match -1")
					}
				} else {
					if data.(int) != expect[i][j] {
						t.Errorf("%d doesn't match %d", data.(int), expect[i][j])
					}
				}
			}
		}
	})

	t.Run("test depth", func(t *testing.T) {
		if tree.Depth() != 4 {
			t.Errorf("depth wrong:expect %d get %d", 4, tree.Depth())
		}
	})

	t.Run("test delete", func(t *testing.T) {
		node, _ := tree.Point(7)
		node.DeleteChild(Left)
		node.DeleteChild(Right)
		var expect [][]int = [][]int{{1}, {2, 3}, {4, 5, 6, 7}}
		for i, datas := range ToDataArray(tree.LayerScan()) {
			for j, data := range datas {
				if data == nil {
					if expect[i][j] != -1 {
						t.Errorf("nil doesn't match -1")
					}
				} else {
					if data.(int) != expect[i][j] {
						t.Errorf("%d doesn't match %d", data.(int), expect[i][j])
					}
				}
			}
		}
	})

	t.Run("test depth", func(t *testing.T) {
		if tree.Depth() != 3 {
			t.Errorf("depth wrong:expect %d get %d", 3, tree.Depth())
		}
	})
}

func ExampleBiNode_Print() {
	testArrayInt := []int{1, 2, 3, 4, 5, 6, 7}
	testArrayInterface := make([]interface{}, 0, 100)
	for _, v := range testArrayInt {
		testArrayInterface = append(testArrayInterface, v)
	}

	tree, _ := Create(testArrayInterface)
	tree.Print()
	// Output
	// 1
	// 2 3
	// 4 5 6 7
}
