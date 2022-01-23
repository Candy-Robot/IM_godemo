package main

import (
	"fmt"
	"main/chatroom/client/process"
	"os"
)

var userId int
var userPwd string
var userName string
//up := &process.UserProcess{}

func main() {

	var key int

	// 制作循环界面
	for {
		fmt.Println("--------欢迎登陆多人聊天系统-----------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）:")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			// 创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("输入用户的密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("输入用户的名字")
			fmt.Scanf("%s\n", &userName)
			fmt.Println()
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

}
