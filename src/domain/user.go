package domain

//用户信息
type User struct {
	Id          string //唯一标识
	NickName    string //昵称
	UserName    string //用户名称
	Password    string //密码
	HeadPortait string //头像
	Status      string //状态，在线、离线
	Enable      bool   //启用、禁用
}
