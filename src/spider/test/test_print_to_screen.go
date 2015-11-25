package main

import (
	"spider/core/engine"
)

func main() {
	var urls = []string{"http://m.qq.com", "http://m.baidu.com"}
	engine.NewEngine("print_to_screen").SetStartUrls(urls).Start()
}