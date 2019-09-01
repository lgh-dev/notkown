package main

import (
	"net/http"
	"controller"
	"log"
)

const tb_user="tb_user"
const tb_friend="tb_friend"
const tb_dialogues="tb_dialogues"
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
	//注册接口
	http.HandleFunc("/user/register", controller.Register)
	//登录接口
	http.HandleFunc("/user/login", controller.Login)
	//消息接收接口
	http.HandleFunc("/msg/accept", controller.AcceptMsg)
	//消息发送接口
	http.HandleFunc("/msg/send", controller.SendMsg)
	//会话查询接口
	http.HandleFunc("/dialogues/query", controller.QueryDialogues)
	//好友查询
	http.HandleFunc("/friend/query", controller.QueryFriends)
	//聊天记录查询
	http.HandleFunc("/msgs/query", controller.QueryMsgByFriendId)
	//好友查询
	http.HandleFunc("/newFriendRequest/query", controller.QueryNewFriendRequest)
	//系统信息查询
	http.HandleFunc("/SystemInfo/query", controller.QuerySystemInfos)
	//开启http监听接口
	http.ListenAndServe("127.0.0.1:7777", nil)
	log.Println("木知木知，谁也不知！")
	log.Println("Not kown, not kown. Nobody kown！")
}
