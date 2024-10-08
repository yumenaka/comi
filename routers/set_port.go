package routers

import (
	"math/rand"
	"time"

	"github.com/yumenaka/comigo/config"
	"github.com/yumenaka/comigo/util"
	"github.com/yumenaka/comigo/util/locale"
	"github.com/yumenaka/comigo/util/logger"
)

// SetWebServerPort 3、设置服务端口
func SetWebServerPort() {
	//检测端口
	if !util.CheckPort(config.Config.Port) {
		//获取一个空闲可用的系统端口号
		port, err := util.GetFreePort()
		if err != nil {
			logger.Infof("%s", err)
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			if config.Config.Port+2000 > 65535 {
				config.Config.Port = config.Config.Port + r.Intn(1024)
			} else {
				config.Config.Port = 30000 + r.Intn(20000)
			}
		} else {
			config.Config.Port = port
		}
		logger.Infof(locale.GetString("port_busy")+"%s", config.Config.Port)
	}
}
