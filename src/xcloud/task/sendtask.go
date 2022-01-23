package task

import (
	task "github.com/RichardKnop/machinery/v1/tasks"
	"github.com/astaxie/beego"
)

func SendTask(taskname string , arg1 string) {
	tasks := task.Signature{
		Name: taskname,
		Args: []task.Arg{
			{
				Type:  "string",
				Value: arg1,
			},
		},
	}
	res, _ := taskserver.SendTask(&tasks)
	//beego.Info(res.GetState().TaskUUID,err,"yuyyyyyyyyyyyyyyyyyyyyyy")
	bk,_:=taskserver.GetBackend().GetState(res.GetState().TaskUUID)
	beego.Info(bk.State,"tttttttttttttttttttt")
}

