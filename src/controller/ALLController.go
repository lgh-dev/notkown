package controller

import (
	"domain"
	"encoding/json"
	"net/http"
)

//注册接口
func Register(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","注册成功！"})
	w.Write(resultDTO)
}

//登录接口
func Login(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","登录成功！"})
	w.Write(resultDTO)
}

//查询好友列表(允许模糊查询)
func QueryFriends(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","查询好友列表成功！"})
	w.Write(resultDTO)

}

//查询会话列表（全量）
func QueryDialogues(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","查询会话列表成功！"})
	w.Write(resultDTO)
}

//查询与好友聊天信息
func QueryMsgByFriendId(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","查询与好友聊天信息成功！"})
	w.Write(resultDTO)
}

//发消息
func SendMsg(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","发消息成功！"})
	w.Write(resultDTO)
}

//接收消息
func AcceptMsg(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","接收成功！"})
	w.Write(resultDTO)
}

//查询添加好友请求列表（时间倒序排行）
func QueryNewFriendRequest(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","查询添加好友请求列表成功！"})
	w.Write(resultDTO)

}

//查询系统信息列表
func QuerySystemInfos(w http.ResponseWriter, r *http.Request) {
	resultDTO,_:=json.Marshal(domain.ResultDTO{1,"","查询系统信息列表成功！"})
	w.Write(resultDTO)
}
