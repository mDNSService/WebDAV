package main

import (
	"fmt"
	"github.com/mDNSService/utils/nettool"
	"golang.org/x/net/webdav"
	"net/http"
)

func main() {
	var txtInfo = nettool.MDNSServiceBaseInfo
	txtInfo["name"] = "WebDAV文件服务器"
	txtInfo["model"] = "com.iotserv.services.webdav"
	port, err := nettool.GetOneFreeTcpPort()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	server, err := nettool.RegistermDNSService(txtInfo, port)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("正在开启%s服务,端口：%d", txtInfo["name"], port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
	})
	if err != nil {
		fmt.Println(err.Error())
		server.Shutdown()
		return
	}
}
