package routers

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comi/config"
	"github.com/yumenaka/comi/logger"
	"github.com/yumenaka/comi/routers/handlers"
	"github.com/yumenaka/comi/routers/token"
	"github.com/yumenaka/comi/routers/websocket"
	"github.com/yumenaka/comi/types"
)

var protectedAPI *gin.RouterGroup

// 前端需要的 API
func setWebAPI(engine *gin.Engine) {
	// 路由组,方便管理部分相同的URL
	api := engine.Group("/api")

	// 无需认证，不受保护的路由
	publicRoutes := func(rg *gin.RouterGroup) {
		rg.GET("/qrcode.png", handlers.GetQrcode)
		rg.GET("/get_server_info_public", handlers.GetServerInfoPublic)
		websocket.WsDebug = &config.Config.Debug
		rg.GET("/ws", websocket.WsHandler)
	}
	publicRoutes(api)

	// 可能需要认证的路由
	protectedAPI = api.Group("/")
	// 初始化 jwtMiddleware 一次，无论是否设置了密码。
	var jwtMiddleware *jwt.GinJWTMiddleware
	if config.Config.Password != "" {
		var err error
		jwtMiddleware, err = token.NewJwtMiddleware()
		if err != nil {
			log.Fatalf("JWT Error: %s", err.Error()) // 终止程序或其他错误处理
		}
	}
	if jwtMiddleware != nil {
		// 登录、注销和 token 刷新路由只有在设置了密码时才添加
		api.POST("/login", jwtMiddleware.LoginHandler)
		api.POST("/logout", jwtMiddleware.LogoutHandler)
		api.GET("/refresh_token", jwtMiddleware.RefreshHandler)
		// 如果设置了密码，则应用 JWT 中间件到一个新的路由组
		protectedAPI = api.Group("/", jwtMiddleware.MiddlewareFunc())
	}

	//文件上传
	protectedAPI.POST("/upload", handlers.Upload)
	//通过URL字符串参数获取特定文件
	protectedAPI.GET("/get_file", handlers.GetFile)
	//登录后才能查看的服务器状态，包括标题、机器状态等
	protectedAPI.GET("/get_server_info", handlers.GetServerInfo)
	//获取书架信息，不包含每页信息
	protectedAPI.GET("/get_book_infos_by_depth", handlers.GetBookInfosByDepth)
	protectedAPI.GET("/get_book_infos_by_max_depth", handlers.GetBookInfosByMaxDepth)
	protectedAPI.GET("/get_book_infos_by_group_id", handlers.GetBookInfosByGroupID)
	//通过URL字符串参数查询书籍信息
	protectedAPI.GET("/get_book", handlers.GetBook)
	//返回同一文件夹的书籍ID列表
	protectedAPI.GET("/same_group_book_infos", handlers.SameGroupBookInfo)
	//通过链接下载reg配置
	protectedAPI.GET("/comigo.reg", handlers.GetRegFile)
	//通过链接下载toml格式的示例配置
	protectedAPI.GET("/config.toml", handlers.GetConfigToml)

	//config,改写成 RESTful 风格的 API
	//Create	POST/PUT
	//Read	    GET
	//Update	PUT
	//Delete	DELETE
	protectedAPI.GET("/config", handlers.GetConfig)
	protectedAPI.PUT("/config", handlers.UpdateConfig)
	protectedAPI.POST("/config", handlers.SaveConfig)
	//protectedAPI.DELETE("/config/:location", handlers.DeleteConfigByLocation)

	//压缩包直接下载链接
	SetDownloadLink()
}

// SetDownloadLink 压缩包直接下载链接
func SetDownloadLink() {
	if types.GetBooksNumber() >= 1 {
		allBook, err := types.GetAllBookInfoList("name")
		if err != nil {
			logger.Info("设置文件下载失败")
		} else {
			for _, info := range allBook.BookInfos {
				//下载文件
				if info.Type != types.TypeBooksGroup && info.Type != types.TypeDir {
					//staticUrl := "/raw/" + info.BookID + "/" + url.QueryEscape(info.title)
					staticUrl := "/raw/" + info.BookID + "/" + info.Title
					if checkUrlRegistered(info.BookID) {
						if config.Config.Debug {
							logger.Info("路径已注册：", info)
						}
						continue
					} else {
						protectedAPI.StaticFile(staticUrl, info.FilePath)
						staticUrlMap[info.BookID] = staticUrl
					}
				}
			}
		}
	}
}
