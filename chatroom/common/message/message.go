package message

const (
	LoginMesType	= "LoginMes"
	LoginResMesType	= "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// 定义两个消息

type LoginMes struct {
	UserId int `json:"user_id"`
	UserPwd string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

type LoginResMes struct {
	Code int `json:"code"`			// 返回的状态码 500 表示未注册 200 表示登陆成功
	Error string `json:"error"`		// 返回错误信息
}

type RegisterMes struct {
	// 注册的信息
	User User `json:"user"`	// 类型就是user结构体
}

type RegisterResMes struct {
	Code int `json:"code"`			// 返回的状态码 400 表示占用 200 表示注册成功
	Error string `json:"error"`		// 返回错误信息
}