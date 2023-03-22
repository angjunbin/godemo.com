package main

import (
	"fmt"
)

func main() {
	s := "nihao中国521"

	tmp := []rune(s)

	length := len(tmp)

	for i := range tmp {
		if i == length/2 {
			break
		}

		tmp[i], tmp[length-i-1] = tmp[length-i-1], tmp[i]
	}

	fmt.Println(string(tmp))
}
