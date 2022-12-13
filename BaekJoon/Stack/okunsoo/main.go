package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	data []int
	size int
}

func New() *Stack {
	return &Stack{
		data: make([]int, 0),
		size: 0,
	}
}

func (stack *Stack) Push(val int) {
	stack.size++
	stack.data = append(stack.data, val)
}

func (stack *Stack) Pop() int {
	stack.size--
	result := stack.data[stack.size]
	stack.data = stack.data[:stack.size]

	return result
}

func (stack *Stack) Top() int {
	return stack.data[stack.size-1]
}

func (stack *Stack) isEmpty() bool {
	if stack.size > 0 {
		return false
	} else {
		return true
	}
}

// 오른쪽에 있는 수 중에서 자기보다 큰 가장 첫번째 숫자
/*
3, 5, 2, 7
9, 5, 4, 8
*/
func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanf(r, "%d\n", &n)

	input := make([]int, n)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d", &input[i])
	}

	stack := New()
	for i := 0; i < n; i++ {
		if stack.isEmpty() {
			stack.Push(i)
		}

		for !stack.isEmpty() && input[stack.Top()] < input[i] {
			result[stack.Pop()] = input[i]
		}

		stack.Push(i)
	}

	for !stack.isEmpty() {
		result[stack.Pop()] = -1
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d ", result[i])
	}
	fmt.Fprintln(w)
}
