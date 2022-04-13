package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

var JsonIter jsoniter.API

func init() {
	JsonIter = jsoniter.ConfigCompatibleWithStandardLibrary
}

type Request1 struct {
	Method string `json:"method"` // 方法标识
	Seq    string `json:"seq"`    // 消息标识
}

type AuthRx struct {
	Request1
	Data *AuthDataRx
}

type AuthDataRx struct {
	Sn    string `json:"sn"`    //SN码
	Mac   string `json:"mac"`   //设备mac地址
	Token string `json:"token"` // 验证token
}

func main() {
	str1 := `{"method":"authDev","data":{"sn":"sn","mac":"mac","token":"token"},"seq":"1357924680"}`
	str2 := `{"method":
	"authDev","authDev":{"mac":"000e1faba8bd","sn":"130101000994888",
	"product":"T18.Pro","model":"AC1200",   "softVer":"V2.0.0#30_20220408181902.","hardVer":"V1","wwlinkerVer":"V1.0.0","devToken":"000e1faba8bd130101000994888"},"seq":"0880997380"}`

	str3 := "{\"method\":\"authDev\",\"authDev\":{\"mac\":\"000e1faba8bd\",\"sn\":\"130101000994888\",\"product\":\"T18 Pro\",\"model\":\"AC1200\",\"softVer\":\"V2.0.0#33_20220409113447\",\"hardVer\":\"V1\",\"wwlinkerVer\":\"V1.0.0\",\"devToken\":\"000e1faba8bd130101000994888\"},\"seq\":\"0955231818\"}"
	r1 := &Request1{}
	err := JsonIter.Unmarshal([]byte(str1), &r1)

	fmt.Println(r1)
	fmt.Println(err)

	r2 := &Request1{}
	err = JsonIter.Unmarshal([]byte(str2), &r2)

	fmt.Println(r2)
	fmt.Println(err)

	r3 := &Request1{}
	err = JsonIter.Unmarshal([]byte(str3), &r3)

	fmt.Println(r2)
	fmt.Println(err)

	a1 := &AuthRx{}
	err = JsonIter.Unmarshal([]byte(str1), &a1)
	fmt.Println(a1.Data)
	fmt.Println(err)
}
