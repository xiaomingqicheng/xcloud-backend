package main


import (
	_ "xcloud/routers"
	"github.com/astaxie/beego"
	"xcloud/task"
)



func main() {
	//if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	go	task.StartWorker()
	beego.Run()

}
