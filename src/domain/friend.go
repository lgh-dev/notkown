package domain

//好友信息
type Friend struct {
	id          string //唯一标识
	netAddress  string //网络地址
	nickName    string //昵称
	headPortait string //头像
	status      string //状态
	enable      bool   //启用、禁用
}
