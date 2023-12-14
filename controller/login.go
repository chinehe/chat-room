package controller

import (
	"chat-room/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var websocketUpgrader = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func LoginHandle(resp http.ResponseWriter, req *http.Request) {
	// 获取用户名
	name := req.URL.Query().Get("name")
	if name == "" {
		_, _ = resp.Write([]byte("Invalid UserName"))
		return
	}
	// websocket 连接
	websocketConn, err := websocketUpgrader.Upgrade(resp, req, nil)
	if err != nil {
		log.Printf("webscoket upgrade error:%v", err)
		return
	}
	// 用户登录
	user := service.User{
		Name:          name,
		IP:            req.RemoteAddr,
		WebsocketConn: websocketConn,
	}
	user.Login()
}
