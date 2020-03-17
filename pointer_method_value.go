package main

import (
    "fmt"
)

type B struct {
    thing int
}

func (b *B) change() { b.thing = 1 }

// 传值
func (b B) write() string { 
	b.thing = 2 
	return fmt.Sprint(b) 
}

// 指针方法和值方法都可以在指针或非指针上被调用
type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

func main() {
    var b1 B // b1是值
    b1.change()
    fmt.Println(b1.write())
    fmt.Println(b1.thing)

    b2 := new(B) // b2是指针
    b2.change()
    fmt.Println(b2.write())
    fmt.Println(b2.thing)


    // 值
    var lst List
    lst.Append(1)
    fmt.Printf("%v (len: %d)\n", lst, lst.Len()) // [1] (len: 1)

    // 指针
    plst := new(List)
    plst.Append(2)
    fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
}
