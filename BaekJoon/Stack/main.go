package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	data []int
	size int
}

func New(size int) *Stack {
	return &Stack{
		data: make([]int, size),
		size: 0,
	}
}

func (stack *Stack) Push(num int) error {
	if len(stack.data) > stack.size {
		stack.data[stack.size] = num
		stack.size++

		return nil
	}

	return fmt.Errorf("out of memory")
}

func (stack *Stack) Pop() int {
	if stack.Empty() == 1 {
		return -1
	} else {
		result := stack.data[stack.size-1]
		stack.size--

		return result
	}
}

func (stack *Stack) Empty() int {
	if stack.size == 0 {
		return 1
	} else {
		return 0
	}
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Top() int {
	if stack.Empty() == 1 {
		return -1
	} else {
		return stack.data[stack.size-1]
	}
}

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	scn := bufio.NewScanner(os.Stdin)

	var n int
	fmt.Scanln(&n)

	stack := New(n)

	for i := 0; i < n; i++ {
		scn.Scan()
		text := scn.Text()
		split := strings.Split(text, " ")

		switch split[0] {
		case "push":
			value, _ := strconv.Atoi(split[1])
			err := stack.Push(value)
			if err != nil {
				panic(err)
			}
		case "top":
			fmt.Fprintf(w, "%d\n", stack.Top())
		case "pop":
			fmt.Fprintf(w, "%d\n", stack.Pop())
		case "empty":
			fmt.Fprintf(w, "%d\n", stack.Empty())
		case "size":
			fmt.Fprintf(w, "%d\n", stack.Size())
		}
	}
}
