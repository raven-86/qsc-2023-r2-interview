package gin_api

import (
	"net/http"
	. "qsc/data"
	Token "qsc/token"

	"github.com/gin-gonic/gin"
)

// 判断是否存在用户
func Check(name string) (bool, User) {
	var user User
	// 如果长度为0说明尚未有用户注册
	if len(Slice) == 0 {
		return false, user
	}
	// 遍历切片
	for _, i := range Slice {
		if i.Name == name {
			return true, i
		}
	}
	return false, user
}

// 测试网络连通性
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "pong!",
		"data": nil,
	})
}

// 用户注册
func Signup(c *gin.Context) {
	// 获取用户名、密码
	name := c.PostForm("username")
	passwd := c.PostForm("password")
	// 用户名或密码为空
	if name == "" || passwd == "" {
		State["code"] = 10
		State["msg"] = "invalid username or password"
		State["data"] = nil
	} else {
		// 判断用户是否存在
		Bool, _ := Check(name)
		if Bool {
			// 用户已存在
			State["code"] = 99
			State["msg"] = "user already exists"
			State["data"] = nil
		} else {
			// 用户不存在即添加用户
			token := AddStruct(name, passwd)
			State["code"] = 0
			State["msg"] = ""
			State["data"] = nil
			State["access_token"] = token
		}
	}
	// 把状态码和注册状态返回给客户端
	c.JSON(http.StatusOK, State)
}

// 用户登录
func Signin(c *gin.Context) {
	name := c.PostForm("username")
	passwd := c.PostForm("password")
	// 判断用户是否存在
	Bool, user := Check(name)
	if Bool {
		if passwd == user.Passwd {
			// 登录成功
			State["code"] = 0
			State["msg"] = ""
			State["data"] = nil
			State["access_token"] = user.Token
		} else {
			// 密码错误
			State["code"] = 1
			State["msg"] = "wrong password"
			State["data"] = nil
		}
	} else {
		// 用户不存在
		State["code"] = 2
		State["msg"] = "user not exists"
		State["data"] = nil
	}
	// 把状态码和注册状态返回给客户端
	c.JSON(http.StatusOK, State)
}

func Checkin(c *gin.Context) {
	token := c.PostForm("access_token")
	// checkwd := c.PostForm("checkword")
	user, _ := Token.ParseToken(token)
	if user == nil {
		// token有问题
		State["code"] = 666
		State["msg"] = "invalid token"
		State["data"] = nil
	} else {
		// 签到成功
		State["code"] = 0
		State["msg"] = ""
		State["data"] = nil
		State["point"] = 1
	}
	c.JSON(http.StatusOK, State)
}

// 添加用户
func AddStruct(name string, passwd string) string {
	var user User
	user.Name = name
	user.Passwd = passwd
	user.Token = Token.GenerateToken(name, passwd)
	user.Id = len(Slice) + 1
	Slice = append(Slice, user)
	return user.Token
}
