package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// 用struct将结果匹配
type Page struct {
	URL string
	Size int
}

func responseSize(url string, channel chan Page) {
	fmt.Println("Getting", url)	
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal()
	}
	channel <- Page{URL: url,  Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"https://baidu.com",
					 "https://zhihu.com",
					 "https://im.qq.com"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <- pages
		fmt.Printf("%s %d\n", page.URL, page.Size)
	}
}
