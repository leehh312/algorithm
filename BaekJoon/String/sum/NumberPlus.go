package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r := bufio.NewScanner(os.Stdin)

	var n int
	fmt.Scanln(&n)

	r.Scan()
	splits := strings.Split(r.Text(), "")
	var result int = 0
	for i := 0; i < len(splits); i++ {
		num, _ := strconv.Atoi(splits[i])
		result += num
	}

	fmt.Fprintln(w, result)
}
