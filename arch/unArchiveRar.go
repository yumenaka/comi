package arch

import (
	"context"
	"fmt"
	"github.com/yumenaka/archiver/v4"
	"github.com/yumenaka/comi/tools"
	"os"
)

// UnArchiveRar  一次性解压rar文件
func UnArchiveRar(filePath string, extractPath string) error {
	extractPath = tools.GetAbsPath(extractPath)
	//如果解压路径不存在，创建路径
	err := os.MkdirAll(extractPath, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	//打开文件，只读模式
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0400) //Use mode 0400 for a read-only // file and 0600 for a readable+writable file.
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//是否是压缩包
	format, _, err := archiver.Identify(filePath, file)
	if err != nil {
		return err
	}
	//如果是rar
	if ex, ok := format.(archiver.Rar); ok {
		ctx := context.Background()
		//WithValue返回parent的一个副本，该副本保存了传入的key/value，而调用Context接口的Value(key)方法就可以得到val。注意在同一个context中设置key/value，若key相同，值会被覆盖。
		ctx = context.WithValue(ctx, "extractPath", extractPath)
		err := ex.LsAllFile(ctx, file, extractFileHandler)
		if err != nil {
			return err
		}
		fmt.Println("rar文件解压完成：" + tools.GetAbsPath(filePath) + " 解压到：" + tools.GetAbsPath(extractPath))
	}
	return nil
}