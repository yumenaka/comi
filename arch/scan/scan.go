package scan

import (
	"context"
	"errors"
	"fmt"
	"github.com/klauspost/compress/zip"
	"github.com/sirupsen/logrus"
	"github.com/yumenaka/archiver/v4"
	"github.com/yumenaka/comi/arch"
	"github.com/yumenaka/comi/locale"
	"github.com/yumenaka/comi/storage"
	"github.com/yumenaka/comi/types"
	"github.com/yumenaka/comi/util"
	"io/fs"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type ScanConfig struct {
	ReScanFile            bool     // 是否重新扫描文件
	StoresPath            []string // 书库路径
	MaxScanDepth          int      // 扫描深度
	MinImageNum           int      // 最小图片数量
	TimeoutLimitForScan   int      // 扫描超时时间
	ExcludePath           []string // 排除路径
	SupportMediaType      []string // 支持的媒体类型
	SupportFileType       []string // 支持的文件类型
	ZipFileTextEncoding   string   // 非UTF-8编码的ZIP文件，尝试用什么编码解析，默认GBK
	EnableDatabase        bool     // 启用数据库
	ClearDatabaseWhenExit bool     // 启用数据库时，扫描完成后，清除不存在的书籍
	Debug                 bool
}

func NewScanConfig(
	reScanFile bool,
	storesPath []string,
	maxScanDepth int,
	minImageNum int,
	timeoutLimitForScan int,
	excludePath []string,
	supportMediaType []string,
	supportFileType []string,
	zipFileTextEncoding string,
	enableDatabase bool,
	clearDatabaseWhenExit bool,
	debug bool,
) ScanConfig {
	return ScanConfig{
		ReScanFile:            reScanFile,
		StoresPath:            storesPath,
		MaxScanDepth:          maxScanDepth,
		MinImageNum:           minImageNum,
		TimeoutLimitForScan:   timeoutLimitForScan,
		ExcludePath:           excludePath,
		SupportMediaType:      supportMediaType,
		SupportFileType:       supportFileType,
		ZipFileTextEncoding:   zipFileTextEncoding,
		EnableDatabase:        enableDatabase,
		ClearDatabaseWhenExit: clearDatabaseWhenExit,
		Debug:                 debug,
	}
}

// IsSupportMedia 判断压缩包内的文件是否需要展示（包括图片、音频、视频、PDF在内的媒体文件）
func (c *ScanConfig) IsSupportMedia(checkPath string) bool {
	for _, ex := range c.SupportMediaType {
		suffix := strings.ToLower(path.Ext(checkPath)) //strings.ToLower():某些文件会用大写文件名
		if ex == suffix {
			return true
		}
	}
	return false
}

// IsSupportArchiver 是否是支持的压缩文件
func (c *ScanConfig) IsSupportArchiver(checkPath string) bool {
	for _, ex := range c.SupportFileType {
		suffix := path.Ext(checkPath)
		if ex == suffix {
			return true
		}
	}
	return false
}

// IsSkipDir  检查路径是否应该跳过（排除文件，文件夹列表）。
func (c *ScanConfig) IsSkipDir(path string) bool {
	for _, substr := range c.ExcludePath {
		if strings.HasSuffix(path, substr) {
			return true
		}
	}
	return false
}

// ScanStorePath 3、扫描配置文件指定的的书籍库
func ScanStorePath(scanConfig ScanConfig) error {
	if len(scanConfig.StoresPath) > 0 {
		for _, p := range scanConfig.StoresPath {
			addList, err := ScanAndGetBookList(p, scanConfig)
			if err != nil {
				fmt.Println(locale.GetString("scan_error"), p, err)
				return err
			} else {
				AddBooksToStore(addList, p, scanConfig.MinImageNum)
			}
		}
	}
	return nil
}

// SaveResultsToDatabase 4，保存扫描结果到数据库，并清理不存在的书籍
func SaveResultsToDatabase(ConfigPath string, ClearDatabaseWhenExit bool, Debug bool) error {
	err := storage.InitDatabase(ConfigPath)
	if err != nil {
		return err
	}
	AllBook := types.GetAllBookList()
	//设置清理数据库的时候，是否清理没扫描到的书籍信息
	if ClearDatabaseWhenExit {
		for _, b := range AllBook {
			storage.ClearBookData(b, Debug)
		}
	}
	saveErr := storage.SaveBookListToDatabase(AllBook)
	if saveErr != nil {
		fmt.Println(saveErr)
		return saveErr
	}
	return nil
}

// AddBooksToStore 添加一组书到书库
func AddBooksToStore(bookList []*types.Book, basePath string, MinImageNum int) {
	err := types.AddBooks(bookList, basePath, MinImageNum)
	if err != nil {
		fmt.Println(locale.GetString("AddBook_error"), basePath)
	}
	// 然后生成对应的虚拟书籍组
	if err := types.Stores.GenerateBookGroup(); err != nil {
		fmt.Println(err)
	}
}

// ScanAndGetBookList 扫描路径，取得路径里的书籍
func ScanAndGetBookList(storePath string, scanConfig ScanConfig) (newBookList []*types.Book, err error) {
	// 路径不存在的Error，不过目前并不会打印出来
	if !util.PathExists(storePath) {
		return nil, errors.New(locale.GetString("PATH_NOT_EXIST"))
	}
	storePathAbs, err := filepath.Abs(storePath)
	if err != nil {
		storePathAbs = storePath
		fmt.Println(err)
	}
	fmt.Println(locale.GetString("SCAN_START_HINT") + storePathAbs)
	err = filepath.Walk(storePathAbs, func(walkPath string, fileInfo os.FileInfo, err error) error {
		// 是否需要跳过
		skip := false
		if !scanConfig.ReScanFile {
			for _, p := range types.GetAllBookList() {
				AbsW, err := filepath.Abs(walkPath) // 取得绝对路径
				if err != nil {
					// 无法取得的情况下，用相对路径
					AbsW = walkPath
					fmt.Println(err, AbsW)
				}
				if walkPath == p.FilePath || AbsW == p.FilePath {
					skip = true
					newBookList = append(newBookList, p)
				}
			}
		} else {
			skip = false
		}
		if skip {
			fmt.Println(locale.GetString("FoundInDatabase") + walkPath)
			return nil
		}
		// 路径深度
		depth := strings.Count(walkPath, "/") - strings.Count(storePathAbs, "/")
		if runtime.GOOS == "windows" {
			depth = strings.Count(walkPath, "\\") - strings.Count(storePathAbs, "\\")
		}
		if depth > scanConfig.MaxScanDepth {
			fmt.Printf(locale.GetString("ExceedsMaximumDepth")+" %d，base：%s scan: %s:\n", scanConfig.MaxScanDepth, storePathAbs, walkPath)
			return filepath.SkipDir // 当WalkFunc的返回值是filepath.SkipDir时，Walk将会跳过这个目录，照常执行下一个文件。
		}
		if scanConfig.IsSkipDir(walkPath) {
			fmt.Println(locale.GetString("SkipPath") + walkPath)
			return filepath.SkipDir
		}
		if fileInfo == nil {
			return err
		}
		// 如果不是文件夹
		if !fileInfo.IsDir() {
			if !scanConfig.IsSupportArchiver(walkPath) {
				return nil
			}
			// 得到书籍文件数据
			getBook, err := scanFileGetBook(walkPath, storePathAbs, depth, scanConfig)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			newBookList = append(newBookList, getBook)
		}
		// 如果是文件夹
		if fileInfo.IsDir() {
			// 得到书籍文件数据
			getBook, err := scanDirGetBook(walkPath, storePathAbs, depth, scanConfig)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			newBookList = append(newBookList, getBook)
		}
		return nil
	})
	// 所有可用书籍，包括压缩包与文件夹
	if len(newBookList) > 0 {
		fmt.Printf(locale.GetString("FOUND_IN_PATH"), len(newBookList), storePathAbs)
	}
	return newBookList, err
}

func scanDirGetBook(dirPath string, storePath string, depth int, scanConfig ScanConfig) (*types.Book, error) {
	// 初始化，生成UUID
	newBook, err := types.New(dirPath, time.Now(), 0, storePath, depth, types.TypeDir)
	if err != nil {
		return nil, err
	}
	//// 获取目录中的文件和子目录的详细信息
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	for _, file := range infos {
		// 跳过子目录, 只搜寻目录中的文件
		if file.IsDir() {
			continue
		}
		// 输出绝对路径
		strAbsPath, errPath := filepath.Abs(dirPath + "/" + file.Name())
		if errPath != nil {
			fmt.Println(errPath)
		}
		if scanConfig.IsSupportMedia(file.Name()) {
			TempURL := "api/getfile?id=" + newBook.BookID + "&filename=" + url.QueryEscape(file.Name())
			newBook.Pages.Images = append(newBook.Pages.Images, types.ImageInfo{RealImageFilePATH: strAbsPath, FileSize: file.Size(), ModeTime: file.ModTime(), NameInArchive: file.Name(), Url: TempURL})
		}
	}
	newBook.SortPages("default")
	// 在添加到书库时判断页数
	return newBook, err
}

// 扫描一个路径，并返回对应书籍
func scanFileGetBook(filePath string, storePath string, depth int, scanConfig ScanConfig) (*types.Book, error) {
	// 打开文件
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0o400) // Use mode 0400 for a read-only // file and 0600 for a readable+writable file.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(file)
	FileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}
	// 初始化一本书，设置文件路径等等
	newBook, err := types.New(filePath, FileInfo.ModTime(), FileInfo.Size(), storePath, depth, types.GetBookTypeByFilename(filePath))
	if err != nil {
		return nil, err
	}
	// 根据文件类型，走不同的初始化流程
	switch newBook.Type {
	// 为解决archiver/v4的BUG “zip文件无法读取2级目录” 单独处理zip文件
	case types.TypeZip, types.TypeCbz, types.TypeEpub:
		// 使用Archiver的虚拟文件系统，无法处理非UTF-8编码
		fsys, zipErr := zip.OpenReader(filePath)
		if zipErr != nil {
			// fmt.Println(zipErr)
			return nil, errors.New(locale.GetString("NOT_A_VALID_ZIP_FILE") + filePath)
		}
		err = walkUTF8ZipFs(fsys, "", ".", newBook, scanConfig)
		// 如果扫描ZIP文件的时候遇到了 fs.PathError ，则扫描到NonUTF-8 ZIP文件，需要特殊处理
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			if scanConfig.Debug {
				fmt.Println("NonUTF-8 ZIP:" + filePath + "  Error:" + err.Error())
			}
			// 忽略 fs.PathError 并换个方式扫描
			err = scanNonUTF8ZipFile(filePath, newBook, scanConfig)
		}
		// epub文件，需要根据 META-INF/container.xml 里面定义的rootfile （.opf文件）来重新排序
		if newBook.Type == types.TypeEpub {
			imageList, err := arch.GetImageListFromEpubFile(newBook.FilePath)
			if err != nil {
				fmt.Println(err)
			} else {
				newBook.SortPagesByImageList(imageList)
			}
			// 根据metadata，改写书籍信息
			metaData, err := arch.GetEpubMetadata(newBook.FilePath)
			if err != nil {
				fmt.Println(err)
			} else {
				newBook.Author = metaData.Creator
				newBook.Press = metaData.Publisher
			}
		}
	// TODO:服务器解压速度太慢，网页用PDF.js解析？
	case types.TypePDF:
		newBook.AllPageNum = 1
		newBook.InitComplete = true
		newBook.Cover = types.ImageInfo{RealImageFilePATH: "", FileSize: FileInfo.Size(), ModeTime: FileInfo.ModTime(), NameInArchive: "", Url: "/images/pdf.png"}
	// TODO：简单的网页播放器
	case types.TypeVideo:
		newBook.AllPageNum = 1
		newBook.InitComplete = true
		newBook.Cover = types.ImageInfo{NameInArchive: "video.png", Url: "/images/video.png"}
	case types.TypeAudio:
		newBook.AllPageNum = 1
		newBook.InitComplete = true
		newBook.Cover = types.ImageInfo{NameInArchive: "audio.png", Url: "/images/audio.png"}
	case types.TypeUnknownFile:
		newBook.AllPageNum = 1
		newBook.InitComplete = true
		newBook.Cover = types.ImageInfo{NameInArchive: "unknown.png", Url: "/images/unknown.png"}
	// 其他类型的压缩文件或文件夹
	default:
		// archiver.FileSystem可以配合ctx了，加个默认超时时间
		const shortDuration = 10 * 1000 * time.Millisecond // 超时时间，10秒
		ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
		defer cancel()
		fsys, err := archiver.FileSystem(ctx, filePath)
		if err != nil {
			return nil, err
		}
		err = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
			if scanConfig.IsSkipDir(path) {
				fmt.Println("Skip Scan:" + path)
				return fs.SkipDir
			}
			f, errInfo := d.Info()
			if errInfo != nil {
				fmt.Println(errInfo)
				return fs.SkipDir
			}
			if !scanConfig.IsSupportMedia(path) {
				logrus.Debugf(locale.GetString("unsupported_file_type") + path)
			} else {
				u, ok := f.(archiver.File) // f.Name不包含路径信息.需要转换一下
				if !ok {
					// 如果是文件夹+图片
					newBook.Type = types.TypeDir
					////用Archiver的虚拟文件系统提供图片文件，理论上现在不应该用到
					//newBook.Pages = append(newBook.Pages, ImageInfo{RealImageFilePATH: "", FileSize: f.Size(), ModeTime: f.ModTime(), NameInArchive: "", Url: "/cache/" + newBook.BookID + "/" + url.QueryEscape(path)})
					//实验：用getfile接口提供文件服务
					TempURL := "api/getfile?id=" + newBook.BookID + "&filename=" + url.QueryEscape(path)
					newBook.Pages.Images = append(newBook.Pages.Images, types.ImageInfo{RealImageFilePATH: "", FileSize: f.Size(), ModeTime: f.ModTime(), NameInArchive: "", Url: TempURL})
					// fmt.Println(locale.GetString("unsupported_extract")+" %s", f)
				} else {
					// 替换特殊字符的时候，额外将“+替换成"%2b"，因为gin会将+解析为空格。
					TempURL := "api/getfile?id=" + newBook.BookID + "&filename=" + url.QueryEscape(u.NameInArchive)
					// 不替换特殊字符
					// TempURL := "api/getfile?id=" + newBook.BookID + "&filename=" + u.NameInArchive
					newBook.Pages.Images = append(newBook.Pages.Images, types.ImageInfo{RealImageFilePATH: "", FileSize: f.Size(), ModeTime: f.ModTime(), NameInArchive: u.NameInArchive, Url: TempURL})
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	// 不管页数，直接返回：在添加到书库时判断页数
	newBook.SortPages("default")
	return newBook, err
}

func scanNonUTF8ZipFile(filePath string, b *types.Book, scanConfig ScanConfig) error {
	b.NonUTF8Zip = true
	reader, err := arch.ScanNonUTF8Zip(filePath, scanConfig.ZipFileTextEncoding)
	if err != nil {
		return err
	}
	for _, f := range reader.File {
		if scanConfig.IsSupportMedia(f.Name) {
			// 如果是压缩文件
			// 替换特殊字符的时候，额外将“+替换成"%2b"，因为gin会将+解析为空格。
			TempURL := "api/getfile?id=" + b.BookID + "&filename=" + url.QueryEscape(f.Name)
			b.Pages.Images = append(b.Pages.Images, types.ImageInfo{RealImageFilePATH: "", FileSize: f.FileInfo().Size(), ModeTime: f.FileInfo().ModTime(), NameInArchive: f.Name, Url: TempURL})
		} else {
			logrus.Debugf(locale.GetString("unsupported_file_type") + f.Name)
		}
	}
	b.SortPages("default")
	return err
}

// 手动写的递归查找，功能与fs.WalkDir()相同。发现一个Archiver/V4的BUG：zip文件的虚拟文件系统，找不到正确的多级文件夹？
// https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter06/06.3.html
func walkUTF8ZipFs(fsys fs.FS, parent, base string, b *types.Book, scanConfig ScanConfig) error {
	// 一般zip文件的处理流程
	// fmt.Println("parent:" + parent + " base:" + base)
	dirName := path.Join(parent, base)
	dirEntries, err := fs.ReadDir(fsys, dirName)
	for _, dirEntry := range dirEntries {
		name := dirEntry.Name()
		f, errInfo := dirEntry.Info()
		if errInfo != nil {
			continue
		}
		if dirEntry.IsDir() == true {
			switch name {
			case ".comigo":
				return fs.SkipDir
			case "flutter_ui":
				return fs.SkipDir
			case "node_modules":
				return fs.SkipDir
			default:
			}
			joinPath := path.Join(parent, name)
			err = walkUTF8ZipFs(fsys, joinPath, base, b, scanConfig)
		} else if !scanConfig.IsSupportMedia(name) {
			logrus.Debugf(locale.GetString("unsupported_file_type") + name)
		} else {
			inArchiveName := path.Join(parent, f.Name())
			TempURL := "api/getfile?id=" + b.BookID + "&filename=" + url.QueryEscape(inArchiveName)
			// 替换特殊字符的时候,不要用url.PathEscape()，PathEscape不会把“+“替换成"%2b"，会导致BUG，让gin会将+解析为空格。
			b.Pages.Images = append(b.Pages.Images, types.ImageInfo{RealImageFilePATH: "", FileSize: f.Size(), ModeTime: f.ModTime(), NameInArchive: inArchiveName, Url: TempURL})
		}
	}
	b.SortPages("default")
	return err
}