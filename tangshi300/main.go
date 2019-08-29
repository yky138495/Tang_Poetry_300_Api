package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/mattn/go-sqlite3"
	"time"
	_ "tangshi300/routers"
)


// func GetDBConn() (*sql.DB, error) {
// 	return sql.Open("sqlite3", "./data/Tangshi300.db")
// }

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./data/唐诗三百首.db")
	orm.RunSyncdb("default", false, true)


}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}


func main() {
	// admin.Run()

	// 开启 ORM 调试模式
	orm.Debug = true
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 自动建表
	orm.RunSyncdb("default", false, true)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//　CORS 跨域请求 参考:https://gocn.io/question/426
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	beego.Run()
}
