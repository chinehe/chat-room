package service

import (
	"github.com/gorilla/websocket"
	"log"
)

// Login 用户登录
func (user *User) Login() {
	err := Login(user)
	if err != nil {
		log.Printf("user(%s:%s) login error:%v", user.IP, user.Name, err)
		user.CloseConn()
		return
	}
	log.Printf("user(%s:%s) login success.", user.IP, user.Name)
	HandleMessage(user, []byte("-------------login---------------------"))
	// 监听
	user.Listen()
}

// Listen 监听
func (user *User) Listen() {
	for {
		// 读取消息
		_, message, err := user.WebsocketConn.ReadMessage()
		if err != nil {
			log.Printf("user(%s:%s) listen message error:%v.", user.IP, user.Name, err)
			break
		}
		// 消息处理
		HandleMessage(user, message)
	}
	user.Logout()
}

// Write 写入消息
func (user *User) Write(message []byte) {
	err := user.WebsocketConn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Printf("user(%s:%s) write message error:%v.", user.IP, user.Name, err)
		user.Logout()
	}
}

// CloseConn 关闭连接
func (user *User) CloseConn() {
	err := user.WebsocketConn.Close()
	if err != nil {
		log.Printf("user(%s:%s) close connection error:%v", user.IP, user.Name, err)
	}
	log.Printf("user(%s:%s) close connection success.", user.IP, user.Name)
}

// Logout 登出
func (user *User) Logout() {
	err := Logout(user)
	if err != nil {
		log.Printf("user(%s:%s) logout error:%v", user.IP, user.Name, err)
	}
	log.Printf("user(%s:%s) logout success.", user.IP, user.Name)
	HandleMessage(user, []byte("-------------logout---------------------"))
	err = user.WebsocketConn.Close()
	if err != nil {
		log.Printf("user(%s:%s) close connection error:%v", user.IP, user.Name, err)
	}
}
