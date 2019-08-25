package main

import (
	"db"
	"domain"
	"log"
)
const tb_user="tb_user";
const tb_friend="tb_friend";
const tb_message="tb_message";
const tb_systeminfo="tb_systeminfo";

//主方法，启动
func main() {
	var tb =db.CreateLocalDB();
	user := domain.User{Id:tb.NextId(), NickName: "美丽的天使", UserName: "xiejie", Password: "123456", HeadPortait: "头像", Status: "online", Enable: true}
	tb.Add(tb_user,user.Id,user)
	var userView= tb.QueryUser(tb_user,user.Id)
	log.Println(userView)
	}


