package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/* 1.将url传给http.Get函数，返回http.Response对象及错误
 * 2.response对象是一个struct,body字段表示页面内容，

 */
func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // 一旦main函数退出，就释放链接，defer是延迟
	body, err := ioutil.ReadAll(response.Body) // 读取响应中的所有数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body)) // 数据转换成字符串
	
}
