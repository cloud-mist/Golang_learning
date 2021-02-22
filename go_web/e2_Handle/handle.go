// Package main provides ...
package main

import "net/http"

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

// 创建web服务器
func main() {
	mh := myHandler{}
	// 方法一：http.ListenAndServe("localhost:8080", nil)

	/* 方法二：更灵活
	    http.Server是一个struct,
		Addr字段，网络地址，若为空表示80端口;
		Handler字段，若为nil,是DefaultServeMux
	*/
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
}
