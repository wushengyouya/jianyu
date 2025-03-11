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

	userHasToken := logic.NewUser(conn, token, nickname, req.RemoteAddr)

	// 2.开启给用户发送消息的goroutine
	go userHasToken.SendMessage(req.Context())

	// 3.给当前用户发送欢迎信息
	userHasToken.MessageChannel <- logic.NewWelcomeMessage(userHasToken)

	// 避免token泄漏
	tmpUser := *userHasToken
	user := &tmpUser
	user.Token = ""

	// 给所有用户告知新用户的到来
	msg := logic.NewUserEnterMessage(user)
	logic.Broadcaster.Broadcast(msg)

	// 4.将该用户加入广播器的用户列表中
	logic.Broadcaster.UserEntering(user)
	log.Println("user: ", nickname, "joins chat")

	// 5.接收用户发送来的消息
	err = user.ReceiverMessage(req.Context())

	// 6.用户离开
	logic.Broadcaster.UserLeaving(user)
	msg = logic.NewUserLeaveMessage(user)
	logic.Broadcaster.Broadcast(msg)
	log.Println("user: ", nickname, "leaves chat")

}
