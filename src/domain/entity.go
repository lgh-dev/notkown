package domain

import "time"

/**
 * @describe   : 用户信息
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:40
 * @version V1.0
 **/

//用户信息
type User struct {
	Id           string    //唯一标识
	NickName     string    //昵称
	UserName     string    //用户名称
	Password     string    //密码
	HeadPortait  string    //头像
	Status       string    //状态，在线、离线
	Enable       bool      //启用、禁用
	RegisterTime time.Time //注册时间
}

/**
 * @describe   : 好友信息
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:39
 * @version V1.0
 **/

//好友信息
type Friend struct {
	id           string    //唯一标识
	netAddress   string    //网络地址
	nickName     string    //昵称
	headPortait  string    //头像
	status       string    //状态
	enable       bool      //启用、禁用
	RegisterTime time.Time //注册时间
}

/**
 * @describe   : 消息内容
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:40
 * @version V1.0
 **/

type Message struct {
	id       string // 唯一标识
	friendId string //用户id
	msg      string //消息
	msgType  string //消息内容
}

/**
 * @class_name :
 * @describe   : 会话
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:39
 * @version V1.0
 **/

type Dialogue struct {
	lastTime time.Time //最近会话更新时间。
	friend   Friend    //好友
	lastMsg  string    //最近消息
}

/**
 * @describe   : 新好友添加请求
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:44
 * @version V1.0
 **/
type NewFriendRequest struct {
	Friend Friend    //好友信息。
	Status int16     //待添加：0；已添加：1；已拒绝。-1;
	Time   time.Time //请求时间。
}

/**
 * @describe   : 系统信息
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/25/025
 * @creat_time : 23:40
 * @version V1.0
 **/

//系统信息
type Systeminfo struct {
	Version    string //版本
	Author     string //作者
	Desc       string //描述
	newVersion string //新版本
}

/**
 * @describe   : 对前端返回数据结构
 * @creat_user : liangguohui
 * @creat_email: 18520640132@163.com
 * @creat_date : 2019/8/26/026
 * @creat_time : 0:07
 * @version V1.0
 **/

type ResultDTO struct {
	Code int64       //异常码
	Data interface{} //返回数据
	Msg  string      //返回消息
}
