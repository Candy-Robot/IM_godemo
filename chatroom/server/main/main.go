package main

import (
	"fmt"
	"main/chatroom/server/model"
	"net"
	"time"
)

// 处理与客户端通讯
func process(conn net.Conn){
	// 需要延时关闭conn
	defer conn.Close()

	// 调用总控
	processor := &Processor{
		Conn: conn,
	}
	err := processor.mainProcess()
	if err != nil {
		fmt.Println("客户端与服务器通讯协程错误，err = ", err)
		return
	}
}
// 编写一个函数完成对userDao的初始化任务
func initUserDao() {
	// 初始化userDao要在初始化链接池之后。这样才有pool
	model.MyUserDao = model.NewUserDao(pool)
}
func main(){
	// 当服务器启动时，就初始化redis当链接池
	initPool("localhost:6379", 8,0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器new在8889端口监听~~~!!")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	// listen 欢迎嵌套字
	defer listen.Close()
	if err != nil {
		fmt.Println("net listen err= ", err)
		return
	}
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept() // conn 链接嵌套字
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦链接成功，则启动一个协程和客户端保持通讯。。
		go process(conn)

	}

}