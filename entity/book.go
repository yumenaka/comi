package entity

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/xxjwxc/gowp/workpool"
	"github.com/yumenaka/comigo/util"
	"github.com/yumenaka/comigo/util/locale"
	"github.com/yumenaka/comigo/util/logger"
)

// https://wnanbei.github.io/post/go-%E5%B9%B6%E5%8F%91%E5%AE%89%E5%85%A8%E7%9A%84-sync.map/
// sync.Map 是标准库 sync 中实现的并发安全的 map。
//
// 操作 	            普通map 	             sync.Map
// map获取某个key 	  map[1] 	            sync.Load(1)
// map添加元素 	      map[1] = 10 	        sync.Store(1, 10)
// map删除一个key 	  delete(map, 1) 	    sync.Delete(1)
// 遍历map 	          for…range 	        sync.Range()
//
// sync.Map 两个特有的函数:
// LoadOrStore - sync.Map 存在就返回，不存在就插入
// LoadAndDelete - sync.Map 获取某个 key，如果存在的话，同时删除这个 key
var (
	mapBooks     sync.Map // 实际存在的书，通过扫描生成 原本是 map[string]*Book 但是为了并发安全，改成sync.Map
	mapBookGroup sync.Map // 通过分析路径与深度生成的书组。不备份，也不存储到数据库。key是BookID
	MainFolder   = Folder{
		SortBy: "name",
	}
)

// ResetBookMap 将 sync.Map 的变量设置为一个新的实例。这样做会让原来的 sync.Map 实例失去引用，随后被垃圾回收器清理。
func ResetBookMap() {
	mapBooks = sync.Map{}
	mapBookGroup = sync.Map{}
	MainFolder = Folder{
		SubFolders: sync.Map{},
		SortBy:     "name",
	}
}

// Book 定义书籍，BooID不应该重复，根据文件路径生成
type Book struct {
	BookInfo
	Pages Pages `json:"pages"` // storm:"inline" 内联字段，结构体嵌套时使用
}

type BookInterface interface {
	GetAuthor(bookID string) string
	GetPageCount(bookID string) int
	GetBookInfo(bookID string) *BookInfo
	GetFileData(bookID string, resourceURI string) string
}

// CheckBookExist 查看内存中是否已经有了这本书,有了就false，让调用者跳过
func CheckBookExist(filePath string, bookType SupportFileType, storePath string) (exit bool) {
	// 如果是文件夹，就不用检查了
	if bookType == TypeDir || bookType == TypeBooksGroup {
		exit = false
		return exit
	}
	// 实际存在的书，通过扫描生成
	for _, value := range mapBooks.Range {
		// id := key.(string)
		realBook := value.(*Book)
		fileAbaPath, err := filepath.Abs(filePath)
		if err != nil {
			logger.Info(err, fileAbaPath)
			if realBook.FilePath == filePath && realBook.ParentFolder == storePath && realBook.Type == bookType {
				exit = true
			}
		} else {
			if realBook.FilePath == fileAbaPath && realBook.Type == bookType {
				exit = true
			}
		}
	}
	return exit
}

// NewBook  初始化Book，设置文件路径、书名、BookID等等
func NewBook(filePath string, modified time.Time, fileSize int64, storePath string, depth int, bookType SupportFileType) (*Book, error) {
	if CheckBookExist(filePath, bookType, storePath) {
		return nil, errors.New("skip:" + filePath)
	}
	// 初始化书籍
	b := Book{
		BookInfo: BookInfo{
			Author:        "",
			Modified:      modified,
			FileSize:      fileSize,
			InitComplete:  false,
			Depth:         depth,
			BookStorePath: storePath,
			Type:          bookType,
		},
	}
	// 方法链： https://colobu.com/gotips/005.html
	b.setFilePath(filePath).setParentFolder(filePath).setTitle(filePath).setAuthor().initBookID()
	return &b, nil
}

// 初始化Book时，设置页数
func (b *Book) setPageNum() {
	b.PageCount = len(b.Pages.Images)
}

// 初始化Book时， 设置封面信息
func (b *Book) initClover() {
	if len(b.Pages.Images) >= 1 {
		b.Cover = b.Pages.Images[0]
	}
}

// AddBooks 添加一组书
func AddBooks(list []*Book, basePath string, minPageNum int) (err error) {
	for _, b := range list {
		if b.GetPageCount() < minPageNum {
			continue
		}
		err = AddBook(b, basePath, minPageNum)
		if err != nil {
			return err
		}
	}
	return err
}

// RestoreDatabaseBooks 从数据库中读取的书籍信息，放到内存中
func RestoreDatabaseBooks(list []*Book) (err error) {
	for _, b := range list {
		if b.Type == TypeZip || b.Type == TypeRar || b.Type == TypeCbz || b.Type == TypeCbr || b.Type == TypeTar || b.Type == TypeEpub {
			mapBooks.Store(b.BookID, b)
		}
	}
	return err
}

// AddBook 添加一本书
func AddBook(b *Book, basePath string, minPageNum int) error {
	// 没有初始化BookID
	if b.BookID == "" {
		return errors.New("add book Error：empty BookID")
	}
	// 页数不符合要求
	if b.GetPageCount() < minPageNum {
		return errors.New("add book Error：minPageNum = " + strconv.Itoa(b.GetPageCount()))
	}
	if _, ok := MainFolder.SubFolders.Load(basePath); !ok {
		if err := MainFolder.AddSubFolder(basePath); err != nil {
			logger.Infof("%s", err)
		}
	}
	// 加入到书籍总表
	mapBooks.Store(b.BookID, b)
	return MainFolder.AddBookToSubFolder(basePath, &b.BookInfo)
}

// DeleteBookByID 删除一本书
func DeleteBookByID(bookID string) {
	// 如果key存在在删除此数据；如果不存在，delete不进行操作，也不会报错
	mapBooks.Delete(bookID)
}

// GetBooksNumber 获取书籍总数，当然不包括BookGroup
func GetBooksNumber() int {
	// 用于计数的变量
	var count int
	// 遍历 map 并递增计数器
	for _, _ = range mapBooks.Range {
		count++
	}
	return count
}

func GetAllBookList() []*Book {
	var list []*Book
	// 加上所有真实书籍
	for _, value := range mapBooks.Range {
		b := value.(*Book)
		list = append(list, b)
	}
	return list
}

func GetArchiveBooks() []*Book {
	var list []*Book
	// 所有真实书籍
	for _, value := range mapBooks.Range {
		b := value.(*Book)
		if b.Type == TypeZip || b.Type == TypeRar || b.Type == TypeCbz || b.Type == TypeCbr || b.Type == TypeTar || b.Type == TypeEpub {
			list = append(list, b)
		}
	}
	return list
}

// GetBookByID 获取特定书籍，复制一份数据
func GetBookByID(id string, sortBy string) (*Book, error) {
	// 根据id查找
	b, ok := mapBooks.Load(id)
	if ok {
		b := b.(*Book)
		b.SortPages(sortBy)
		return b, nil
	}
	g, ok := mapBookGroup.Load(id)
	if ok {
		g := g.(*BookGroup)
		temp := Book{
			BookInfo: g.BookInfo,
		}
		return &temp, nil
	}
	return nil, errors.New("can not found book,id=" + id)
}

// GetRandomBook 随机获取一本书
func GetRandomBook() (*Book, error) {
	for _, value := range GetAllBookList() {
		return value, nil
	}
	return nil, errors.New("can not found book")
}

func GetBookGroupIDByBookID(id string) (group_id string, err error) {
	// 根据id查找
	for _, value := range mapBookGroup.Range {
		group := value.(*BookGroup)
		for _, value := range group.ChildBook.Range {
			b := value.(*BookInfo)
			if b.BookID == id {
				group_id = group.BookID
				break
			}
		}
	}
	if group_id != "" {
		return group_id, nil
	}
	return "", errors.New("can not found group,id=" + id)
}

func GetBookGroupInfoByChildBookID(id string) (g *BookGroup, err error) {
	// 根据id查找
	for _, value := range mapBookGroup.Range {
		group := value.(*BookGroup)
		for _, value := range group.ChildBook.Range {
			b := value.(*BookInfo)
			if b.BookID == id {
				g = group
				break
			}
		}
	}
	if g != nil {
		return g, nil
	}
	return nil, errors.New("can not found group,id=" + id)
}

// GetBookByAuthor 获取同一作者的书籍。
func GetBookByAuthor(author string, sortBy string) ([]*Book, error) {
	var bookList []*Book
	for _, value := range mapBooks.Range {
		b := value.(*Book)
		if b.Author == author {
			b.SortPages(sortBy)
			bookList = append(bookList, b)
		}
	}

	if len(bookList) > 0 {
		return bookList, nil
	}
	return nil, errors.New("can not found book,author=" + author)
}

type Pages struct {
	Images []ImageInfo `json:"images"`
	SortBy string      `json:"sort_by"`
}

func (s Pages) Len() int {
	return len(s.Images)
}

// Less 按时间或URL，将图片排序
func (s Pages) Less(i, j int) (less bool) {
	// 如何定义 Images[i] < Images[j]
	switch s.SortBy {
	case "filename": // 根据文件名(自然语言字符串)
		return util.Compare(s.Images[i].NameInArchive, s.Images[j].NameInArchive)
	case "filesize": // 根据文件大小
		return s.Images[i].FileSize < s.Images[j].FileSize
	case "modify_time": // 根据修改时间
		return s.Images[i].ModeTime.Before(s.Images[j].ModeTime) // Images[i] 的修改时间，是否比 Images[j] 晚
	// 如何定义 Images[i] < Images[j](反向)
	case "filename_reverse": // 根据文件名(反向)
		return !util.Compare(s.Images[i].NameInArchive, s.Images[j].NameInArchive)
	case "filesize_reverse": // 根据文件大小(反向)
		return !(s.Images[i].FileSize < s.Images[j].FileSize)
	case "modify_time_reverse": // 根据修改时间(反向)
		return !s.Images[i].ModeTime.Before(s.Images[j].ModeTime) // Images[i] 的修改时间，是否比 Images[j] 晚
	default: // 默认根据文件名
		return util.Compare(s.Images[i].NameInArchive, s.Images[j].NameInArchive)
	}
}

func (s Pages) Swap(i, j int) {
	s.Images[i], s.Images[j] = s.Images[j], s.Images[i]
}

// SortPages 上面三个函数定义好了，终于可以使用sort包排序了
func (b *Book) SortPages(s string) {
	if b.Type == TypeEpub && s == "default" {
		return
	}
	if s != "" {
		b.Pages.SortBy = s
		sort.Sort(b.Pages)
	}
	b.initClover() // 重新排序后重新设置封面
}

// SortPagesByImageList 根据一个既定的文件列表，重新对页面排序。用于epub文件。
func (b *Book) SortPagesByImageList(imageList []string) {
	if len(imageList) == 0 {
		return
	}
	imageInfos := b.Pages.Images
	// 如果在有序表中，按照有序表的顺序重排
	var reSortList []ImageInfo
	for i := 0; i < len(imageList); i++ {
		checkSrc := imageList[i]
		for j := 0; j < len(imageInfos); j++ {
			if imageInfos[j].NameInArchive == checkSrc {
				reSortList = append(reSortList, imageInfos[j])
			}
		}
	}
	if len(reSortList) == 0 {
		logger.Infof(locale.GetString("EPUB_CANNOT_RESORT")+"%s", b.FilePath)
		return
	}
	// 不在表中的话，就不改变顺序，并加在有序表的后面
	for i := 0; i < len(imageInfos); i++ {
		checkName := imageInfos[i].NameInArchive
		find := false
		for j := 0; j < len(imageList); j++ {
			if imageList[j] == checkName {
				find = true
			}
		}
		if !find {
			reSortList = append(reSortList, imageInfos[i])
		}
	}
	b.Pages.Images = reSortList
	b.initClover() // 重新排序后重新设置封面
}

func md5string(s string) string {
	r := md5.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func getShortBookID(fullID string, minLength int) string {
	if len(fullID) <= minLength {
		logger.Infof("can not short ID:%s", fullID)
		return fullID
	}
	shortID := fullID[0:minLength]
	notFound := true
	add := 0
	pass := false
	for notFound {
		pass = true
		for _, value := range mapBooks.Range {
			b := value.(*Book)
			if shortID == b.BookID {
				add++
				shortID = fullID[0 : minLength+add]
				pass = false
			}
		}
		for _, value := range mapBookGroup.Range {
			group := value.(*BookGroup)
			if shortID == group.BookID {
				add++
				shortID = fullID[0 : minLength+add]
				pass = false
			}
		}
		if pass {
			notFound = false
			return shortID
		}
	}
	return fullID
}

// GetBookID  根据路径的MD5，生成书籍ID
func (b *Book) GetBookID() string {
	// 防止未初始化，最好不要用到
	if b.BookID == "" {
		logger.Infof("%s", "BookID未初始化，估计是哪里写错了")
		b.initBookID()
	}
	return b.BookID
}

// GetAuthor  获取作者信息
func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetPageCount() int {
	b.initClover()
	if !b.InitComplete {
		// 设置页数
		b.setPageNum()
		b.InitComplete = true
	}
	return b.PageCount
}

func (b *Book) GetFilePath() string {
	return b.FilePath
}

// ScanAllImage 服务器端分析分辨率、漫画单双页，只适合已解压文件
func (b *Book) ScanAllImage() {
	log.Println(locale.GetString("check_image_start"))
	// Console progress bar
	bar := pb.StartNew(b.GetPageCount())
	tmpl := `{{ red "With funcs:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }} {{percent .}} {{string . "my_green_string" | green}} {{string . "my_blue_string" | blue}}`
	bar.SetTemplateString(tmpl)
	for i := 0; i < len(b.Pages.Images); i++ { // 此处不能用range，因为会修改b.Pages.Images本身
		analyzePageImages(&b.Pages.Images[i], b.FilePath)
		// 进度条计数
		bar.Increment()
	}
	// 进度条跑完
	bar.Finish()
	log.Println(locale.GetString("check_image_completed"))
}

// ScanAllImageGo 并发分析
func (b *Book) ScanAllImageGo() {
	// var wg sync.WaitGroup
	log.Println(locale.GetString("check_image_start"))
	wp := workpool.New(10) // 设置最大线程数
	// res := make(chan string)
	count := 0
	// Console progress bar
	bar := pb.StartNew(b.GetPageCount())
	for i := 0; i < len(b.Pages.Images); i++ { // 此处不能用range，因为会修改b.Pages.Images本身
		// wg.Add(1)
		count++
		ii := i
		// 并发处理，提升图片分析速度
		wp.Do(func() error {
			// defer wg.Done()
			analyzePageImages(&b.Pages.Images[ii], b.FilePath)
			bar.Increment()
			// res <- fmt.Sprintf("Finished %d", i)
			return nil
		})
	}
	// wg.Wait()
	_ = wp.Wait()
	// finish bar
	bar.Finish()
	log.Println(locale.GetString("check_image_completed"))
}

// analyzePageImages 解析漫画的分辨率与blurhash
func analyzePageImages(p *ImageInfo, bookPath string) {
	err := p.analyzeImage(bookPath)
	// log.Println(locale.GetString("check_image_ing"), p.RealImageFilePATH)
	if err != nil {
		log.Println(locale.GetString("check_image_error") + err.Error())
	}
	if p.Width == 0 && p.Height == 0 {
		p.ImgType = "UnKnow"
		return
	}
	if p.Width > p.Height {
		p.ImgType = "DoublePage"
	} else {
		p.ImgType = "SinglePage"
	}
}

// ClearTempFilesALL web加载时保存的临时图片，在在退出后清理
func ClearTempFilesALL(debug bool, cacheFilePath string) {
	for _, value := range mapBooks.Range {
		tempBook := value.(*Book)
		clearTempFilesOne(debug, cacheFilePath, tempBook)
	}
}

// 清空某一本压缩漫画的解压缓存
func clearTempFilesOne(debug bool, cacheFilePath string, book *Book) {
	// logger.Infof(locale.GetString("clear_temp_file_start"))
	haveThisBook := false
	for _, value := range mapBooks.Range {
		tempBook := value.(*Book)
		if tempBook.GetBookID() == book.GetBookID() {
			haveThisBook = true
		}
	}
	if haveThisBook {
		cachePath := path.Join(cacheFilePath, book.GetBookID())
		err := os.RemoveAll(cachePath)
		if err != nil {
			logger.Infof(locale.GetString("clear_temp_file_error")+"%s", cachePath)
		} else {
			if debug {
				logger.Infof(locale.GetString("clear_temp_file_completed")+"%s", cachePath)
			}
		}
	}
}
