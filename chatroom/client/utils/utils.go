package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"main/chatroom/common/message"
	"net"
)

// 将读和写等方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte	// 传输时使用的缓冲
}


// 读取接收到的数据
// 根据接收到的长度len 再接受消息本身
// 接收时要判断实际接收到的消息内容长度是否等于len
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	//fmt.Println("读取客户端发送端数据")
	// conn.Read 在conn没有关闭的情况下，才会阻塞
	// 如果客户端
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	// 根据读到的buf[:4] 转换成一个Uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	// 根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	// 把buf反序列化成message.Message
	// 一定要用取地址符
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("json.Unmarshal fail err")
		return
	}

	return
}
// 发送数据的封装函数
// 发送数据的长度，再发送消息本身
func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var bytes [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write() pkglen fail", err)
		return
	}
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}
	return nil
}
