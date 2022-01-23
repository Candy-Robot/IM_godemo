package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"main/chatroom/common/message"
)
// 在服务器启动后，就初始化一个userDao实例
// 操作redis的时候就可以直接使用
var (
	MyUserDao *UserDao
)


// 定义UserDao 结构体
// 完成对User 结构体对各种操作

type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建UserDao 实力
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return userDao
}

// 得到用户的账号密码 根据用户id返回一个 User实例 + err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	// 通过给定的ID 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		// 错误
		if err == redis.ErrNil { // 表示在users 哈希中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	// 需要把res 反序列化成User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil{
		fmt.Println("json.Unmarshall err=", err)
		return
	}
	return
}

// 完成登陆的校验
// 如果id和密码都正确，返回User实例
// 如果不正确， 返回错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error)  {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil{
		return
	}
	// 用户获取到了，没有判断密码
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}

// 完成注册
func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil{
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil{
		fmt.Println("Register(user *message.User)  json.Marshal err=", err)
		return
	}
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil{
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}