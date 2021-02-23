package main

import (
	"fmt"
	"net"
	"reflect"
)

// run linux
func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}

	fdVal := reflect.Indirect(reflect.ValueOf(listener)).FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")
	fmt.Println(int(pfdVal.FieldByName("Sysfd").Int()))
}
