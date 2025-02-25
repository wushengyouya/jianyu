package main

import (
	"flag"
	"log"
)

// 子命令
var name string

func main() {
	// 先解析，便于后续命令调用
	flag.Parse()
	// 返回命令行的所有参数
	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])

	}

	log.Printf("name: %s", name)
}
