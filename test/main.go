package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("%+v", GetAddress("27.27.226.137"))
}

func GetAddress(ip string) string {
	url := fmt.Sprintf("https://restapi.amap.com/v3/ip?key=64f192da1cb2edfddb573f5afa9097a6&ip=%s", ip)
	response, err := GetRequest(url)
	if err != nil {
		return "未知地址"
	}
	type IpAddress struct {
		Province string `json:"province"`
		City     string `json:"city"`
	}
	data := IpAddress{}
	if err := json.Unmarshal([]byte(response), &data); err != nil {
		return "未知地址"
	}

	return data.Province + data.City
}

func GetRequest(url string) (data string, err error) {
	client := &http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return
}
