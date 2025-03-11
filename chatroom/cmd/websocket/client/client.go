package main

import (
	"context"
	"fmt"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:2021/ws", nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close(websocket.StatusInternalError, "内部错误")

	err = wsjson.Write(ctx, conn, "Hello WebSocker Server")
	if err != nil {
		panic(err)
	}
	var v any
	err = wsjson.Read(ctx, conn, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("接收到服务端的响应: %v\n", v)
	conn.Close(websocket.StatusNormalClosure, "")
}
