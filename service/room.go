package service

import (
	"fmt"
	"log"
)

// chatRoom 聊天室实例
var chatRoom *Room

func init() {
	chatRoom = &Room{
		users: make(map[string]*User),
	}
	log.Printf("init chat room success.")
}

// Login 登录
func Login(user *User) error {
	if _, ok := chatRoom.users[user.Name]; ok {
		return fmt.Errorf("user(%s) already exists", user.Name)
	}
	chatRoom.users[user.Name] = user
	return nil
}

// Logout 登录
func Logout(user *User) error {
	if _, ok := chatRoom.users[user.Name]; !ok {
		return fmt.Errorf("user(%s) already logout", user.Name)
	}
	delete(chatRoom.users, user.Name)
	return nil
}

// HandleMessage 消息处理
func HandleMessage(user *User, message []byte) {
	messageFormat := "<span style='color: blue'>[ %s ]</span>  <span style='color: rgb(255,0,38)'>< %s ></span> : %s"
	message = []byte(fmt.Sprintf(messageFormat, user.IP, user.Name, message))
	log.Printf("user(%s:%s) send:%s",user.IP,user.Name,string(message))
	for _, u := range chatRoom.users {
		u.Write(message)
	}
}
