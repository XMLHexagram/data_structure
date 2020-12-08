package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	test := []int{1,2,3,4,5,6,7}
	test1 := make([]interface{},0,100)

	for _, v := range test {
		test1 = append(test1,v)
	}
	tree := Create(test1)
	tree.Print()
}
