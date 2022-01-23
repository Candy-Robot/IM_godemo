package process2

import (
	"encoding/json"
	"fmt"
	"main/chatroom/common/message"
	"main/chatroom/server/model"
	"main/chatroom/server/utils"
	"net"
)

type UserProcessor struct {
	// 字段分析
	Conn net.Conn
}

// 编写一个函数 severProcessLogin 函数，专门处理登陆请求
func (this *UserProcessor) SeverProcessLogin(mes *message.Message) (err error) {
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

	// 拿到的账号密码应该放到数据库中进行比对
	// 使用 model.MyUserDao到redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if user != nil{
		fmt.Println(user, "用户登陆成功 ")
	}
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS{// 不合法 500 表示该用户不存在
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 404
			loginResMes.Error = "服务器内部错误"
		}
	}else {// 合法 code 200 表示登陆成功
		loginResMes.Code = 200
	}
	fmt.Println(loginResMes.Error)


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
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
// 专门处理注册请求
func (this *UserProcessor) SeverProcessRegister(mes *message.Message) (err error) {
	// 先从mes 中取出mes.data 并且反序列化成registerMess
	var registerMess message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMess)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//fmt.Println(registerMess.User)
	// 1、声明一个返回客户端的消息
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	// 2、再声明一个 RegisterResMes
	var registerResMes message.RegisterResMes
	// 到数据库完成注册
	err = model.MyUserDao.Register(&registerMess.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误信息"
		}
	}else {
		registerResMes.Code = 200
	}

	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal loginResMes fail", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal resMes fail", err)
		return
	}
	// 发送数据 writePkg 函数
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}