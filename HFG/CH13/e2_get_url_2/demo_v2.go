package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/* 1.将url传给http.Get函数，返回http.Response对象及错误
2.response对象是一个struct,body字段表示页面内容，

*/
func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() // 一旦main函数退出，就释放链接，defer是延迟
	body, err := ioutil.ReadAll(response.Body) // 读取响应中的所有数据
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func main() {
	go responseSize("https://example.com")
	go responseSize("https://baidu.com")
	go responseSize("https://zhihu.com")
	time.Sleep(time.Second * 5) // main goro 休眠5s
}

// 结果：输出内容是混合的，但速度快了不少，但是我们需要等待main中time.Sleep的调用完成，仍需5s.
// 所以，休眠不是解决问题的最佳方式。

/*
Getting https://example.com
Getting https://zhihu.com
Getting https://baidu.com
298584
50504
1256

[Process exited 0]
*/

