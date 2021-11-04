package main

import (
	"fmt"
	"sync"
)

type TestVal struct {
	ID int
}
type test struct {
	sync.Map
}

var testMap = &test{}

type TestService struct {
	TestMap *test
}

func newTestService() *TestService {
	return &TestService{TestMap: testMap}
}

func (ts *TestService) Get(key int) *TestVal {
	if testVal, ok := ts.TestMap.Load(key); ok {
		return testVal.(*TestVal)
	}
	return nil
}

func (ts *TestService) Set(id int) {
	ts.TestMap.Store(id, &TestVal{id})
}

func (ts *TestService) Pull(key int) *TestVal {
	if v, ok := ts.TestMap.LoadAndDelete(key); ok {
		return v.(*TestVal)
	}
	return nil
}

func (ts *TestService) Put(id int) *TestVal {
	if v, ok := ts.TestMap.LoadOrStore(id, &TestVal{id}); ok {
		return v.(*TestVal)
	}
	return nil
}

func (ts *TestService) Del(id int) {
	ts.TestMap.Delete(id)
}

func main() {
	testService := newTestService()
	testService.Set(1)

	val1 := testService.Get(1)
	fmt.Println(val1)

	val2 := testService.Pull(1)
	fmt.Println(val2)

	val3 := testService.Get(1)
	fmt.Println(val3)

	val4 := testService.Put(4) // 原来没有
	fmt.Println(val4)
	val5 := testService.Get(4)
	fmt.Println(val5)

	testService.Del(4)
	val6 := testService.Get(4)
	fmt.Println(val6)

	testService.Set(5)
	testService.Set(6)
	testService.Set(7)
	testService.Set(8)
	testMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		if k == 5 {
			testMap.Delete(5)
		}
		if k == 7 {
			return false
		}
		return true
	})
}
