package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "HTTP Hello")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// 升级到websocket连接
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close(websocket.StatusInternalError, "内部出错了")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		// 读取客户端的数据
		var v any
		err = wsjson.Read(ctx, conn, &v)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("接收到客户端: %v\n", v)

		// 写数据给客户端
		err = wsjson.Write(ctx, conn, "Hello WebSocket Client")
		if err != nil {
			log.Println(err)
			return
		}
		conn.Close(websocket.StatusNormalClosure, "")
	})
	log.Fatal(http.ListenAndServe(":2021", nil))
}
