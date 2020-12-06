package devilLanguage

import (
	"fmt"
	"github.com/lmx-Hexagram/data_structure/stack"
)

var translateMap = map[string]string{
	"t": "天",
	"d": "地",
	"s": "上",
	"a": "一只",
	"e": "鹅",
	"z": "追",
	"g": "赶",
	"x": "下",
	"n": "蛋",
	"h": "恨",
}

var simplifyMap = map[string]string{
	"B": "tAdA",
	"A": "sae",
}

func Go() {
	var char string
	var temp string
	var temp1 string

	s := stack.Init()

	for i := 1; ; i++ {
		_, _ = fmt.Scanf("%c", &char)
		if char == "#" {
			break
		}

		if char == "B" {
			temp = simplifyMap[char]
			for _, v := range temp {
				if v == "A" {
					
				}
			}
		}

		s.Push(char)
	}
}
