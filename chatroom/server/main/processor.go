package main

import (
	"fmt"
	"io"
	"main/chatroom/common/message"
	"main/chatroom/server/process"
	"main/chatroom/server/utils"
	"net"
)

type Processor struct {
	Conn net.Conn
}

// 编写一个ServerProcessMes 函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
		case message.LoginMesType :
			// 处理登陆的逻辑
			// 创建用户登陆的实例
			up := &process2.UserProcessor{
				Conn: this.Conn,
			}
			up.SeverProcessLogin(mes)
		case message.RegisterMesType:
			// 处理注册的逻辑
			up := &process2.UserProcessor{
				Conn: this.Conn,
			}
			up.SeverProcessRegister(mes)
		default:
			fmt.Println("消息类型不存在，无法处理")
	}
	return
}

func (this *Processor) mainProcess() (err error){
	//循环读取客户端发送的信息
	for {

		// 将读取数据包，直接封装成一个函数readPkg(),返回Message，err
		// 创建Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出")
				return nil
			}else {
				fmt.Println("readPkg err=", err)
				return err
			}
		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil{
			fmt.Println(err)
			return err
		}
	}
}