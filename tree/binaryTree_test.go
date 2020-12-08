package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := Init()
	tree.Data = 1
	tree.InsertChild(Left, &BiNode{
		Data:       2,
		LeftChild:  nil,
		RightChild: nil,
		Parent:     nil,
	})
	tree.InsertChild(Right,&BiNode{
		Data:       3,
		LeftChild:  nil,
		RightChild: nil,
		Parent:     nil,
	})

	fmt.Println(tree.LayerScan())
	tree.Print()
	fmt.Println(tree.Depth())
}
