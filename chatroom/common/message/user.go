package message

// 定义一个用户的结构体
type User struct {
	// 为了序列化和反序列化成功
	// 必须保证json字符串的key和tag匹配成功
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}