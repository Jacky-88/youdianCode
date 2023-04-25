package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//定义请求url
	apiUrl := "http://localhost:8005/req/get"
	//定义url参数
	data := url.Values{}
	data.Set("name", "httpGet-root")
	//组装url
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		panic(err)
	}
	u.RawQuery = data.Encode()
	fmt.Println("请求路由为: ", u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		panic("响应失败" + err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("返回数据为：", string(b))
}
