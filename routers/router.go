package routers

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanity-io/litter"
	"github.com/yumenaka/comi/arch"
	"github.com/yumenaka/comi/common"
	"github.com/yumenaka/comi/locale"
	"github.com/yumenaka/comi/tools"
	"html/template"
	"io/fs"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed static/index.html
var TemplateString string

//go:embed  static
var staticFS embed.FS

//go:embed  static/assets
var staticAssetFS embed.FS

//go:embed  static/images
var staticImageFS embed.FS

//退出时清理
func init() {
	common.SetupCloseHander()
}

//1、设置静态文件
func setStaticFiles(engine *gin.Engine) {
	//获取模板，命名为"template-data"，同时把左右分隔符改为 [[ ]]
	tmpl := template.Must(template.New("template-data").Delims("[[", "]]").Parse(TemplateString))
	//使用模板
	engine.SetHTMLTemplate(tmpl)
	if common.Config.LogToFile {
		// 关闭 log 打印的字体颜色。输出到文件不需要颜色
		//gin.DisableConsoleColor()
		// 中间件，输出 log 到文件
		engine.Use(tools.LoggerToFile(common.Config.LogFilePath, common.Config.LogFileName))
		//禁止控制台输出
		gin.DefaultWriter = ioutil.Discard
	}
	//自定义分隔符，避免与vue.js冲突
	engine.Delims("[[", "]]")
	//https://stackoverflow.com/questions/66248258/serve-embedded-filesystem-from-root-path-of-url
	assetsEmbedFS, err := fs.Sub(staticAssetFS, "static/assets")
	if err != nil {
		fmt.Println(err)
	}
	engine.StaticFS("/assets/", http.FS(assetsEmbedFS))
	imagesEmbedFS, errStaticImageFS := fs.Sub(staticImageFS, "static/images")
	if errStaticImageFS != nil {
		fmt.Println(errStaticImageFS)
	}
	engine.StaticFS("/images/", http.FS(imagesEmbedFS))
	//单独一张静态图片
	//singleStaticFiles(engine, "/favicon.ico", "static/images/favicon.ico", "image/x-icon")
	engine.GET("/favicon.ico", func(c *gin.Context) {
		file, _ := staticFS.ReadFile("static/images/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})
	//解析模板到HTML
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "template-data", gin.H{
			"title": common.ReadingBook.Name, //页面标题
		})
	})
	if !common.ReadingBook.IsDir {
		engine.StaticFile("/raw/"+common.ReadingBook.Name, common.ReadingBook.GetFilePath())
	}
}

//2、设置获取书籍信息、图片文件的 API
func setWebAPI(engine *gin.Engine) {
	enableAuth := common.Config.UserName != "" && common.Config.Password != ""
	var api *gin.RouterGroup
	if enableAuth {
		//简单http认证的路由组
		// 路由组：https://learnku.com/docs/gin-gonic/1.7/examples-grouping-routes/11399
		//使用 BasicAuth 中间件  https://learnku.com/docs/gin-gonic/1.7/examples-using-basicauth-middleware/11377
		api = engine.Group("/api", gin.BasicAuth(gin.Accounts{
			common.Config.UserName: common.Config.Password,
		}))
	} else {
		// 简单的路由组: api
		api = engine.Group("/api")
	}
	//解析json
	api.GET("/book.json", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, common.ReadingBook)
	})
	//解析书架json
	api.GET("/bookshelf.json", func(c *gin.Context) {
		bookShelf, err := common.GetBookShelf()
		if err != nil {
			fmt.Println(err)
		}
		c.PureJSON(http.StatusOK, bookShelf)
	})
	//通过字符串参数 查询书籍
	// 示例 URL： /get?uuid=2b15a130-06c1-4462-a3fe-5276b566d9db
	// 示例 URL： /get?&author=Doe&name=book_name
	api.GET("/get", func(c *gin.Context) {
		author := c.DefaultQuery("author", "")
		if author != "" {
			bookList, err := common.GetBookByAuthor(author)
			if err != nil {
				fmt.Println(err)
			} else {
				c.PureJSON(http.StatusOK, bookList)
			}
			return
		}
		uuid := c.DefaultQuery("uuid", "")
		if uuid != "" {
			b, err := common.GetBookByUUID(uuid)
			if err != nil {
				fmt.Println(err)
			} else {
				c.PureJSON(http.StatusOK, b)
			}
			return
		}
	})
	//服务器设定
	api.GET("/setting.json", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, common.Config)
	})
	//服务器设定
	api.GET("/config.yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, common.Config)
	})
	//初始化websocket
	api.GET("/ws", wsHandler)
	////相关参数：
	// uuid：书籍的UUID，必须项目       							&uuid=2b17a130-1
	// filename:获取的文件名，必须项目   							&filename=01.jpg
	////可选参数：
	// resize_width:指定宽度，缩放图片  							&resize_width=300
	// resize_height:指定高度，缩放图片 							&resize_height=300
	// resize_max_width:指定宽度上限，图片宽度大于这个上限时缩小图片  	&resize_max_width=300
	// resize_max_height:指定高度上限，图片高度大于这个上限时缩小图片  	&resize_max_height=300
	// auto_crop:自动切白边，数字是切白边的阈值，范围是0~100 			&auto_crop=10
	// gray:黑白化												&gray=true
	// blurhash:获取对应图片的blurhash，而不是原始图片 				&blurhash=3
	// blurhash_image:获取对应图片的blurhash图片，而不是原始图片  	&blurhash_image=3
	
	// 示例 URL： 127.0.0.1:1234/getfile?uuid=2b17a130-06c1-4222-a3fe-55ddb5ccd9db&filename=1.jpg
	//缩放文件，会转化为jpeg：http://127.0.0.1:1234/api/getfile?resize_width=300&resize_height=400&uuid=597e06&filename=01.jpeg
	api.GET("/getfile", func(c *gin.Context) {
		uuid := c.DefaultQuery("uuid", "")
		needFile := c.DefaultQuery("filename", "")
		if uuid != "" && needFile != "" {
			book, err := common.GetBookByUUID(uuid)
			if err != nil {
				fmt.Println(err)
			}
			bookPath := book.GetFilePath()
			//fmt.Println(bookPath)
			var imgData []byte
			//如果是特殊编码的ZIP文件
			if book.NonUTF8Zip && !book.IsDir {
				imgData, err = arch.GetSingleFile(bookPath, needFile, "gbk")
				if err != nil {
					fmt.Println(err)
				}
			}
			//如果是一般压缩文件
			if !book.NonUTF8Zip && !book.IsDir {
				imgData, err = arch.GetSingleFile(bookPath, needFile, "")
				if err != nil {
					fmt.Println(err)
				}
			}
			//如果是本地文件夹
			if book.IsDir {
				//直接读取磁盘文件
				imgData, err = ioutil.ReadFile(filepath.Join(bookPath, needFile))
				if err != nil {
					fmt.Println(err)
				}
			}
			//处理图像文件
			if imgData != nil {
				//默认的媒体类型，根据文件后缀判断。可能在后面有变动。
				contentType := tools.GetContentTypeByFileName(needFile)

				//读取图片Resize用的resizeWidth
				resizeWidth, errX := strconv.Atoi(c.DefaultQuery("resize_width", "0"))
				if errX != nil {
					fmt.Println(errX)
				}
				//读取图片Resize用的resizeHeight
				resizeHeight, errY := strconv.Atoi(c.DefaultQuery("resize_height", "0"))
				if errY != nil {
					fmt.Println(errY)
				}
				//图片Resize, 按照固定的width height缩放
				if errX == nil && errY == nil && resizeWidth > 0 && resizeHeight > 0 {
					imgData = tools.ImageResize(imgData, resizeWidth, resizeHeight)
					contentType = tools.GetContentTypeByFileName(".jpg")
				}

				//图片Resize, 按照 width 等比例缩放
				if errX == nil && errY != nil && resizeWidth > 0 {
					imgData = tools.ImageResizeByWidth(imgData, resizeWidth)
					contentType = tools.GetContentTypeByFileName(".jpg")
				}
				//图片Resize, 按照 height 等比例缩放
				if errY == nil && errX != nil && resizeHeight > 0 {
					imgData = tools.ImageResizeByHeight(imgData, resizeHeight)
					contentType = tools.GetContentTypeByFileName(".jpg")
				}
				//图片Resize, 按照 maxWidth 限制大小
				resizeMaxWidth, errMX := strconv.Atoi(c.DefaultQuery("resize_max_width", "0"))
				if errMX != nil {
					fmt.Println(errMX)
				}
				if errMX == nil && resizeMaxWidth > 0 {
					tempData, limitErr := tools.ImageResizeByMaxWidth(imgData, resizeMaxWidth)
					if limitErr != nil {
						fmt.Println(limitErr)
					} else {
						imgData = tempData
						contentType = tools.GetContentTypeByFileName(".jpg")
					}
				}
				//图片Resize, 按照 MaxHeight 限制大小
				resizeMaxHeight, errMY := strconv.Atoi(c.DefaultQuery("resize_max_height", "0"))
				if errMY != nil {
					fmt.Println(errMY)
				}
				if errMY == nil && resizeMaxHeight > 0 {
					tempData, limitErr := tools.ImageResizeByMaxHeight(imgData, resizeMaxHeight)
					if limitErr != nil {
						fmt.Println(limitErr)
					} else {
						imgData = tempData
						contentType = tools.GetContentTypeByFileName(".jpg")
					}
				}

				//自动切白边
				autoCrop, errCrop := strconv.Atoi(c.DefaultQuery("auto_crop", "-1"))
				if errCrop != nil {
					fmt.Println(errCrop)
				}
				if errCrop == nil && autoCrop > 0 && autoCrop <= 100 {
					imgData = tools.ImageAutoCrop(imgData, float32(autoCrop))
					contentType = tools.GetContentTypeByFileName(".jpg")
				}

				//转换为黑白图片
				gray := c.DefaultQuery("gray", "false")
				if gray == "true" {
					imgData = tools.ImageGray(imgData)
					contentType = tools.GetContentTypeByFileName(".jpg")
				}

				//获取对应图片的blurhash字符串并返回，不是图片
				blurhash, blurErr := strconv.Atoi(c.DefaultQuery("blurhash", "0"))
				if blurErr != nil {
					fmt.Println(blurErr)
				}
				//虽然blurhash components 理论上最大可以设到9，但反应速度太慢，毫无实用性、所以最大为2
				if blurhash >= 1 && blurhash <= 3 && blurErr == nil {
					hash := tools.GetImageDataBlurHash(imgData, blurhash)
					contentType = tools.GetContentTypeByFileName(".txt")
					imgData = []byte(hash)
				}

				//返回图片的blurhash图
				blurhashImage, blurImageErr := strconv.Atoi(c.DefaultQuery("blurhash_image", "0"))
				if blurImageErr != nil {
					fmt.Println(blurImageErr)
				}
				//虽然blurhash components 理论上最大可以设到9，但反应速度太慢，毫无实用性、所以最大为3
				if blurhashImage >= 1 && blurhashImage <= 3 && blurErr == nil {
					imgData = tools.GetImageDataBlurHashImage(imgData, blurhashImage)
					contentType = tools.GetContentTypeByFileName(".jpg")
				}
				c.Data(http.StatusOK, contentType, imgData)
			}
		}
	})

	//// getFileApi正常运作的话，就不需要这个虚拟文件系统
	////使用虚拟文件系统，设置服务路径（每本书都设置一遍）
	////参考了: https://bitfieldconsulting.com/golang/filesystems
	//for _, book := range common.BookList {
	//	if book.NonUTF8Zip {
	//		continue
	//	}
	//	ext := path.Ext(book.GetFilePath())
	//	if (ext == ".zip" || ext == ".epub" || ext == ".cbz") && !book.NonUTF8Zip {
	//		fsys, zipErr := zip.OpenReader(book.GetFilePath())
	//		if zipErr != nil {
	//			fmt.Println(zipErr)
	//		}
	//		httpFS := http.FS(fsys)
	//		if book.IsDir {
	//			engine.Static("/cache/"+book.BookID, book.GetFilePath())
	//		} else {
	//			engine.StaticFS("/cache/"+book.BookID, httpFS)
	//		}
	//	} else {
	//		// 通过archiver/v4，建立虚拟FS。非UTF zip文件有编码问题
	//		fsys, err := archiver.FileSystem(book.GetFilePath())
	//		httpFS := http.FS(fsys)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		if book.IsDir {
	//			engine.Static("/cache/"+book.BookID, book.GetFilePath())
	//		} else {
	//			engine.StaticFS("/cache/"+book.BookID, httpFS)
	//		}
	//	}
	//}
}

//3、选择服务端口
func setPort() {
	//检测端口
	if !tools.CheckPort(common.Config.Port) {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		if common.Config.Port+2000 > 65535 {
			common.Config.Port = common.Config.Port + r.Intn(1024)
		} else {
			common.Config.Port = 30000 + r.Intn(20000)
		}
		fmt.Println(locale.GetString("port_busy") + strconv.Itoa(common.Config.Port))
	}
}

////4、setWebpServer TODO：新的webp模式
//func setWebpServer(engine *gin.Engine) {
//	//webp反向代理
//	if common.Config.EnableWebpServer {
//		webpError := common.StartWebPServer(common.CacheFilePath+"/webp_config.json", common.ReadingBook.ExtractPath, common.CacheFilePath+"/webp", common.Config.Port+1)
//		if webpError != nil {
//			fmt.Println(locale.GetString("webp_server_error"), webpError.Error())
//			//engine.Static("/cache", common.CacheFilePath)
//
//		} else {
//			fmt.Println(locale.GetString("webp_server_start"))
//			engine.Use(reverse_proxy.ReverseProxyHandle("/cache", reverse_proxy.ReverseProxyOptions{
//				TargetHost:  "http://localhost",
//				TargetPort:  strconv.Itoa(common.Config.Port + 1),
//				RewritePath: "/cache",
//			}))
//		}
//	} else {
//		if common.ReadingBook.IsDir {
//			engine.Static("/cache/"+common.ReadingBook.BookID, common.ReadingBook.GetFilePath())
//		} else {
//			engine.Static("/cache", common.CacheFilePath)
//		}
//	}
//}

//5、setFrpClient
func setFrpClient() {
	//frp服务
	if common.Config.EnableFrpcServer {
		if common.Config.FrpConfig.RandomRemotePort {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			common.Config.FrpConfig.RemotePort = 50000 + r.Intn(10000)
		} else {
			if common.Config.FrpConfig.RemotePort <= 0 || common.Config.FrpConfig.RemotePort > 65535 {
				common.Config.FrpConfig.RemotePort = common.Config.Port
			}
		}
		frpcError := common.StartFrpC(common.CacheFilePath)
		if frpcError != nil {
			fmt.Println(locale.GetString("frpc_server_error"), frpcError.Error())
		} else {
			fmt.Println(locale.GetString("frpc_server_start"))
		}
	}
}

//6、printCMDMessage
func printCMDMessage() {
	//cmd打印链接二维码
	enableTls := common.Config.CertFile != "" && common.Config.KeyFile != ""
	tools.PrintAllReaderURL(common.Config.Port, common.Config.OpenBrowser, common.Config.EnableFrpcServer, common.Config.PrintAllIP, common.Config.Host, common.Config.FrpConfig.ServerAddr, common.Config.FrpConfig.RemotePort, common.Config.DisableLAN, enableTls)
	//打印配置，调试用
	if common.Config.Debug {
		litter.Dump(common.Config)
	}
	fmt.Println(locale.GetString("ctrl_c_hint"))
}

// StartWebServer 启动web服务
func StartWebServer() {
	//设置 gin
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	//1、setStaticFiles
	setStaticFiles(engine)
	//2、setWebAPI
	setWebAPI(engine)
	//TODO：设定第一本书
	if len(common.BookList) >= 1 {
		common.ReadingBook = common.BookList[0]
	}
	//3、setPort
	setPort()
	//4、setWebpServer
	//setWebpServer(engine)
	//5、setFrpClient
	setFrpClient()
	//6、printCMDMessage
	printCMDMessage()
	//7、StartWebServer 监听并启动web服务
	//是否对外服务
	webHost := ":"
	if common.Config.DisableLAN {
		webHost = "localhost:"
	}
	enableTls := common.Config.CertFile != "" && common.Config.KeyFile != ""
	if enableTls {
		err := engine.RunTLS(webHost+strconv.Itoa(common.Config.Port), common.Config.CertFile, common.Config.KeyFile)
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, locale.GetString("web_server_error")+"%q\n", common.Config.Port)
			if err != nil {
				return
			}
		}
	} else {
		// 监听并启动服务
		err := engine.Run(webHost + strconv.Itoa(common.Config.Port))
		if err != nil {
			_, err := fmt.Fprintf(os.Stderr, locale.GetString("web_server_error")+"%q\n", common.Config.Port)
			if err != nil {
				return
			}
		}
	}
}

////单独设定某个文件
//func singleStaticFiles(engine *gin.Engine, fileUrl string, filePath string, contentType string) {
//	engine.GET(fileUrl, func(c *gin.Context) {
//		file, _ := staticFS.ReadFile(filePath)
//		c.Data(
//			http.StatusOK,
//			contentType,
//			file,
//		)
//	})
//}
