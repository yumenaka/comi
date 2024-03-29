package cmd

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yumenaka/comi/config"
	"github.com/yumenaka/comi/locale"
	"github.com/yumenaka/comi/logger"
	"github.com/yumenaka/comi/routers"
	"github.com/yumenaka/comi/routers/handlers"
)

// cobra & viper sample:
// https://qiita.com/nirasan/items/cc2ab5bc2889401fe596
var runtimeViper *viper.Viper

func init() {
	runtimeViper = viper.New()
}

// rootCmd 没有任何子命令的情况下时的基本命令
var rootCmd = &cobra.Command{
	Use:     locale.GetString("comigo_use"),
	Short:   locale.GetString("short_description"),
	Example: locale.GetString("comigo_example"),
	Version: config.Version,
	Long:    locale.GetString("long_description"),
	// 不加参数的命令。
	Run: func(cmd *cobra.Command, args []string) {
		//解析命令，扫描文件
		StartScan(args)
		//设置临时文件夹
		config.SetTempDir()
		//SetPort
		routers.SetPort()
		//设置书籍API
		routers.StartWebServer()
		//退出时清理临时文件
		SetShutdownHandler()
	},
}

// 读取顺序：RAM（代码当中设定的默认值）+命令行参数  -> HomeDirectory -> ProgramDirectory -> WorkingDirectory
func initConfigFile() {
	home, err := homedir.Dir()
	if err != nil {
		logger.Infof("%s", err)
	}
	//在HomeDir搜索配置
	homeConfigDir := path.Join(home, ".config/comigo")
	runtimeViper.AddConfigPath(homeConfigDir)
	// 在ProgramDirectory(二进制程序所在文件夹）的配置
	ProgramDirectory, err := os.Executable()
	if err != nil {
		logger.Infof("Failed to get ProgramDirectory:", err)
		return
	}
	// 将ProgramDirectory转换为绝对路径
	absPath, err := filepath.Abs(ProgramDirectory)
	if err != nil {
		logger.Infof("Failed to get absolute path:%s", err)
		return
	}
	logger.Infof("ProgramDirectory:%s", absPath)
	runtimeViper.AddConfigPath(absPath)

	// WorkingDirectory：当前执行目录
	WorkingDirectory, err := os.Getwd()
	if err != nil {
		logger.Infof("Failed to get WorkingDirectory:%s", err)
	}
	runtimeViper.AddConfigPath(WorkingDirectory)

	runtimeViper.SetConfigType("toml")
	runtimeViper.SetConfigName("config.toml")

	//用户命令行指定的目录或文件
	if config.Config.ConfigPath != "" {
		//SetConfigFile 显式定义配置文件的路径、名称和扩展名。 Viper 将使用它并且不检查任何配置路径。
		runtimeViper.SetConfigFile(config.Config.ConfigPath)
	}

	// 读取设定文件
	if err := runtimeViper.ReadInConfig(); err != nil {
		if config.Config.ConfigPath == "" {
			logger.Infof("%s", err)
		}
	} else {
		//获取当前使用的配置文件路径
		//https://github.com/spf13/viper/issues/89
		tempConfigPath := runtimeViper.ConfigFileUsed()
		logger.Infof(locale.GetString("FoundConfigFile")+"%s", tempConfigPath)
	}
	// 把设定文件的内容，解析到构造体里面。
	if err := runtimeViper.Unmarshal(&config.Config); err != nil {
		logger.Infof("%s", err)
		os.Exit(1)
	}
	////监听文件修改
	//runtimeViper.WatchConfig()
	////文件修改时，执行重载设置、服务重启的函数
	//runtimeViper.OnConfigChange(handlerConfigReload)
}

// Execute 执行将所有子命令添加到根命令并适当设置标志。
// 这是由 main.main() 调用的。 rootCmd 只需要执行一次。
func Execute() {
	//初始化命令行参数。不能放在初始化配置文件之后。
	initFlags()
	//初始化配置文件
	cobra.OnInitialize(initConfigFile) // "OnInitialize"传入的函数，应该会在所有命令执行之前，包括rootCmd.Run之前执行。
	//执行命令
	if err := rootCmd.Execute(); err != nil {
		logger.Infof("%s", err)
		time.Sleep(3 * time.Second)
		os.Exit(1)
	}
}

func initFlags() {
	//加载环境变量，改写对应值
	runtimeViper.AutomaticEnv()
	//设置环境变量的前缀，将 PORT变为 COMI_PORT
	runtimeViper.SetEnvPrefix("COMI")
	cobra.MousetrapHelpText = ""       //屏蔽鼠标提示，支持拖拽、双击运行
	cobra.MousetrapDisplayDuration = 5 //"这是命令行程序"的提醒表示时间
	//登陆用户名、密码
	rootCmd.PersistentFlags().BoolVar(&config.Config.EnableLogin, "login", false, locale.GetString("ENABLE_LOGIN"))
	rootCmd.PersistentFlags().StringVarP(&config.Config.Username, "username", "u", "admin", locale.GetString("USERNAME"))
	rootCmd.PersistentFlags().StringVarP(&config.Config.Password, "password", "k", "", locale.GetString("PASSWORD"))
	rootCmd.PersistentFlags().IntVarP(&config.Config.Timeout, "timeout", "t", 65535, locale.GetString("TIMEOUT"))
	//TLS设定
	rootCmd.PersistentFlags().BoolVar(&config.Config.EnableTLS, "tls", false, locale.GetString("TLS_ENABLE"))
	rootCmd.PersistentFlags().StringVar(&config.Config.CertFile, "tls-crt", "", locale.GetString("TLS_CRT"))
	rootCmd.PersistentFlags().StringVar(&config.Config.KeyFile, "tls-key", "", locale.GetString("TLS_KEY"))
	//指定配置文件
	rootCmd.PersistentFlags().StringVarP(&config.Config.ConfigPath, "config", "c", "", locale.GetString("CONFIG"))
	//启用数据库，保存扫描数据
	rootCmd.PersistentFlags().BoolVarP(&config.Config.EnableDatabase, "database", "e", false, locale.GetString("EnableDatabase"))
	//服务端口
	rootCmd.PersistentFlags().IntVarP(&config.Config.Port, "port", "p", 1234, locale.GetString("PORT"))
	//本地Host
	rootCmd.PersistentFlags().StringVar(&config.Config.Host, "host", "DefaultHost", locale.GetString("LOCAL_HOST"))
	//DEBUG
	rootCmd.PersistentFlags().BoolVar(&config.Config.Debug, "debug", false, locale.GetString("DEBUG_MODE"))
	//启用文件上传功能
	rootCmd.PersistentFlags().BoolVar(&config.Config.EnableUpload, "enable-upload", true, locale.GetString("ENABLE_FILE_UPLOAD"))
	//上传文件的保存路径
	rootCmd.PersistentFlags().StringVar(&config.Config.UploadPath, "upload-path", "", locale.GetString("UPLOAD_PATH"))
	if config.Config.EnableUpload && config.Config.UploadPath == "" {
		//获取当前目录
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			config.Config.UploadPath = path.Join(dir, "upload")
		}
	}
	//打开浏览器
	rootCmd.PersistentFlags().BoolVarP(&config.Config.OpenBrowser, "open-browser", "o", false, locale.GetString("OPEN_BROWSER"))
	if runtime.GOOS == "windows" {
		config.Config.OpenBrowser = true
	}
	//不对局域网开放
	rootCmd.PersistentFlags().BoolVarP(&config.Config.DisableLAN, "disable-lan", "d", false, locale.GetString("DISABLE_LAN"))
	//文件搜索深度
	rootCmd.PersistentFlags().IntVarP(&config.Config.MaxScanDepth, "max-depth", "m", 3, locale.GetString("MAX_DEPTH"))
	////服务器解析书籍元数据，如果生成blurhash，需要消耗大量资源
	rootCmd.PersistentFlags().BoolVar(&config.Config.GenerateMetaData, "generate-metadata", false, locale.GetString("GENERATE_METADATA"))
	//打印所有可用网卡ip
	rootCmd.PersistentFlags().BoolVar(&config.Config.PrintAllPossibleQRCode, "print-all", false, locale.GetString("PRINT_ALL_IP"))
	//至少有几张图片，才认定为漫画压缩包
	rootCmd.PersistentFlags().IntVar(&config.Config.MinImageNum, "min-image", 1, locale.GetString("MIN_MEDIA_NUM"))
	//输出log文件
	rootCmd.PersistentFlags().BoolVar(&config.Config.LogToFile, "log", false, locale.GetString("LOG_TO_FILE"))
	//web图片缓存
	rootCmd.PersistentFlags().BoolVar(&config.Config.UseCache, "use-cache", false, locale.GetString("CACHE_FILE_ENABLE"))
	//图片缓存路径
	rootCmd.PersistentFlags().StringVar(&config.Config.CachePath, "cache-path", "", locale.GetString("CACHE_FILE_PATH"))
	//退出时清除缓存
	rootCmd.PersistentFlags().BoolVar(&config.Config.ClearCacheExit, "cache-clean", true, locale.GetString("CACHE_FILE_CLEAN"))
	handlers.EnableUpload = &config.Config.EnableUpload
	handlers.UploadPath = &config.Config.UploadPath
	//手动指定zip文件编码 gbk、shiftjis……
	rootCmd.PersistentFlags().StringVar(&config.Config.ZipFileTextEncoding, "zip-encode", "gbk", locale.GetString("ZIP_ENCODE"))
}
