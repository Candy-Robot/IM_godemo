package main

import (
	"fmt"
	"main/chatroom/common/message"
	"net"
)

// 编写一个ServerProcessMes 函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType :
		// 处理登陆的逻辑
		severProcessLogin(conn, mes)
	case message.RegisterMesType:
		// 处理注册的逻辑
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}