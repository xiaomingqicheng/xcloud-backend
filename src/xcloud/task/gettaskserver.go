package task


import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/astaxie/beego"
)


var taskserver *machinery.Server
func GetTaskserver() *machinery.Server {
	taskserver ,_ = machinery.NewServer(&config.Config{
			Broker:                  "redis://localhost:6379",
			ResultBackend:           "redis://localhost:6379",
		})
	err := taskserver.RegisterTasks(map[string]interface{}{
		"build": Build,
	})
	if err != nil {
		beego.Info(err,"oooooooooooooooooooooooo")
	}
	return taskserver
}