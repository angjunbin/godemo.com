package main

import "fmt"

func init() {
	// fmt.Println("init")
}

func main() {
<<<<<<< Updated upstream
	fmt.Println("hello java")
=======
	fmt.Println("hello go")
>>>>>>> Stashed changes

	defer func() {
		fmt.Println("recover: ", recover())
	}()
	panic("not success")

}
