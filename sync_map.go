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

func (ts *TestService) Incr(key, num int) {
	if testVal := ts.Get(key); testVal != nil {
		testVal.ID += num
	}
}

func main() {
	testService := newTestService()
	id := 1
	testService.Set(id)
	val := testService.Get(id)

	fmt.Println(val.ID)
	testService.Incr(id, 2)
	fmt.Println(val.ID)
}
