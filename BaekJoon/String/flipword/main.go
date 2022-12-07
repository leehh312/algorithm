package main

import (
	"bufio"
	"os"
)

type Stack struct {
	data []rune
	size int
}

func New() *Stack {
	return &Stack{
		data: make([]rune, 0),
		size: 0,
	}
}

func (stack *Stack) Push(param rune) {
	stack.data = append(stack.data, param)
	stack.size++
}

func (stack *Stack) Pop() rune {
	if len(stack.data) == 0 {
		return ' '
	}

	stack.size--
	result := stack.data[stack.size]
	stack.data = stack.data[:stack.size]
	return result
}

func (stack *Stack) Size() int {
	return stack.size
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1000001), 1000001)
	sc.Scan()
	text := sc.Text()

	stack := New()
	var tagArea bool = false
	for _, word := range text {
		if word == '<' { // '<' 일 경우
			// 태그 내부 시작
			tagArea = true

			if stack.Size() > 0 {
				reverseWord(stack, w)
			}
			w.WriteRune(word)
		} else if word == '>' { // '>' 일 경우
			// 태그 내부 끝
			tagArea = false

			w.WriteRune(word)
		} else if word == ' ' { // 띄어쓰기 일 경우
			if stack.Size() > 0 {
				reverseWord(stack, w)
			}
			w.WriteRune(word)
		} else {
			// 태그 내부 단어일 경우
			if tagArea {
				w.WriteRune(word)
				continue
			}

			stack.Push(word)
		}
	}

	if stack.Size() > 0 {
		reverseWord(stack, w)
	}
}

func reverseWord(stack *Stack, result *bufio.Writer) {
	for stack.Size() > 0 {
		result.WriteRune(stack.Pop())
	}
}
