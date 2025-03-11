package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

type User struct {
	ID             int
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

type Message struct {
	OwnerID int
	Content string
}

func (u User) String() string {
	return strconv.Itoa(u.ID) + u.Addr
}

var (
	// 广播专用的用户普通消息 channel，缓冲是尽可能避免出现异常情况堵塞，这里简单给了 8，具体值根据情况调整
	messageChannel  = make(chan *Message, 8)
	enteringChannel = make(chan *User)
	leavingChannel  = make(chan *User)
)

func main() {
	listener, err := net.Listen("tcp", ":2020")
	if err != nil {
		panic(err)
	}
	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster 用于记录聊天室的用户,并进行消息广播
// 1.新用户进来 2.用户普通消息 3.用户离开
func broadcaster() {
	users := make(map[*User]struct{})
	for {
		select {
		case user := <-enteringChannel:
			// 新用户进入
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users, user)
			// 避免goroutine泄漏
			close(user.MessageChannel)
		case msg := <-messageChannel:
			// 给所有在线用户法消息
			for user := range users {
				// 过滤自己的消息
				if user.ID == msg.OwnerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 1.新用户进来，构建该用户的实例
	user := &User{
		ID:             GenUserID(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	// 2.当前在一个新的goroutine中，用来进行读操作，因此需要开一个goroutine用户写操作
	// 读写goroutine可以通过channel进行通信
	go sendMessage(conn, user.MessageChannel)
	// 3.给当前用户发送欢迎信息，给所有用户告知新用户的到来
	user.MessageChannel <- "Welcome, " + user.String()
	messageChannel <- &Message{OwnerID: user.ID, Content: "user: " + strconv.Itoa(user.ID) + " has enter"}
	// 4.将该记录到到全局的用户列表中，避免用锁
	enteringChannel <- user

	// 控制超时用户踢出
	userActive := make(chan struct{})
	go func() {

		d := time.Minute * 5
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	// 5.循环读取用户的输入
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- &Message{OwnerID: user.ID, Content: strconv.Itoa(user.ID) + ":" + input.Text()}
		// 用户活跃
		userActive <- struct{}{}
	}

	// 6.用户离开
	leavingChannel <- user
	messageChannel <- &Message{OwnerID: user.ID, Content: "user:" + strconv.Itoa(user.ID) + " has left"}
}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func GenUserID() int {
	return int(time.Now().Unix())
}
