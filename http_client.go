package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type IPLocationInfo struct {
	ProvinceId uint `json:"province_id"`
	CityId     uint `json:"city_id"`
	DistrictId uint `json:"district_id"`
}

type IPLocation struct {
	Status int8            `json:"status"`
	Msg    string          `json:"msg"`
	Info   *IPLocationInfo `json:"info"`
}

func main() {
	resp, err := http.Get("http://ipser.agzzu.com/ip/query-location?ip=183.14.133.237")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(body)

	location := IPLocation{}

	json.Unmarshal(body, &location)
	fmt.Println(location)
	fmt.Println(location.Info)
}
