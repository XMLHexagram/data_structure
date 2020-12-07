package devilLanguage

import (
	"fmt"
	"github.com/lmx-Hexagram/data_structure/stack"
)

//var translateMap = map[string]string{
//	"t": "天",
//	"d": "地",
//	"s": "上",
//	"a": "一只",
//	"e": "鹅",
//	"z": "追",
//	"g": "赶",
//	"x": "下",
//	"n": "蛋",
//	"h": "恨",
//}

var simplifyMap = map[rune]string{
	'B': "tAdA",
	'A': "sae",
}

func Go() {
	var devilSentence string

	resultS := stack.Init()
	initS := stack.Init()
	tempS := stack.Init()

	var expandString string
	for _, v := range simplifyMap['B'] {
		if v == 'A' {
			expandString += simplifyMap[v]
			continue
		}
		expandString += string(v)
	}
	simplifyMap['B'] = expandString

	fmt.Println("start")
	_, _ = fmt.Scanln(&devilSentence)

	for _, v := range devilSentence {
		initS.Push(v)
	}

	for !initS.IsEmpty() {
		char := initS.Pop().(rune)

		if char > rune('A') && char < rune('Z') {
			temp := simplifyMap[char]
			for i := len(temp); i > 0; i-- {
				resultS.Push(rune(temp[i-1]))
			}
		}

		if char == ')' {
			char := initS.Pop().(rune)
			for char != '(' {
				tempS.Push(char)
				char = initS.Pop().(rune)
			}

			repeatChar := tempS.Pop().(rune)
			resultS.Push(repeatChar)
			for !tempS.IsEmpty() {
				resultS.Push(tempS.Pop().(rune))
				resultS.Push(repeatChar)
			}
		}
		//s.Push(char)
	}

	for !resultS.IsEmpty() {
		fmt.Print(string(resultS.Pop().(rune)))
	}
}
