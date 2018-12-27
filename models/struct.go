package models

import (
	"fmt"
	//"strconv"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
)

// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
    Phone string `orm:"size(11)"`
    Email string `orm:"size(100)"`
	Age int `orm:"size(100)"`
	CompanyId int `orm:"size(100)"`
	Balance float64 `orm:"size(100)"`

}

// Model Struct
type Login struct {
    Id   int
    Name string `orm:"size(100)"`
    Password string `orm:"size(11)"`
}

func RegisterDB() {
	orm.Debug = true
	// set default database
	// 参数1        数据库的别名，用来在 ORM 中切换数据库使用
	// 参数2        driverName
	// 参数3        对应的链接字符串
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	// maxIdle := 30
	// maxConn := 30
	// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)

    orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8", 30)

    // register model
    orm.RegisterModel(new(User),new(Login))

    // create table
    orm.RunSyncdb("default", false, true)
}
// 添加 map[string]interface{}  //  map 类型的书写
func AddPost(AddData User) bool {
	var res bool
	o := orm.NewOrm()
	//类型断言 只对interface{} 有用 断言成对应类型的 
	// user := User{Name:AddData["name"].(string),Phone:AddData["phone"].(string),Email:AddData["email"].(string),Balance:AddData["balance"].(float64)}
	id, err := o.Insert(&AddData)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err == nil{
		res = true
	}
	return res
}

// 查询列表
func QueryTest() interface{}{
	var maps []orm.Params
	o := orm.NewOrm()
	num, _ := o.QueryTable("user").Values(&maps)
	fmt.Println(num)
	return maps
}

// 查询详情
func GetUsers(Id int) interface{} {
	var res bool
	o := orm.NewOrm()
	u := User{Id: Id}
    err := o.Read(&u)
    fmt.Println(err)
	if err == nil{
		res = true
	}
	fmt.Println(res)
	return u
}
// 删除PutUsers
func DeleteUsers(Id int) bool {
	var res bool
	o := orm.NewOrm()
	u := User{Id: Id}
	num, err := o.Delete(&u)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	if err == nil{
		res = true
	}
	return res
}
// 更改
func PutUsers(AddData User) bool {
	var res bool
	o := orm.NewOrm()
	num, err := o.Update(&AddData)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	if err == nil{
		res = true
	}
	return res
}

//  登录
func Logins(LoginSql Login) bool {
	var res bool
	o := orm.NewOrm()
	u := Login{}
	err := o.QueryTable("login").Filter("name", LoginSql.Name).One(&u)
	//err := o.Read(&u)
	if err == nil {
		if u.Password == LoginSql.Password {
			res = true
		}
	}
	return res
}

//  注册
func Registers(LoginSql Login) bool {
	var res bool
	o := orm.NewOrm()
	u := Login{}
	err := o.QueryTable("login").Filter("name", LoginSql.Name).One(&u)
	if err != nil{
		id, err := o.Insert(&LoginSql)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		if err == nil{
			res = true
		}
	}
	
	return res
}