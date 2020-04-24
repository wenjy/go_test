package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("In main()")
    go longWait()
    go shortWait()
    fmt.Println("About to sleep in main()")
	// 当 main() 函数返回的时候，程序退出：它不会等待任何其他非 main 协程的结束，这里sleep 3s 等待测试协程执行
    time.Sleep(3 * 1e9)
    fmt.Println("At the end of main()")
}

func longWait() {
    fmt.Println("Beginning longWait()")
    time.Sleep(2 * 1e9) // sleep for 5 seconds
    fmt.Println("End of longWait()")
}

func shortWait() {
    fmt.Println("Beginning shortWait()")
    time.Sleep(1 * 1e9) // sleep for 2 seconds
    fmt.Println("End of shortWait()")
}