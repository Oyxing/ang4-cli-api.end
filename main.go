package main

import (
	"hello/models"
	_ "hello/routers" //  引入路由表

	"github.com/astaxie/beego"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
