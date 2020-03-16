package main

import (
    "fmt"
    "time"
)

const LIM = 21

var fibs [LIM]uint64

func main() {
    var result uint64 = 0
    start := time.Now()
    for i := 0; i < LIM; i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
    end := time.Now()
    delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	
	result1 := 0
	start1 := time.Now()
	for i := 0; i <= LIM; i++ {
		result1 = fibonacci_old(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result1)
	}
	end1 := time.Now()
	delta1 := end1.Sub(start1)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta1)
}

func fibonacci(n int) (res uint64) {
    // memoization: check if fibonacci(n) is already known in array:
    if fibs[n] != 0 {
        res = fibs[n]
        return
    }
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    fibs[n] = res
    return
}

func fibonacci_old(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci_old(n-1) + fibonacci_old(n-2)
	}
	return
}