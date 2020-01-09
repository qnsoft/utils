package main

import (
	"github.com/astaxie/beego"
)

func main() {
	//-----------------swagger相关------------------------//
	// if beego.BConfig.RunMode == "dev" {
	// 	beego.BConfig.WebConfig.DirectoryIndex = true
	// 	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	// }
	//-----------------基本业务处理------------------------//
	// 定时任务
	//task.Init()
	//redis信息
	//initRedis()
	//计划任务
	//initJobs()go
	//-----------------运行beego------------------------//
	beego.Run()
}
