package logic

import (
	"time"

	"github.com/spf13/cast"
)

const (
	MsgTypeNormal    = iota // 普通用户消息
	MsgTypeWelcome          // 当前用户欢迎消息
	MsgTypeUserEnter        // 用户进入
	MsgTypeSystem           // 系统消息
	MsgTypeError            // 错误消息
	MsgTypeUserLeave        // 用户退出
	MsgTypeUserList         // 发送当前用户列表
)

// 给用户发送的消息
type Message struct {
	// 哪个用户发送的消息
	User           *User            `json:"user"`
	Type           int              `json:"type"`
	Content        string           `json:"content"`
	MsgTime        time.Time        `json:"msg_time"`
	ClientSendTime time.Time        `json:"client_send_time"`
	Users          map[string]*User `json:"users"`
}

func NewMessage(user *User, content, clientTime string) *Message {
	message := &Message{
		User:    user,
		Type:    MsgTypeNormal,
		Content: content,
		MsgTime: time.Now(),
	}
	if clientTime != "" {
		message.ClientSendTime = time.Unix(0, cast.ToInt64(clientTime))
	}
	return message
}

func NewWelcomeMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeWelcome,
		Content: user.NickName + "你好，欢迎加入聊天室",
		MsgTime: time.Now(),
	}
}

func NewUserEnterMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserEnter,
		Content: user.NickName + "加入了聊天室",
		MsgTime: time.Now(),
	}
}

func NewUserLeaveMessage(user *User) *Message {
	return &Message{
		User:    user,
		Type:    MsgTypeUserLeave,
		Content: user.NickName + " 离开了聊天室",
		MsgTime: time.Now(),
	}
}

func NewErrorMessage(content string) *Message {
	return &Message{
		User:    System,
		Type:    MsgTypeError,
		Content: content,
		MsgTime: time.Now(),
	}
}
