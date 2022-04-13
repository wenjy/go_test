package main

import "fmt"

func main() {
	s1 := make([]byte, 10)
	fmt.Println("s1 cap", cap(s1)) // 10
	fmt.Println("s1 len", len(s1)) // 10

	s2 := s1[3:6]
	fmt.Println("s2 cap", cap(s2)) // 7
	fmt.Println("s2 len", len(s2)) // 3

	s3 := s1[0:3]
	fmt.Println("s3 cap", cap(s3)) // 10
	fmt.Println("s3 len", len(s3)) // 3

	s4 := s1[10:]
	fmt.Println("s4 cap", cap(s4)) // 0
	fmt.Println("s4 len", len(s4)) // 0

	s6 := s1[:0]
	fmt.Println("s6 cap", cap(s6)) // 10
	fmt.Println("s6 len", len(s6)) // 0

	copy(s2, []byte{97, 98, 99})
	fmt.Println("s1[0]", s1[0])
	fmt.Println("s2[0]", s2[0])

	var s5 []byte
	fmt.Println("s5 len", len(s5))
}
