package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	remove(slice, 3)
	fmt.Println(slice)
}

func remove(slices []int, i int) []int {
	copy(slices[i:], slices[i+1:])
	fmt.Println("copy after:", slices)
	return slices[:len(slices)-1]
}
