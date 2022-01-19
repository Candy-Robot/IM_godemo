package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/chatroom/common/message"
	"net"
)
// 读取接收到的数据
// 根据接收到的长度len 再接受消息本身
// 接收时要判断实际接收到的消息内容长度是否等于len
func readPkg(conn net.Conn) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	// conn.Read 在conn没有关闭的情况下，才会阻塞
	// 如果客户端
	_, err = conn.Read(buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	// 根据读到的buf[:4] 转换成一个Uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	// 根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	// 把buf反序列化成message.Message
	// 一定要用取地址符
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json.Unmarshal fail err")
		return
	}

	return
}
// 发送数据的封装函数
// 发送数据的长度，再发送消息本身
func writePkg(conn net.Conn, data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:4], pkgLen)

	n, err := conn.Write(bytes[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write() pkglen fail", err)
		return
	}
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return nil
}

// 编写一个函数 severProcessLogin 函数，专门处理登陆请求
func severProcessLogin(conn net.Conn, mes *message.Message) (err error) {
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
	err = writePkg(conn, data)
	return
}


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


// 处理与客户端通讯
func process(conn net.Conn){
	// 需要延时关闭conn
	defer conn.Close()

	//循环读取客户端发送的信息
	for {

		// 将读取数据包，直接封装成一个函数readPkg(),返回Message，err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出")
				return
			}else {
				fmt.Println("readPkg err=", err)
				return
			}
		}
		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil{
			fmt.Println(err)
			return
		}
	}

}

func main(){
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen err= ", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦链接成功，则启动一个协程和客户端保持通讯。。
		go process(conn)

	}

}