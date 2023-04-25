package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//定义请求url
	apiUrl := "http://localhost:8005/req/post"
	//定义请求体以及header
	contentType := "application/json"
	data := `{"name": "httpPost-root"}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		panic("响应失败" + err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("返回数据为：", string(b))
}
