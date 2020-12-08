package tree

import (
	"errors"
	"fmt"
	"github.com/lmx-Hexagram/data_structure/queue"
	"math"
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
	InsertData(LR int, Data interface{})
	DeleteChild(LR int, node *BiNode)
	LayerScan() [][]*BiNode
	Print()
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

func Create(datas []interface{}) *BiNode {
	T := Init()
	q := queue.Init(301)
	_ = q.Put(T)

	depth := int(math.Floor(math.Sqrt(float64(len(datas))))) + 1

	k := 0
	for i := 0; i < depth; i++ {
		length := q.Len()
		for j := 0; j < length; j++ {
			nodeInterface, _ := q.Poll()
			node := nodeInterface.(*BiNode)
			node.Data = datas[k]
			k++
			if i+1 != depth {
				node.LeftChild = new(BiNode)
				_ = q.Put(node.LeftChild)

				node.RightChild = new(BiNode)
				_ = q.Put(node.RightChild)
			}
		}
	}
	return T
}

func (T *BiNode) IsEmpty() bool {
	if T.Data == nil && T.LeftChild == nil && T.RightChild == nil {
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
		return r + 1
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

func (T *BiNode) InsertData(LR int, data interface{}) {
	T.InsertChild(LR, &BiNode{
		Data:       data,
		LeftChild:  nil,
		RightChild: nil,
		Parent:     nil,
	})
}

func (T *BiNode) DeleteChild(LR int) {
	if LR == Left {
		T.LeftChild = nil
	} else {
		T.RightChild = nil
	}
}

func (T *BiNode) LayerScan() [][]*BiNode {
	res := make([][]*BiNode, 0, 100)
	list := make([]*BiNode, 0, 100)
	q := queue.Init(301)

	depth := T.Depth()

	_ = q.Put(T)
	for i := 0; i < depth; i++ {
		count := q.Len()
		for i := 0; i < count; i++ {
			nodeInterface, _ := q.Poll()
			//fmt.Println(nodeInterface)
			//if reflect.ValueOf(nodeInterface)== nil {
			//	fmt.Println("###")
			//
			//}
			//fmt.Println(nodeInterface)
			node := nodeInterface.(*BiNode)
			if node == nil {
				//fmt.Println("success")
				list = append(list, nil)
				_ = q.Put(node)
				_ = q.Put(node)
				continue
			}
			list = append(list, node)
			//if node.LeftChild != nil {
			_ = q.Put(node.LeftChild)
			//}
			//if node.RightChild != nil {
			_ = q.Put(node.RightChild)
			//}
		}
		res = append(res, list)
		list = make([]*BiNode, 0, 100)
	}

	return res
}

func (T *BiNode) Print() {
	for _, nodes := range T.LayerScan() {
		for _, node := range nodes {
			if node != nil {
				fmt.Print(node.Data, " ")
			} else {
				fmt.Print("nil", " ")
			}

		}
		fmt.Println()
	}
}
