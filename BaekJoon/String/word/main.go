package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var temp map[string]int = make(map[string]int)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var text string
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &text)

	splits := strings.Split(strings.ToUpper(text), "")
	for _, v := range splits {
		cnt := temp[v]
		temp[v] = cnt + 1
	}

	var result string
	var maxWord int = -1
	for key, value := range temp {
		if value > maxWord {
			result = key
			maxWord = value
		} else if value == maxWord {
			result = "?"
		}
	}

	fmt.Fprintln(w, result)
}
