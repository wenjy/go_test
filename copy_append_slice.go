package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	// func copy(dst, src []T) int
	// copy 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，
	// 并且返回拷贝的元素个数
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
