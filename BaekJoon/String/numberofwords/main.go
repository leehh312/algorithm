package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r := bufio.NewReaderSize(os.Stdin, 1000001)
	line, _, _ := r.ReadLine()
	text := strings.TrimSpace(string(line))

	if text == "" {
		fmt.Fprintf(w, "%d", 0)
	} else {
		splits := strings.Split(text, " ")
		fmt.Fprintf(w, "%d", len(splits))
	}
}
