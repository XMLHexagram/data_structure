package josephus

import (
	"fmt"
	"github.com/lmx-Hexagram/data_structure/link/circularLinkedList"
)

type Data struct {
	ID   int
	Code int
}

func Go() {
	// 约瑟夫环 Josephus problem
	var startAt int
	var code int

	var inputCode []int

	fmt.Print("start at:")
	_, _ = fmt.Scanf("%d", &startAt)
	fmt.Print("input:")
	for i := 1; ; i++ {
		_, _ = fmt.Scanf("%d", &code)
		if code == -1 {
			break
		}

		inputCode = append(inputCode, code)
	}

	ll := CreateJosephus(inputCode)
	list := DealJosephus(startAt, ll)
	for _, v := range list {
		fmt.Printf("kill: %d,Code: %d\n",v.Kill,v.Code)
	}
}

func CreateJosephus(testInput []int) *circularLinkedList.CNode {
	ll := circularLinkedList.Init()
	for i := 1; i <= len(testInput); i++ {
		//fmt.Println(testInput[i-1])
		ll.AddAtEnd(Data{
			ID:   i,
			Code: testInput[i-1],
		})
	}
	return ll
}

type List struct {
	Kill int
	Code int
}

func DealJosephus(startAt int, linkList *circularLinkedList.CNode) []List {
	list := make([]List, 0, 100)
	var step = startAt
	var length = linkList.Len()

	//linkList.Print()

	node, _ := linkList.Delete(step)

	list = append(list, List{
		Kill: node.Data.(Data).ID,
		Code: node.Data.(Data).Code,
	})
	//linkList.Print()
	//node.Next.Print()
	length--

	for ; length > 0; length-- {
		step = node.Data.(Data).Code

		node, _ = node.Next.Delete(step)
		//node.Next.Print()
		//fmt.Println(node.Next.Len())
		//fmt.Println("###################")
		list = append(list, List{
			Kill: node.Data.(Data).ID,
			Code: node.Data.(Data).Code,
		})
	}

	//node.Next.Print()
	//data := node.Next.Next.Data
	//fmt.Println("alive:", data.(Data).ID, ";code:", data.(Data).Code)
	return list
}
