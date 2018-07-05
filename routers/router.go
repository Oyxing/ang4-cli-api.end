package router
import (
    "github.com/astaxie/beego"
    "hello/controllers"  // 引入控制器
)



func init() { // init 初始化
    //固定路由
    // beego.Router("/", &controllers.MainController{})
    //自定义路由
    // beego.Router("/Login", &controllers.IndexController{})
    //自动匹配路由
    beego.AutoRouter(&controllers.APIController{})
}