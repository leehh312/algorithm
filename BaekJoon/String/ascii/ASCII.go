package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
var char rune
fmt.Scanln(&char)
백준사이트에서 룬타입을 인식? 못하는거 같다.
*/
func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewScanner(os.Stdin)
	defer w.Flush()
	r.Scan()
	text := r.Text()

	fmt.Fprintf(w, "%d\n", []rune(text)[0])
}
