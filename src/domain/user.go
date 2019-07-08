package domain

//用户信息
type User struct {
	id          string //唯一标识
	nickName    string //昵称
	userName    string //用户名称
	password    string //密码
	headPortait string //头像
	status      string //状态，在线、离线
	enable      bool   //启用、禁用
}
