package main

import (
	_ "hello/routers"  //  引入路由表
	"github.com/astaxie/beego"
	"hello/models"
)

func init(){
	models.RegisterDB()
}

func main() {
	beego.Run()
}



