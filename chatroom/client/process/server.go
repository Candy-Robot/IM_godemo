package process

import (
	"fmt"
	"main/chatroom/server/utils"
	"net"
	"os"
)

// 显示登陆成功后的界面
func ShowMenu() {
	for {
		fmt.Println("----------登陆成功---------")
		fmt.Println("------1. 显示在线用户列表----")
		fmt.Println("------2. 发送信息------")
		fmt.Println("------3. 信息列表------")
		fmt.Println("------4. 退出系统------")
		fmt.Println("请选择（1-4）")
		var key int
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送信息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你输入的不正确")
		}
	}
}

// 和服务器端保持通讯
func serverProcessMes(conn net.Conn){
	// 创建一个transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		// 读取到消息，下一步处理
		fmt.Println(mes)
	}
}