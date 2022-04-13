package main

import (
	"encoding/json"
	"fmt"
)

type Iln struct {
	ProvinceId uint `json:"province_id"`
	CityId     uint `json:"city_id"`
	DistrictId uint `json:"district_id"`
}

type Il struct {
	Status int8   `json:"status"`
	Msg    string `json:"msg"`
	Info   *Iln   `json:"info"`
}

func main() {
	str1 := `["1.0.0","1.0.1","1.0.2"]`
	str2 := `[]`
	str3 := `["20210301090101"]`

	arr1 := []string{}
	arr2 := []string{}
	arr3 := []string{}
	json.Unmarshal([]byte(str1), &arr1)
	json.Unmarshal([]byte(str2), &arr2)
	json.Unmarshal([]byte(str3), &arr3)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	str4 := `{"status":1,"msg":"ok","info":{"province_id":440000,"city_id":440300,"district_id":0}}`
	lo := Il{}
	json.Unmarshal([]byte(str4), &lo)
	fmt.Println(lo, lo.Info)
}
