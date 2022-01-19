package process

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"main/chatroom/server/utils"
	"net"
)

type UserProcessor struct {
	// 字段分析
	conn net.Conn
}

// 编写一个函数 severProcessLogin 函数，专门处理登陆请求
func (this *UserProcessor) severProcessLogin(mes *message.Message) (err error) {
	// 先从mes 中取出mes.data 并且反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	// 1、声明一个返回客户端的消息
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// 2、再声明一个 LoginResMes
	var loginResMes message.LoginResMes

	// 拿到的账号密码应该放到数据库中进行比对，写一个函数。暂时先放着
	// 如果用户id = 100，密码=123456 就是合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
		// 合法 code 200 表示登陆成功
		loginResMes.Code = 200
	}else {
		// 不合法 500 表示该用户不存在
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
	}
	// 4、将loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal loginResMes fail", err)
		return
	}
	// 5、将data赋值给resMes
	resMes.Data = string(data)
	// 对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes fail", err)
		return
	}
	// 发送数据 writePkg 函数
	// 采用分层模式， 先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.conn,
	}
	err = tf.WritePkg(data)
	return
}
