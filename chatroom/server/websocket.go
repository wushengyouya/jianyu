package server

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/wushengyouya/chatroom/logic"
)

func WebSocketHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		log.Println("websocket accept error: ", err)
		return
	}

	// 1.新用户进来，构建用户的实例
	token := req.FormValue("token")
	nickname := req.FormValue("nickname")

	// 昵称长度校验
	if l := len(nickname); l < 2 || l > 20 {
		log.Println("nickname illegal: ", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("非法昵称，昵称长度2-20"))
		conn.Close(websocket.StatusUnsupportedData, "nickname illegal")
		return
	}

	// 判断是否同名
	if !logic.Broadcaster.CanEnterRoom(nickname) {
		log.Println("昵称已经存在：", nickname)
		wsjson.Write(req.Context(), conn, logic.NewErrorMessage("改昵称已经存在"))
		conn.Close(websocket.StatusUnsupportedData, "nickname exists")
		return
	}

}
