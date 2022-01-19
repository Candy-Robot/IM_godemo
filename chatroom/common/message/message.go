package message

const (
	LoginMesType	= "LoginMes"
	LoginResMesType	= "LoginResMes"
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