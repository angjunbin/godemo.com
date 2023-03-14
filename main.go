package main

import "fmt"

func init() {
	// fmt.Println("init")
}

func main() {
	fmt.Println("hello java")

	defer func() {
		fmt.Println("recover: ", recover())
	}()
	panic("not success")

}
