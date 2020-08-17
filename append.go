package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)

	// 将切片 b1 的元素追加到切片 a1 之后：a1 = append(a1, b1...)
	a1 := []int{1}
	b1 := []int{2, 3}
	a1 = append(a1, b1...) // 第二个参数扩展成一个列表
	printSlice("a1", a1)

	// 复制a1到b2
	b2 := make([]int, len(a1))
	copy(b2, a1)
	printSlice("b2", b2)

	// 删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
	a2 := []int{1, 2, 3, 4}
	a2 = append(a2[:1], a2[1+1:]...)
	printSlice("a2", a2)

	// 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j:]...)，不包括j
	a3 := []int{1, 2, 3, 4}
	a3 = append(a3[:2], a3[3:]...)
	printSlice("a3", a3)

	// 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
	a4 := []int{1, 2, 3, 4}
	a4 = append(a4, make([]int, 1)...)
	printSlice("a4", a4)

	// 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
	a5 := []int{1, 2, 3, 4}
	a5 = append(a5[:2], append([]int{5}, a5[2:]...)...)
	printSlice("a5", a5)

	// 在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
	a6 := []int{1, 2, 3, 4}
	a6 = append(a6[:2], append(make([]int, 2), a6[2:]...)...)
	printSlice("a6", a6)

	// 在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
	a7 := []int{1, 2, 3, 4}
	b7 := []int{7, 8}
	a7 = append(a7[:2], append(b7, a7[2:]...)...)
	printSlice("a7", a7)

	// 取出位于切片 a 最末尾的元素 x：x, a = a[len(a)-1], a[:len(a)-1]
	a8 := []int{1, 2, 3, 4}
	var x = []int{}
	x, a8 = a8[len(a8)-1:], a8[:len(a8)-1]
	printSlice("a8", a8)
	printSlice("x", x)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
