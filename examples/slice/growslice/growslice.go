package main

import "fmt"

// 测试slice扩容的规律
func main() {
	s := make([]int, 0)
	oldCap := cap(s)

	for i := 0; i < 2048; i++ {
		s = append(s, i)
		newCap := cap(s)

		if newCap != oldCap {
			fmt.Printf("[%d -> %4d] cap = %-4d  |  after append %-4d  cap = %-4d\n", 0, i, oldCap, i, newCap)
			oldCap = newCap
		}
	}
}
