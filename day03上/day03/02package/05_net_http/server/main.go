package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//定义入参的结构体
type Data struct {
	Name string `json:"name"`
}

//一个路由对应一个handler， “/api/xxx” 对应的是一个处理方法
//处理get请求
//http.ResponseWriter 用来响应内容
//*http.Request 用来拿到请求头或者请求参数
func dealGetReqHandler(w http.ResponseWriter, r *http.Request) {
	//获取请求参数
	query := r.URL.Query()
	//通过map获取
	if len(query["name"]) > 0 {
		name := query["name"][0]
		fmt.Println("通过map中切片的下标获取：", name)
	}
	//通过get方法获取
	name2 := query.Get("name") //没有的话会返回空值
	fmt.Println("通过get方式获取：", name2)
	data := &Data{Name:name2}
	//响应json数据
	json.NewEncoder(w).Encode(data)
	//w.Write([]byte(name2))
}
//处理post处理
func dealPostReqHandler(w http.ResponseWriter, r *http.Request) {
	//请求体数据
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	//转成结构体是为了可以做更多的事情，并不是转成结构体才能用json返回
	bodyStr := string(bodyByte)
	data := &Data{}
	json.Unmarshal([]byte(bodyStr), &data)
	fmt.Println(data)
	//返回响应内容
	json.NewEncoder(w).Encode(data)
	//w.Write([]byte(data.Name))
}

//启动http服务
func main() {
	//注册路由
	http.HandleFunc("/req/post", dealPostReqHandler)
	http.HandleFunc("/req/get", dealGetReqHandler)
	http.ListenAndServe(":8005", nil)
}