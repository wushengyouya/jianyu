package main

import "fmt"

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	// &User 逃逸到堆上
	return &User{
		ID:     13746734,
		Name:   "mikasa",
		Avatar: "www.baiduc.om",
	}
}

func main() {
	GetUserInfo()
	str := new(string)
	*str = "miksa"
	// 不执行输出，str不会逃逸到堆上
	fmt.Println(str)
}
