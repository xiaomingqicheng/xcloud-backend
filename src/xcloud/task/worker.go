package task

import (
	"github.com/astaxie/beego"

)


func StartWorker() error {
	taskserver:=GetTaskserver()
	worker := taskserver.NewWorker("machinery_worker", 10)
	if err := worker.Launch(); err != nil {
		beego.Info(err,"pppppppppppppppppppppppp")
		return err
	}
	return nil
}

