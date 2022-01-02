package controllers

import (
	"github.com/astaxie/beego"
	"k8s.io/client-go/tools/remotecommand"
	"encoding/json"
	"xcloud/util"
	"github.com/gorilla/websocket"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"fmt"
)

type ContainerWebshellController struct {
	beego.Controller
}



// ssh流式处理器
type streamHandler struct {
	wsConn      *WsConnection
	resizeEvent chan remotecommand.TerminalSize
}


type xtermMessage struct {
	MsgType string `json:"type"`  // 类型:resize客户端调整终端, input客户端输入
	Input   string `json:"input"` // msgtype=input情况下使用
	Rows    uint16 `json:"rows"`  // msgtype=resize情况下使用
	Cols    uint16 `json:"cols"`  // msgtype=resize情况下使用
}


// executor回调获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	ret := <-handler.resizeEvent
	size = &ret
	return
}

// executor回调读取web端的输入
func (handler *streamHandler) Read(p []byte) (size int, err error) {
	var (
		msg      *WsMessage
		xtermMsg xtermMessage
	)

	// 读web发来的输入
	if msg, err = handler.wsConn.WsRead(); err != nil {
		return
	}

	// 解析客户端请求
	if err = json.Unmarshal(msg.Data, &xtermMsg); err != nil {
		return
	}
	beego.Info(xtermMsg.Input,"YYYYYYYYYYY")
	//web ssh调整了终端大小
	if xtermMsg.MsgType == "resize" {
		// 放到channel里，等remotecommand executor调用我们的Next取走
		handler.resizeEvent <- remotecommand.TerminalSize{Width: xtermMsg.Cols, Height: xtermMsg.Rows}
	} else if xtermMsg.MsgType == "input" { // web ssh终端输入了字符
		// copy到p数组中
		size = len(xtermMsg.Input)
		copy(p, xtermMsg.Input)
		beego.Info(xtermMsg,"yyyyyyyyyyyyyyy")
	}
	return
}

// executor回调向web端输出
func (handler *streamHandler) Write(p []byte) (size int, err error) {
	var (
		copyData []byte
	)

	// 产生副本
	copyData = make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	beego.Info(copyData,"rrrrrrrrrrrrrr")
	err = handler.wsConn.WsWrite(websocket.TextMessage, copyData)
	return
}


// @router / [get]
func (this *ContainerWebshellController) WsHandler(){

	var webshell map[string]interface{}
	json.Unmarshal([]byte(this.Ctx.Input.RequestBody), &webshell)



	//clustername := webshell["clustername"]

	var (
		wsConn    *WsConnection
		restConf  rest.Config
		//sshReq    *rest.Request
		pod       string
		namespace string
		container string
		executor  remotecommand.Executor
		handler   *streamHandler
		err       error
	)
	namespace = this.GetString("namespace")
	pod = this.GetString("pod")
	container = this.GetString("container")

	// 得到websocket长连接
	if wsConn, err = InitWebsocket(this.Ctx.ResponseWriter, this.Ctx.Request); err != nil {
		return
	}



	// 获取k8s rest client配置


	var ClientSet kubernetes.Clientset
	ClientSet = util.Getclient("集群1")
	//sshReq = new(rest.Request)
	sshReq := ClientSet.CoreV1().RESTClient().Post().Resource("pods").Name(pod).Namespace(namespace).SubResource("exec")
	sshReq.VersionedParams(&v1.PodExecOptions{
			Container: container,
			Command:   []string{"bash"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	 restConf = util.Gettnlsconfig("集群1")

	// 创建到容器的连接
	if executor, err = remotecommand.NewSPDYExecutor(&restConf, "POST", sshReq.URL()); err != nil {
		goto END
	}

	// 配置与容器之间的数据流处理回调
	handler = &streamHandler{wsConn: wsConn, resizeEvent: make(chan remotecommand.TerminalSize)}
	beego.Info(*handler,"uuuuuuuuuuuuuuu")
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		TerminalSizeQueue: handler,
		Tty:               true,
	}); err != nil {
		beego.Info(err,"5555555555555555")
		goto END
	}
	return

END:
	fmt.Println(err,"ooooooooooooooooooo")
	wsConn.WsClose()
}