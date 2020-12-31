package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32 = 0
	var b int32 = 0
	// 原子增减
	newA := atomic.AddInt32(&a, int32(1))
	fmt.Println("AddInt32 new", newA)

	newA = atomic.AddInt32(&a, int32(-1))
	fmt.Println("AddInt32 new", newA)

	// 比较并交换----Compare And Swap 简称CAS
	// 假设被操作的值未曾被改变（即与旧值相等），并一旦确定这个假设的真实性就立即进行值替换
	swapped := atomic.CompareAndSwapInt32(&b, int32(0), int32(1))
	fmt.Println("CompareAndSwapInt32", swapped)

	// 原子读取
	b1 := atomic.LoadInt32(&b)
	fmt.Println("LoadInt32", b1)

	// 原子保存
	atomic.StoreInt32(&b, int32(2))
	b1 = atomic.LoadInt32(&b)
	fmt.Println("LoadInt32", b1)

	// 原子交换
	oldB := atomic.SwapInt32(&b, int32(3))
	newB := atomic.LoadInt32(&b)
	fmt.Println("LoadInt32 newB", newB)
	fmt.Println("SwapInt32 oldB", oldB)

	// 原子操作由底层硬件支持，而锁则由操作系统的调度器实现。
	// 锁应当用来保护一段逻辑，对于一个变量更新的保护，原子操作通常会更有效率，并且更能利用计算机多核的优势，
	// 如果要更新的是一个复合对象，则应当使用atomic.Value封装好的实现。
	var readOnly atomic.Value

	ac := &atomicValue{1}
	readOnly.Store(ac)
	readOnly.Store(ac)
	ac.id = 3
	ac1 := readOnly.Load().(*atomicValue)
	fmt.Println("atomicValue", ac1)

}

type atomicValue struct {
	id int
}
