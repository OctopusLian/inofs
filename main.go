package main

import (
	_ "inofs/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	username := beego.AppConfig.String("username")
	pwd := beego.AppConfig.String("pwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	db := beego.AppConfig.String("db")

	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8mb4"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	orm.RunCommand()
	beego.Run()
}
