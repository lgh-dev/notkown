package kits

import "log"

//异常处理
func Panic(msg string,err error)  {
	if err!=nil {
		log.Println(msg)
		log.Panic(err)
	}
}
