package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yumenaka/comigo/util"
	"github.com/yumenaka/comigo/util/locale"
	"github.com/yumenaka/comigo/util/logger"
	"net/http"
	"os"
	"path/filepath"
)

var LocalRescanBroadcast *chan string
var EnableUpload *bool
var UploadPath *string

// UploadFile 下载服务器配置
// 除了设置头像以外，也可以做上传文件并阅读功能
// Set a lower memory limit for multipart forms (default is 32 MiB)

// UploadFile engine.MaxMultipartMemory = 60 << 20  // 60 MiB  只限制程序在上传文件时可以使用多少内存，而是不限制上传文件的大小。(default is 32 MiB)
func UploadFile(c *gin.Context) {
	if !*EnableUpload {
		logger.Infof("%s", locale.GetString("UPLOAD_DISABLE_HINT"))
		return
	}
	//logger.Infof("EnableUpload:", *EnableUpload)
	//默认的上传路径
	path := "upload"
	//如果设置过上传路径
	if *UploadPath != "" {
		path = *UploadPath
	}
	//如果保存路径不存在，就新建路径
	if !util.IsExist(path) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Infof("mkdir failed![%s]\n", err)
		} else {
			logger.Infof("%s", "mkdir success!\n")
		}
	}

	// 上传单个文件 sample: https://github.com/gin-gonic/examples/blob/master/upload-file/single/main.go
	file, err := c.FormFile("file")
	if err != nil { //没有传文件会报错，处理这个错误。
		logger.Infof("%s", err)
	}

	// UploadFile the file to specific dst.
	err = c.SaveUploadedFile(file, filepath.Join(path, file.Filename))
	if err != nil {
		logger.Infof("%s", err)
	}
	/*
	  也可以直接使用io操作，拷贝文件数据。
	  out, err := os.Create(filename)
	  defer out.Close()
	  _, err = io.Copy(out, file)
	*/
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	*LocalRescanBroadcast <- "upload"

	// 上传多个文件，未调通，示例：
	//https://github.com/gin-gonic/examples/blob/master/upload-file/multiple/main.go
	//form, err := c.MultipartForm()
	//if err != nil {
	//	c.String(http.StatusBadRequest, "get form err: %s", err.Error())
	//	return
	//}
	//files := form.File["files"]
	//logger.Infof("files:", files)
	//for _, file := range files {
	//	filename := filepath.Base(file.Filename)
	//	if err := c.SaveUploadedFile(file, filepath.Join(path, filename)); err != nil {
	//		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
	//		return
	//	}
	//}
	//c.String(http.StatusOK, "Uploaded successfully %d files", len(files))
	//*LocalRescanBroadcast <- "upload"
}
