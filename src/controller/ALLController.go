package controller

import (
	"net/http"
	"log"
	"fmt"
)

//注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "注册成功！")

}

//登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("登录成功！")

}

//查询好友列表(允许模糊查询)
func QueryFriends(w http.ResponseWriter, r *http.Request) {
	log.Println("查询好友列表成功！")

}

//查询会话列表（全量）
func QueryDialogues(w http.ResponseWriter, r *http.Request) {
	log.Println("查询会话列表成功！")

}

//查询与好友聊天信息
func QueryMsgByFriendId(w http.ResponseWriter, r *http.Request) {
	log.Println("查询与好友聊天信息成功！")

}

//发消息
func SendMsg(w http.ResponseWriter, r *http.Request) {
	log.Println("发消息成功！")

}

//接收消息
func AcceptMsg(w http.ResponseWriter, r *http.Request) {
	log.Println("接收成功！")

}

//查询添加好友请求列表（时间倒序排行）
func QueryNewFriendRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("查询添加好友请求列表成功！")

}

//查询系统信息列表
func QuerySystemInfos(w http.ResponseWriter, r *http.Request) {
	log.Println("查询系统信息列表成功！")

}
