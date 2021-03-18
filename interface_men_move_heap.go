package main

type S123 struct {
	s1 int
}

func (s *S123) M1(i int) { s.s1 = i }

type I123 interface {
	M1(int)
}

func g() {
	var s1 S123 // 逃逸
	var s2 S123 // 不逃逸
	var s3 S123 // 不逃逸

	f123(&s1)
	f223(&s2)
	f323(&s3)
}

func f123(s I123)  { s.M1(42) }
func f223(s *S123) { s.M1(42) }
func f323(s I123)  { s.(*S123).M1(42) }

// go build -gcflags='-m=2' ./interface_men_move_heap.go
