package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	stack []rune = make([]rune, 0)
	size  int    = 0
)

func Push(param rune) {
	size++
	stack = append(stack, param)
}

func Pop() rune {
	size--
	result := stack[size]
	stack = stack[:size]

	return result
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 100001), 100001)
	sc.Scan()
	text := sc.Text()

	var stick int = 0
	var prevWord rune
	for _, word := range text {
		if word == '(' {
			Push(word)
		} else {
			Pop()

			if prevWord == '(' { // '(' 다음 ')' 들어오면 레이저
				// 스택 안의 괄호 갯수만큼 막대 카운트
				stick += size
			} else { // ')' 다음 ')'가 들어온 경우 막대기 끝의 개수 하나 카운트
				stick++
			}
		}
		prevWord = word
	}

	fmt.Fprintln(w, stick)
}
