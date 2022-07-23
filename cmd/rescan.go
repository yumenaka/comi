package cmd

import (
	"fmt"

	"github.com/yumenaka/comi/common"
	"github.com/yumenaka/comi/locale"
	"github.com/yumenaka/comi/routers"
	"github.com/yumenaka/comi/routers/handler"
)

//用于由客户端发送消息的队列，扮演通道的角色。后面定义了一个 goroutine 来从这个通道读取新消息，然后将它们发送给其它连接到服务器的客户端。
var rescanBroadcast = make(chan string) // broadcast channel
func init() {
	// Start listening for incoming chat messages
	go waitRescanMessages()
	handler.LocalRescanBroadcast = &rescanBroadcast
}

//一个简单循环，从“broadcast”中连续读取数据，然后通过各自的 WebSocket 连接将消息传播到客户端。
func waitRescanMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-rescanBroadcast //广播频道
		// Send it out to every client that is currently connected
		switch msg {
		case "ComigoUpload":
			fmt.Println("扫描上传文件夹：", msg)
			ReScanUploadPath()
		case "SomePath":
			fmt.Println("收到重新扫描消息：", msg)
			ReScanPath(msg)
		default:
			continue
		}
	}
}

//ReScanUploadPath 重新扫描上传目录,因为需要设置下载路径，gin 初始化后才能执行
func ReScanUploadPath() {
	//没启用上传，则不扫描
	if !common.Config.EnableUpload {
		return
	}
	uploadPath := "ComigoUpload"
	if common.Config.UploadPath != "" {
		uploadPath = common.Config.UploadPath
	}
	ReScanPath(uploadPath)
}

func ReScanPath(path string) {
	//扫描上传目录的文件
	addList, err := common.ScanAndGetBookList(path, databaseBookList)
	if err != nil {
		fmt.Println(locale.GetString("scan_error"), path)
	} else {
		common.AddBooksToStore(addList, path)
	}
	//保存扫描结果到数据库
	SaveResultsToDatabase()
	//重新设置文件下载链接
	routers.SetDownloadLink()
}