package main

import (
	"fmt"
	"runtime"
	"time"
)

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	// fmt.Print(time.Now())
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// like if else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	fmt.Println(shouldEscape('?'))

	i := 0
	switch i {
		case 0: fallthrough
		case 1:
			fmt.Println(i) // 当 i == 0 时函数也会被调用
	}

	k := 6
	switch k {
		case 4: fmt.Println("was <= 4"); fallthrough;
		case 5: fmt.Println("was <= 5"); fallthrough;
		case 6: fmt.Println("was <= 6"); fallthrough;
		case 7: fmt.Println("was <= 7"); fallthrough;
		case 8: fmt.Println("was <= 8"); fallthrough;
		default: fmt.Println("default case")
	}

}
