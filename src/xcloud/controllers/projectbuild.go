package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"xcloud/task"
	//"encoding/json"
)


// Operations about object
type ProjectbuildController struct {
	beego.Controller
}




func (this *ProjectbuildController) Ws(){

	// http升级websocket协议的配置
	var wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, _ := wsUpgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	for {
		_, msgStr, err := conn.ReadMessage()
		if err != nil {
			break
		}
		beego.Info("WS-----------receive: "+string(msgStr))
		task.SendTask("build","pppppppppppppppppppppppppppppppp")
	}

}



func (this *ProjectbuildController) GetBuildResult()  {

	taskuuid := this.GetString("taskuuid")
	beego.Info(taskuuid,"yyyyyyyyyyyyyyyyyy")
	state:=task.GetResult(taskuuid)
	this.Data["json"] = state
	this.ServeJSON()
}