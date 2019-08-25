package main

import (
	"net/http"
	"controller"
	"log"
)

const tb_user="tb_user"
const tb_friend="tb_friend"
const tb_message="tb_message"
const tb_systeminfo="tb_systeminfo"

//主方法，启动
//func main() {
//	var tb =db.CreateLocalDB();
//	user := domain.User{Id:tb.NextId(), NickName: "美丽的天使", UserName: "xiejie", Password: "123456", HeadPortait: "头像", Status: "online", Enable: true}
//	tb.Add(tb_user,user.Id,user)
//	var userView= tb.QueryUser(tb_user,user.Id)
//	log.Println(userView)
//	}



func main() {

	http.HandleFunc("/user/register", controller.Register)
	http.HandleFunc("/user/login", controller.Login)
	http.HandleFunc("/msg/accept", controller.AcceptMsg)
	http.HandleFunc("/msg/send", controller.SendMsg)
	http.HandleFunc("/dialogues/query", controller.QueryDialogues)
	http.HandleFunc("friend/query", controller.QueryFriends)
	http.HandleFunc("msgs/query", controller.QueryMsgByFriendId)
	http.HandleFunc("/newFriendRequest/query", controller.QueryNewFriendRequest)
	http.HandleFunc("/SystemInfo/query", controller.QuerySystemInfos)
	http.ListenAndServe("127.0.0.1:7777", nil)
	log.Println("木知木知，谁也不知！")
	log.Println("Not kown, not kown. Nobody kown！")
}
