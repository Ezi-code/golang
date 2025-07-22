package main

import "fmt"

func main() {

	fmt.Println("Go slices...")

	mySlice := []int{1, 2, 3, 4}

	newSlice := make([]int, 10)
	y := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	mySlice = append(newSlice, y...)

	fmt.Println(mySlice)
	fmt.Println(newSlice)
}
