package main

import (
	"fmt"
	"net"
	"sync"
	"unsafe"
)

func main() {
	var bool1 bool = true
	fmt.Println(unsafe.Sizeof(bool1)) // 1

	var byte1 byte = 1
	fmt.Println(unsafe.Sizeof(byte1)) // 1

	var uint81 uint8 = 1
	fmt.Println(unsafe.Sizeof(uint81)) // 1
	var int81 int8 = 1
	fmt.Println(unsafe.Sizeof(int81)) // 1

	var uint161 uint16 = 1
	fmt.Println(unsafe.Sizeof(uint161)) // 2
	var int161 int16 = 1
	fmt.Println(unsafe.Sizeof(int161)) // 2

	var uint321 uint32 = 1
	fmt.Println(unsafe.Sizeof(uint321)) // 4
	var int321 int32 = 1
	fmt.Println(unsafe.Sizeof(int321)) // 4
	var float321 float32 = 1
	fmt.Println(unsafe.Sizeof(float321)) // 4

	var uint641 uint64 = 1
	fmt.Println(unsafe.Sizeof(uint641)) // 8
	var int641 int64 = 1
	fmt.Println(unsafe.Sizeof(int641)) // 8
	var float641 float64 = 1
	fmt.Println(unsafe.Sizeof(float641)) // 8
	var complex641 complex64 = 1
	fmt.Println(unsafe.Sizeof(complex641)) // 8

	var complex1281 complex128 = 1
	fmt.Println(unsafe.Sizeof(complex1281)) // 16

	var int1 int = 1
	fmt.Println(unsafe.Sizeof(int1)) // 8 64位系统，32位为4 不推荐
	var uint1 uint = 1
	fmt.Println(unsafe.Sizeof(uint1)) // 8 64位系统

	var uintptr1 uintptr = 1
	fmt.Println(unsafe.Sizeof(uintptr1)) // 8

	var str1 string = "a"
	fmt.Println(unsafe.Sizeof(str1)) // 16

	uint1ptr := &uint1
	fmt.Println(unsafe.Sizeof(uint1ptr)) // 8

	var sp1 = []byte{1, 2, 3}
	fmt.Println(unsafe.Sizeof(sp1)) // 24

	var map1 = map[int]int{1: 1}
	fmt.Println(unsafe.Sizeof(map1)) // 8

	chan1 := make(chan int, 10)
	fmt.Println(unsafe.Sizeof(chan1)) // 8

	f1 := func() {}
	fmt.Println(unsafe.Sizeof(f1)) // 8

	fmt.Println(unsafe.Sizeof(Sa{}))     // 0
	fmt.Println(unsafe.Sizeof([0]int{})) // 0

	// 具体结构体
	var t111 t11 = t11{}
	fmt.Println(unsafe.Sizeof(t111)) // 136
	var t121 t12 = t12{}
	fmt.Println(unsafe.Sizeof(t121)) // 128

	var t211 t21 = t21{}
	fmt.Println(unsafe.Sizeof(t211)) // 80
	var t222 t22 = t22{}
	fmt.Println(unsafe.Sizeof(t222)) // 72

	var t311 t31 = t31{}
	fmt.Println(unsafe.Sizeof(t311)) // 112

	var t411 t41 = t41{}
	fmt.Println(unsafe.Sizeof(t411)) // 48
	var t421 t42 = t42{}
	fmt.Println(unsafe.Sizeof(t421)) // 40

	var t511 t51 = t51{}
	fmt.Println(unsafe.Sizeof(t511)) // 168

	var t611 t61 = t61{}
	fmt.Println(unsafe.Sizeof(t611)) // 56
	var t621 t62 = t62{}
	fmt.Println(unsafe.Sizeof(t621)) // 48
}

type Sa struct{}

type t11 struct {
	a string       // 16
	b uint8        // 1
	c string       // 16
	d net.Conn     // 接口 16
	e *net.UDPConn // 8
	f net.Addr     // 2个函数 16
	g []byte       // 24
	h []byte       // 24
	i bool         // 1
}

// 优化后
type t12 struct {
	a string       // 16
	c string       // 16
	d net.Conn     // 接口 16
	e *net.UDPConn // 8
	f net.Addr     // 2个函数 16
	g []byte       // 24
	h []byte       // 24
	i bool         // 1
	b uint8        // 1
}

type t21 struct {
	a []byte // 24
	b uint8  // 1
	c string // 16
	d uint16 // 2
	e []byte // 24
}

// 优化后
type t22 struct {
	a []byte // 24
	c string // 16
	e []byte // 24
	b uint8  // 1
	d uint16 // 2
}

// 无可优化
type t31 struct {
	a net.Conn         // 接口 16
	b int64            // 8
	c string           // 16
	d int64            //8
	e int64            // 8
	f uint32           // 4
	g []byte           // 24
	h int64            // 8
	i chan interface{} // 8
	j chan interface{} //8
}

type t41 struct {
	a net.Conn // 接口 16
	b uint8    // 1
	c uint8
	d uint8
	e string // 16
	f uint16 // 2
	g uint8
	h uint8
}

// 优化后
type t42 struct {
	a net.Conn // 接口 16
	e string   // 16
	b uint8    // 1
	c uint8
	d uint8
	f uint16 // 2
	g uint8
	h uint8
}

// 无可优化
type t51 struct {
	a sync.RWMutex // 24
	b map[int]int  //8
	c map[int]int  // 8
	d chan bool    // 8
	e sync.Once    // 12
	f int          // 8
	g *int         // 8
	h *int         // 8
	i string       // 16
	j int          // 8
	k string       // 16
	l int64        // 8
	m uint32       // 4
	n uint32       //4
	o int64        //8
	p int64        // 8
	q uint32       // 4
	r uint32       // 4
}

type t61 struct {
	a uint16 // 2
	b uint8  // 1
	c uint8  // 1
	d string // 16
	e uint16 // 2
	f []byte // 24
}

// 优化后
type t62 struct {
	a uint16 // 2
	b uint8  // 1
	c uint8  // 1
	e uint16 // 2
	d string // 16
	f []byte // 24
}
