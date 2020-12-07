package tree

import (
	"errors"
	"github.com/lmx-Hexagram/data_structure/queue"
)

var (
	ErrorBinaryTreeIsEmpty  = errors.New("binary tree is empty")
	ErrorNotFound = errors.New("not found")
)

type ADT interface {
	IsEmpty() bool
	Depth() int
}

type BiNode struct {
	data       interface{}
	leftChild  *BiNode
	rightChild *BiNode
}

//type BiTree = BiNode

func Init() *BiNode {
	biNode := &BiNode{
		data:       nil,
		leftChild:  nil,
		rightChild: nil,
	}
	return biNode
}

func Create(stmt string) *BiNode {
	return nil
}

func (T *BiNode) IsEmpty() bool {
	if T.data == nil || T.leftChild == nil || T.rightChild == nil {
		return true
	}
	return false
}

func (T *BiNode) Depth() int {
	var l, r int = 0, 0
	if T.IsEmpty() {
		return 0
	}
	if T.leftChild != nil {
		l = T.leftChild.Depth()
	} else {
		l = 0
	}

	if T.rightChild != nil {
		r = T.rightChild.Depth()
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

func (T *BiNode)Point(data interface{}) (*BiNode,error) {
	if T.IsEmpty() {
		return nil,ErrorBinaryTreeIsEmpty
	}
	q := queue.Init(301)
	_ = q.Put(T)
	for !q.IsEmpty() {
		temp, _ := q.Poll()
		node := temp.(*BiNode)
		if node.data == data {
			return node,nil
		}
		if node.leftChild != nil {
			_ = q.Put(node.leftChild)
		}
		if node.rightChild != nil {
			_ = q.Put(node.rightChild)
		}
	}
	return nil,ErrorNotFound
}

