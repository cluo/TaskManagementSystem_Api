package main

import (
	_ "TaskManagementSystem_Api/routers"
	"runtime"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
	}))
	beego.Run()
}

// http://oauth.hisign.top:6002/o/authorize/?response_type=code&client_id=hAHln3ZKrnPf8odTUdkizuSSbIP3CvRzNY0zBZXD
// http://oauth.hisign.top:6002/o/authorize/?response_type=password&client_id=hAHln3ZKrnPf8odTUdkizuSSbIP3CvRzNY0zBZXD
