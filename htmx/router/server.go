package router

import (
	"errors"
	"github.com/yumenaka/comigo/htmx/comigo"
	"io/fs"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comigo/config"
	"github.com/yumenaka/comigo/util/logger"
)

func noCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}

// RunServer 运行一个新的 HTTP 服务器。
func RunServer() (err error) {

	gin.SetMode(gin.ReleaseMode)
	// 创建一个新的Gin服务器。
	router := gin.Default()

	// 使用 noCache 中间件
	router.Use(noCache())

	// 扫描漫画
	comigo.StartComigoServer(router)
	// 为模板引擎定义 HTML 渲染器。
	router.HTMLRender = &TemplRender{}
	// 静态文件。
	//router.Static("/static", "./router/static")
	// 嵌入静态文件。
	staticFS, err := fs.Sub(static, "static")
	if err != nil {
		logger.Infof("%s", err)
	}
	router.StaticFS("/static/", http.FS(staticFS))
	// 设置路由
	bindURL(router)

	// 发消息
	slog.Info("Starting server...", "port", config.Config.Port)

	// 是否对外服务
	webHost := ":"
	if config.Config.DisableLAN {
		webHost = "localhost:"
	}
	// 是否启用TLS
	enableTLS := config.Config.CertFile != "" && config.Config.KeyFile != ""
	server := &http.Server{
		Addr:         webHost + strconv.Itoa(config.Config.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router, // gin.Engine本身可以作为一个Handler传递到http包,用于启动服务器
	}
	// 监听并启动服务(TLS)
	if enableTLS {
		if err = server.ListenAndServeTLS(config.Config.CertFile, config.Config.KeyFile); err != nil && !errors.Is(err, http.ErrServerClosed) {
			time.Sleep(3 * time.Second)
			logger.Fatalf("listen: %s\n", err)
		}
	}
	if !enableTLS {
		// 监听并启动服务(HTTP)
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			time.Sleep(3 * time.Second)
			logger.Fatalf("listen: %s\n", err)
		}
	}
	return err
}
