package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"strconv"
	"encoding/json"  
	"hello/models"
)

type APIRES struct{
	Success int
	Msg string
}

type APIRES1 struct{
	Success int
	Msg interface{}
}

type APIController struct {
	beego.Controller
}
// 查询列表
func (c *APIController) Get() {
	var res APIRES1
	res.Success = 0
	res.Msg = models.QueryTest()
	c.Data["jsonp"] = res
	c.ServeJSONP()
}
// 添加 
func (c *APIController) Post() {
	var res APIRES1
	var p models.User
	res.Success = 0      
	jsoninfo := c.GetString("jsoninfo")
	//字符串转 json
	json.Unmarshal([]byte(jsoninfo), &p)  
	msg := models.AddPost(p)
	if msg {
		res.Success = 0      
		res.Msg = "添加成功"
	}else{
		res.Success = 1
		res.Msg = "添加失败"
	}
	c.Data["jsonp"] = res
	c.ServeJSONP()
}
//  复杂方法
// func (c *APIController) Post() {
// 	var res APIRES1
// 	var p models.User
// 	res.Success = 0      
// 	jsoninfo := c.GetString("jsoninfo")
// 	json.Unmarshal([]byte(jsoninfo), &p)  
//  json数据 处理成 map类型 
// 	m := p.(map[string]interface{})  
// 	msg := models.AddPost(m)
// 	if msg {
// 		res.Success = 0      
// 		res.Msg = "添加成功"
// 	}else{
// 		res.Success = 1
// 		res.Msg = "添加失败"
// 	}
// 	c.Data["jsonp"] = res
// 	c.ServeJSONP()
// }
// 详情查询 
func (c *APIController) GetUser() {
	var res APIRES1
	Id := c.GetString("id")
	id,_:= strconv.Atoi(Id)
	res.Success = 0    
	msg := models.GetUsers(id)
	res.Msg = msg
	c.Data["jsonp"] = res
	c.ServeJSONP()
}
// 删除
func (c *APIController) DeleteUser() {
	var res APIRES1
	Id := c.GetString("id")
	id,_:= strconv.Atoi(Id)
	msg := models.DeleteUsers(id)
	if msg {
		res.Success = 0  
		res.Msg = "删除成功"
	}else{
		res.Success = 1  
		res.Msg = "删除失败"
	}
	c.Data["jsonp"] = res
	c.ServeJSONP()
}
// 更新
func (c *APIController) PutUser() {
	var res APIRES1
	var p models.User
	jsoninfo := c.GetString("jsoninfo")
	//json转结构体
	json.Unmarshal([]byte(jsoninfo), &p)
	msg := models.PutUsers(p)
	if msg {
		res.Success = 0  
		res.Msg = "更新成功"
	}else{
		res.Success = 1  
		res.Msg = "更新失败"
	}
	c.Data["jsonp"] = res
	c.ServeJSONP()
}

// 登录
func (c *APIController) Login() {
	var msg APIRES
	var data models.Login
	msg.Success = 0      
	jsoninfo := c.GetString("jsoninfo")
	json.Unmarshal([]byte(jsoninfo), &data)  
	res := models.Logins(data)
	if res {
		msg.Success = 0
		msg.Msg = "登录成功"
	}else{
		msg.Success = 1
		msg.Msg = "登录失败"
	}
	fmt.Println(msg)
	c.Data["jsonp"] = msg
	c.ServeJSONP()
}

//  注册
func (c *APIController) Register() {
	var msg APIRES
	var data models.Login
	msg.Success = 0      
	jsoninfo := c.GetString("jsoninfo")
	json.Unmarshal([]byte(jsoninfo), &data)  
	res := models.Registers(data)
	if res {
		msg.Success = 0
		msg.Msg = "插入成功"
	}else{
		msg.Success = 1
		msg.Msg = "插入失败"
	}
	fmt.Println(msg)
	c.Data["jsonp"] = msg
	c.ServeJSONP()
}