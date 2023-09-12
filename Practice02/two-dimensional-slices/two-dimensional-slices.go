package main

import "fmt"

func main() {
	var TwoDArray [8][8]int
	TwoDArray[3][6] = 18
	TwoDArray[7][4] = 3
	fmt.Println(TwoDArray)

	// dynamic slice
	var rows int
	var cols int

	rows = 7
	cols = 9
	var twodslices = make([][]int, rows)

	for i := range twodslices {
		twodslices[i] = make([]int, cols)
	}

	fmt.Println(twodslices)

	arr := []int{5, 6, 7, 8, 9}
	slic1 := arr[:3]
	fmt.Println("slice1", slic1)
	slic2 := arr[1:5]
	fmt.Println("slice2", slic2)

	slic3 := append(slic2, 12)
	fmt.Println("slice3", slic3)
}
