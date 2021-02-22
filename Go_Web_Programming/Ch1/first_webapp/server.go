// Package main provides Hello World
package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func handler(writer http.ResponseWriter, request *http.Request) {
	// 从Request结构中提取相关信息，创建一个HTTP响应，通过ResponseWriter接口将响应返回给客户端。
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler) // 把处理器函数handler设置成根被访问时的处理器。
	http.ListenAndServe(":8080", nil)
}
