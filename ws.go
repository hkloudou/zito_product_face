package zitoproductface

import "github.com/hkloudou/websoketkit"

var ws = &websoketkit.WebsocketHandler{}

//GetWS GetWS
func GetWS() *websoketkit.WebsocketHandler {
	return ws
}

func init() {
	ws.Init() //初始化Websockethandle
}
