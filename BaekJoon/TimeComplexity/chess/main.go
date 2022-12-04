package main

import "fmt"

const (
	king  = 1
	queen = 1
	look  = 2
	bshop = 2
	night = 2
	pone  = 8
)

func main() {
	var a, b, c, d, e, f int
	fmt.Scanf("%d %d %d %d %d %d", &a, &b, &c, &d, &e, &f)
	fmt.Printf("%d %d %d %d %d %d", king-a, queen-b, look-c, bshop-d, night-e, pone-f)
}
