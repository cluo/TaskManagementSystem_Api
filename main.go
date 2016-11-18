package main

import (
	"TaskManagementSystem_Api/models/common"
	_ "TaskManagementSystem_Api/routers"
	"runtime"

	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	err := common.InitRedis()
	if err != nil {
		log.Fatalln("Redis连接初始化失败！")
	}
	defer common.Bunt.CloseDB()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "X-Auth-Token"},
	}))
	beego.Run()
}

// http://oauth.hisign.top:6002/o/authorize/?response_type=code&client_id=ddvaC9RcIAiWYlqeTYy2NbwQzFtQS2hzSZERGyUA&state=task&redirect_uri=http://task.hisign.top:6009/

// http://oauth.hisign.top:6002/o/authorize/?response_type=password&client_id=hAHln3ZKrnPf8odTUdkizuSSbIP3CvRzNY0zBZXD
