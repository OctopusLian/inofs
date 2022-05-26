package main

import (
	"fmt"
	_ "inofs/routers"
	"inofs/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "inofs/models"
)

func init() {
	username := beego.AppConfig.String("username")
	fmt.Println("username:",username)
	pwd := beego.AppConfig.String("pwd")
	fmt.Println("pwd:",pwd)
	host := beego.AppConfig.String("host")
	fmt.Println("host:",host)
	port := beego.AppConfig.String("port")
	fmt.Println("port:",port)
	db := beego.AppConfig.String("db")
	fmt.Println("db:",db)

	// username:pwd@tcp(ip:port)/db?charset=utf8&loc=Local
	dataSource := username + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + db + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)

	ret := fmt.Sprintf("host:%s|port:%s|db:%s", host, port, db)
	logs.Info(ret)
}

func main() {
	// 数据库命令行迁移
	orm.RunCommand()
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true

	// 未登录请求拦截
	beego.InsertFilter("/main/*", beego.BeforeRouter, utils.LoginFilter)

	// 日志
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/ions.log","separate":["error","info"]}`)
	logs.SetLogFuncCallDepth(3)
	beego.SetStaticPath("/upload", "upload")
	// 打印sql
	//orm.Debug = true
	beego.Run()
}
