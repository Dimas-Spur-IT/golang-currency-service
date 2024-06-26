package main

import (
	_ "golang-currency-service/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbUser, _ := web.AppConfig.String("db_user")
	dbPass, _ := web.AppConfig.String("db_pass")
	dbHost, _ := web.AppConfig.String("db_host")
	dbPort, _ := web.AppConfig.String("db_port")
	dbName, _ := web.AppConfig.String("db_name")
	dataSource := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True"

	orm.RegisterDataBase("default", "mysql", dataSource)
}

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
