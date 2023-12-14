package main

import (
	"chat-room/constant"
	"chat-room/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// 接口
	router.HandleFunc("/user/login", controller.LoginHandle)

	// 开始监听
	log.Printf("start server:%v", constant.ServerAdder)
	if err := http.ListenAndServe(constant.ServerAdder, router); err != nil {
		log.Fatalf("start server fail,error:%v", err)
	}
}
