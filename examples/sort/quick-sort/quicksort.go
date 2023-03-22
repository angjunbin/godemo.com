package main

import "fmt"

func main() {
	data := []int{1, 0, 2, 8, 2, 0, 1, 8}
	QuickSort(data)
	fmt.Println(data)
}

// 快排
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	//基准数
	base := data[0]
	//左右两个指针
	l, r := 0, len(data)-1
	for i := 1; i < r; {
		if data[i] > base {
			data[i], data[r] = data[r], data[i]
			r--
		} else {
			data[i], data[l] = data[l], data[i]
			i++
			l++
		}
	}
	// fmt.Println("循环一次：", data)
	QuickSort(data[:l])
	QuickSort(data[l+1:])
}
