package main

import "fmt"

func main() {
	slice := []int{1, 3, 5, 6}
	twiceValue(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Println("new slice value", slice[i])
	}

	//slice = append(slice, 8)
	//fmt.Println("Capacity", cap(slice))
	//fmt.Println("Length", len(slice))
}

func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}
