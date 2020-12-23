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
}
