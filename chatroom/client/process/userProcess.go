package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"main/chatroom/server/utils"
	"net"
)

type UserProcess struct {

}

// 关联一个用户登陆的方法
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
	}
	// 延时关闭
	defer conn.Close()

	// 准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	// 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//fmt.Println(string(data))
	// 将data赋给mes.data字段
	mes.Data = string(data)

	// 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	// 创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	// 将发送数据用函数封装
	err = tf.WritePkg(data)

	// 还需要处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err=", err)
	}
	// 将reMes反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200{
		// 还需要再客户端启动一个协程
		// 该协程保持和服务器端的通许，如果服务器有数据推送给客户端
		// 则接收并且显示在客户端的终端
		go serverProcessMes(conn)
		// 显示登陆成功的菜单
		ShowMenu()
	}else {
		fmt.Println(loginResMes.Error)
	}
	return nil
}
// 关联一个用户注册代码
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
	}
	// 延时关闭
	defer conn.Close()

	// 准备通过conn发送信息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType

	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	// registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//fmt.Println(string(data))

	// 将data赋给mes.data字段
	mes.Data = string(data)

	// 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	// 创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	// 将发送数据用函数封装
	err = tf.WritePkg(data)

	// 还需要处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err=", err)
	}

	// 将reMes反序列化成loginResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200{
		fmt.Println("注册成功，重新登陆")
	}else {
		fmt.Println(registerResMes.Error)
	}
	return nil
}
