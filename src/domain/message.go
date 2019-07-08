package domain

type Message struct {
	id       string // 唯一标识
	friendId string //用户id
	msg      string //消息
	msgType  string //消息内容
}
