package data

//用于存储用户信息的结构体
type User struct {
	Id     int
	Name   string
	Passwd string
	Token  string
}

//用于存储用户的切片
var Slice []User

//用于临时存储用户登录信息的Map
var State = make(map[string]interface{})
