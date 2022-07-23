package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/yumenaka/comi/common"
	"github.com/yumenaka/comi/routers"
)

//Go 每日一库之 fsnotify  https://darjun.github.io/2020/01/19/godailylib/fsnotify/

//优雅地重启或停止  https://learnku.com/docs/gin-gonic/1.7/examples-graceful-restart-or-stop/11376
func configReloadHandler(e fsnotify.Event) {
	//打印配置文件路径与触发事件
	fmt.Printf("配置文件改变，Comigo将在5秒后重启:%s Op:%s\n", e.Name, e.Op)
	//重新读取改变后的配置文件
	if err := viperInstance.ReadInConfig(); err != nil {
		if common.ConfigFile == "" && common.Config.Debug {
			fmt.Println(err)
		}
	}
	// 把设定文件的内容，解析到构造体里面。
	if err := viperInstance.Unmarshal(&common.Config); err != nil {
		fmt.Println(err)
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
	// 上下文用于通知服务器它有 5 秒的时间来完成它当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := common.Srv.Shutdown(ctx); err != nil {
		time.Sleep(3 * time.Second)
		log.Fatal("Server forced to shutdown: ", err)
	}
	<-ctx.Done()
	//3、扫描配置文件指定的书籍库
	ScanStorePathInConfig()
	//4，保存扫描结果到数据库
	SaveResultsToDatabase()
	//5、通过“可执行文件名”设置部分默认参数,目前不生效
	common.Config.SetByExecutableFilename()
	//重新设置文件下载链接
	routers.SetDownloadLink()
	//重启 web 服务器
	routers.StartWebServer()
}