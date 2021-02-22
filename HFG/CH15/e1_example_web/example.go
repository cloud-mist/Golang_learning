// Package main provides ...
package main

import (
	"log"
	"net/http"
)

//http.res用于想浏览器响应写入数据，一个指针表示浏览器的请求
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello Web!")
	_, err := writer.Write(message) //Write方法添加数据
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)            //告诉应用程序每当收到/hello结尾的url请求，调用viewHandle
	err := http.ListenAndServe("localhost:8080", nil) // 这个函数启动web服务器
	log.Fatal(err)                                    //除非遇到错误，才终止运行
}
