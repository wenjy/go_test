package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

type ErrNegativeSqrt float64

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}
	return fmt.Sprintf("Sqrt negative number: %v", float64(e))
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func Sqrt(x float64) (float64, error) {
	return x, ErrNegativeSqrt(x)
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
