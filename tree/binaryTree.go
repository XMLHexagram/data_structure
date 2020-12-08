package tree

import (
	"errors"
	"fmt"
	"github.com/lmx-Hexagram/data_structure/queue"
)

const (
	Left = iota
	Right
)

var (
	ErrorBinaryTreeIsEmpty = errors.New("binary tree is empty")
	ErrorNotFound          = errors.New("not found")
)

type ADT interface {
	IsEmpty() bool
	Depth() int
	Point(data interface{}) (*BiNode, error)
	InsertChild(LR int, node *BiNode)
}

type BiNode struct {
	Data       interface{}
	LeftChild  *BiNode
	RightChild *BiNode
	Parent     *BiNode
}

//type BiTree = BiNode

func Init() *BiNode {
	biNode := &BiNode{
		Data:       nil,
		LeftChild:  nil,
		RightChild: nil,
	}
	return biNode
}

func Create(stmt string) *BiNode {
	return nil
}

func (T *BiNode) IsEmpty() bool {
	if T.Data == nil || T.LeftChild == nil || T.RightChild == nil {
		return true
	}
	return false
}

func (T *BiNode) Depth() int {
	var l, r int = 0, 0
	if T.IsEmpty() {
		return 0
	}
	if T.LeftChild != nil {
		l = T.LeftChild.Depth()
	} else {
		l = 0
	}

	if T.RightChild != nil {
		r = T.RightChild.Depth()
	} else {
		r = 0
	}

	if r > l {
		return r + 1
	} else if r < l {
		return l + 1
	} else {
		return r
	}
}

func (T *BiNode) Point(data interface{}) (*BiNode, error) {
	if T.IsEmpty() {
		return nil, ErrorBinaryTreeIsEmpty
	}
	q := queue.Init(301)
	_ = q.Put(T)
	for !q.IsEmpty() {
		temp, _ := q.Poll()
		node := temp.(*BiNode)
		if node.Data == data {
			return node, nil
		}
		if node.LeftChild != nil {
			_ = q.Put(node.LeftChild)
		}
		if node.RightChild != nil {
			_ = q.Put(node.RightChild)
		}
	}
	return nil, ErrorNotFound
}

func (T *BiNode) InsertChild(LR int, node *BiNode) {
	if LR == Left {
		var temp *BiNode
		var t bool = false
		if T.LeftChild != nil {
			temp = T.LeftChild
			t = true
		}
		T.LeftChild = node
		if t {
			node.LeftChild = temp
		}
	} else {
		var temp *BiNode
		var t bool = false
		if T.RightChild != nil {
			temp = T.RightChild
			t = true
		}
		T.RightChild = node
		if t {
			node.RightChild = temp
		}
	}
}

func (T *BiNode) LayerScan() [][]*BiNode {
	res := make([][]*BiNode, 0, 100)
	list := make([]*BiNode, 0, 100)
	q := queue.Init(301)

	_ = q.Put(T)
	for !q.IsEmpty() {
		count := q.Len()
		for i := 0; i < count; i++ {
			nodeInterface, _ := q.Poll()
			node := nodeInterface.(*BiNode)
			list = append(list, node)
			if node.LeftChild != nil {
				_ = q.Put(node.LeftChild)
			}
			if node.RightChild != nil {
				_ = q.Put(node.RightChild)
			}
		}
		res = append(res, list)
		list = make([]*BiNode, 0, 100)
	}

	return res
}

func (T *BiNode) Print() {
	for _, nodes := range T.LayerScan() {
		for _, node := range nodes {
			fmt.Print(node.Data," ")
		}
		fmt.Println()
	}
}
