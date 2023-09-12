package main

import "fmt"

func h(x, y int) int {
	return x * y
}

func g(l, m int) (x int, y int) {
	x = 2 * 1
	y = 4 * m
	return
}

func main() {
	fmt.Println(h(g(1, 2)))
}
