package main

import (
	"flag"
	"log"
)

// 简单案例
func main() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")

	flag.Parse()

	log.Printf("name: %s", name)

}
