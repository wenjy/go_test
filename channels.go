package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	a1 := a[:len(a)/2]
	fmt.Println(a1)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	//x, y := <-c, <-c // 从 c 中获取
	x := <-c
	y := <-c

	fmt.Println(x, y, x+y)

	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2
	fmt.Println(<-c1)
	fmt.Println(<-c1)

	c2 := make(chan int, 10)
	go fibonacci(cap(c2), c2)
	for i := range c2 {
		fmt.Println(i)
	}
}
