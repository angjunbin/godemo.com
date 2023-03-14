package main

import "fmt"

func init() {
	// fmt.Println("init")
}

type student struct {
	id   int32
	name string
}

func main() {
	defer func() {
		fmt.Println("recover: ", recover())
	}()
	panic("not success")

}
