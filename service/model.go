package service

import (
	"github.com/gorilla/websocket"
)

// Room 聊天室
type Room struct {
	users map[string]*User // 聊天室里的成员
}

// User 用户
type User struct {
	Name          string          `json:"name"` // 用户名称
	IP            string          `json:"ip"`   // IP 地址
	WebsocketConn *websocket.Conn // 连接
}
