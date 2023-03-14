package main

import "fmt"

func init() {
	// fmt.Println("init")
}

func main() {
	defer func() {
		fmt.Println("recover: ", recover())
	}()
	panic("not success")

}
