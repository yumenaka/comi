//go:build !(windows && 386)

package database // Package database 编译条件的注释和 package 语句之间一定要隔一行，不然无法识别编译条件。go:build 是1.18以后“条件编译”的推荐语法。

import (
	"context"
	"errors"
	"fmt"
	"github.com/yumenaka/comi/ent"
	"github.com/yumenaka/comi/ent/book"
	"github.com/yumenaka/comi/ent/singlepageinfo"
	"github.com/yumenaka/comi/types"
	"strconv"
)

// ClearBookData   清空数据库的Book与SinglePageInfo表  // 后台并发执行，所以不能保证结果如预期，不用这个函数???
func ClearBookData(clearBook *types.Book) {
	//如何增删查改： https://entgo.io/zh/docs/crud
	ctx := context.Background()
	_, err := client.Book.
		Delete().
		Where(book.BookIDEQ(clearBook.BookID)).
		Exec(ctx)
	if err != nil {
		fmt.Println("ClearBookData Book:" + err.Error())
	}
	fmt.Println("Clear Book ：" + clearBook.Name)
	deletePageInfoNum, err := client.SinglePageInfo.
		Delete().
		Where(singlepageinfo.BookIDEQ(clearBook.BookID)).
		Exec(ctx)
	if err != nil {
		fmt.Println("ClearBookData SinglePageInfo:" + err.Error())
	}
	fmt.Println("Clear SinglePageInfo Num：" + strconv.Itoa(deletePageInfoNum))
}

// DeleteAllBookInDatabase  清空数据库的Book与SinglePageInfo表
// 后台并发执行，不能保证结果如预期，不用这个函数。
func DeleteAllBookInDatabase(debug bool) {
	//如何增删查改： https://entgo.io/zh/docs/crud
	ctx := context.Background()
	deleteBookNum, err := client.Book.
		Delete().
		Where(book.AllPageNumNEQ(-99999)).
		Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if debug {
		fmt.Println("Delete Book Num：" + strconv.Itoa(deleteBookNum))
	}
	deletePageInfoNum, err := client.SinglePageInfo.
		Delete().
		Where(singlepageinfo.WidthNEQ(-99999)).
		Exec(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if debug {
		fmt.Println("Delete SinglePageInfo Num：" + strconv.Itoa(deletePageInfoNum))
	}
}

// SaveAllBookToDatabase 将Map里面的书籍信息，全部保存到本地数据库中
func SaveAllBookToDatabase(m map[string]*types.Book) {
	for _, b := range m {
		var c = *b
		err := SaveBookToDatabase(&c)
		if err != nil {
			fmt.Println("SaveAllBookToDatabase error :" + err.Error())
		}
	}
}

// SaveBookListToDatabase  向数据库中插入一组书
func SaveBookListToDatabase(bookList []*types.Book) error {
	for _, b := range bookList {
		err := SaveBookToDatabase(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveBookToDatabase 向数据库中插入一本书
func SaveBookToDatabase(save *types.Book) error {
	//如何增删查改： https://entgo.io/zh/docs/crud
	ctx := context.Background()
	b, err := client.Book.
		Create().
		SetName(save.BookInfo.Name).
		SetBookID(save.BookInfo.BookID).
		SetOwner("").
		SetFilePath(save.BookInfo.FilePath).
		SetBookStorePath(save.BookInfo.BookStorePath).
		SetChildBookNum(save.BookInfo.ChildBookNum).
		SetType(string(save.BookInfo.Type)).
		SetDepth(save.BookInfo.Depth).
		SetParentFolder(save.BookInfo.ParentFolder).
		SetAllPageNum(save.BookInfo.AllPageNum).
		SetFileSize(save.BookInfo.FileSize).
		SetAuthors(save.GetAuthor()).
		SetISBN(save.BookInfo.ISBN).
		SetPress(save.BookInfo.Press).
		SetPublishedAt(save.BookInfo.PublishedAt).
		SetExtractPath(save.BookInfo.ExtractPath).
		SetInitComplete(save.BookInfo.InitComplete).
		SetReadPercent(save.BookInfo.ReadPercent).
		SetNonUTF8Zip(save.BookInfo.NonUTF8Zip).
		SetZipTextEncoding(save.BookInfo.ZipTextEncoding).
		SetExtractNum(save.BookInfo.ExtractNum).
		Save(ctx) // 创建并返回 //还有一个SaveX(ctx)，和 Save() 不一样， SaveX 在出错时 panic。
	if err != nil {
		//log.Fatalf("failed creating book: %v", err)
		return err
	}

	//保存封面与页面信息
	bulk := make([]*ent.SinglePageInfoCreate, len(save.Pages.Images))
	for i, p := range save.Pages.Images {
		bulk[i] = client.SinglePageInfo.
			Create().
			SetBookID(save.BookID).
			SetPageNum(p.PageNum).
			SetNameInArchive(p.NameInArchive).
			SetURL(p.Url).
			SetBlurHash(p.Blurhash).
			SetHeight(p.Height).
			SetWidth(p.Width).
			SetModeTime(p.ModeTime).
			SetFileSize(p.FileSize).
			SetRealImageFilePATH(p.RealImageFilePATH).
			SetImgType(p.ImgType)
	}
	pages, err := client.SinglePageInfo.CreateBulk(bulk...).Save(ctx)
	if b != nil && pages != nil {
		//log.Println("book was created: ", b)
		//log.Println("book pages info was created: ", pages)
	}
	return err
}

// GetBookFromDatabase 根据文件路径，从数据库查询一本书的详细信息,避免重复扫描压缩包？
func GetBookFromDatabase(filepath string) (*types.Book, error) {
	ctx := context.Background()
	books, err := client.Book. // UserClient.
					Query(). // 用户查询生成器。
					Where(book.FilePath(filepath)).
					All(ctx) // query and return.
	if err != nil {
		fmt.Println(err)
	}
	if len(books) == 0 {
		return nil, errors.New("not found in database,filepath:" + filepath)
	}
	temp := books[0]
	b := types.Book{
		BookInfo: types.BookInfo{
			Name:            temp.Name,
			BookID:          temp.BookID,
			FilePath:        temp.FilePath,
			BookStorePath:   temp.BookStorePath,
			Type:            types.SupportFileType(temp.Type),
			ChildBookNum:    temp.ChildBookNum,
			Depth:           temp.Depth,
			ParentFolder:    temp.ParentFolder,
			AllPageNum:      temp.AllPageNum,
			FileSize:        temp.FileSize,
			ISBN:            temp.ISBN,
			Press:           temp.Press,
			PublishedAt:     temp.PublishedAt,
			ExtractPath:     temp.ExtractPath,
			Modified:        temp.Modified,
			ExtractNum:      temp.ExtractNum,
			InitComplete:    temp.InitComplete,
			ReadPercent:     temp.ReadPercent,
			NonUTF8Zip:      temp.NonUTF8Zip,
			ZipTextEncoding: temp.ZipTextEncoding,
		},
	}

	//查询数据库里的封面与页面信息
	//https://entgo.io/zh/docs/crud
	pages, err := client.SinglePageInfo. // UserClient.
						Query(). // 用户查询生成器。
						Where(singlepageinfo.BookID(temp.BookID)).
						All(ctx) // query and return.
	for _, v := range pages {
		b.Pages.Images = append(b.Pages.Images, types.ImageInfo{
			PageNum:           v.PageNum,
			NameInArchive:     v.NameInArchive,
			Url:               v.URL,
			Blurhash:          v.BlurHash,
			Height:            v.Height,
			Width:             v.Width,
			ModeTime:          v.ModeTime,
			FileSize:          v.FileSize,
			RealImageFilePATH: v.RealImageFilePATH,
			ImgType:           v.ImgType,
		})
	}
	//设置封面
	if len(b.Pages.Images) > 0 {
		b.Cover = b.Pages.Images[0]
	}
	if err != nil {
		fmt.Println(err)
	}
	return &b, err
}

// GetBooksFromDatabase  根据文件路径，从数据库查询书的详细信息,避免重复扫描压缩包。//忽略文件夹型的书籍
func GetBooksFromDatabase() (list []*types.Book, err error) {
	ctx := context.Background()
	books, err := client.Book. // UserClient.
					Query(). // 用户查询生成器。
		//Where(book.Not(book.Type("dir"))). //忽略文件夹型的书籍
		All(ctx) // query and return.
	if err != nil {
		fmt.Println(err)
	}
	if len(books) == 0 {
		return nil, errors.New("not found in database")
	}
	for _, temp := range books {
		b := types.Book{
			BookInfo: types.BookInfo{
				Name:            temp.Name,
				BookID:          temp.BookID,
				FilePath:        temp.FilePath,
				BookStorePath:   temp.BookStorePath,
				Type:            types.SupportFileType(temp.Type),
				ChildBookNum:    temp.ChildBookNum,
				Depth:           temp.Depth,
				ParentFolder:    temp.ParentFolder,
				AllPageNum:      temp.AllPageNum,
				FileSize:        temp.FileSize,
				ISBN:            temp.ISBN,
				Press:           temp.Press,
				PublishedAt:     temp.PublishedAt,
				ExtractPath:     temp.ExtractPath,
				Modified:        temp.Modified,
				ExtractNum:      temp.ExtractNum,
				InitComplete:    temp.InitComplete,
				ReadPercent:     temp.ReadPercent,
				NonUTF8Zip:      temp.NonUTF8Zip,
				ZipTextEncoding: temp.ZipTextEncoding,
			},
		}
		//查询数据库里的封面与页面信息
		//https://entgo.io/zh/docs/crud
		pages, err := client.SinglePageInfo. // UserClient.
							Query(). // 用户查询生成器。
							Where(singlepageinfo.BookID(temp.BookID)).
							All(ctx) // query and return.
		if err != nil {
			fmt.Println(err)
		}
		for _, v := range pages {
			b.Pages.Images = append(b.Pages.Images, types.ImageInfo{
				PageNum:           v.PageNum,
				NameInArchive:     v.NameInArchive,
				Url:               v.URL,
				Blurhash:          v.BlurHash,
				Height:            v.Height,
				Width:             v.Width,
				ModeTime:          v.ModeTime,
				FileSize:          v.FileSize,
				RealImageFilePATH: v.RealImageFilePATH,
				ImgType:           v.ImgType,
			})
		}
		//设置封面
		if len(b.Pages.Images) > 0 {
			b.Cover = b.Pages.Images[0]
		}
		//硬写一个封面
		switch b.Type {
		case types.TypePDF:
			b.Cover = types.ImageInfo{NameInArchive: "pdf.png", Url: "/images/pdf.png"}
		case types.TypeVideo:
			b.Cover = types.ImageInfo{NameInArchive: "video.png", Url: "/images/video.png"}
		case types.TypeAudio:
			b.Cover = types.ImageInfo{NameInArchive: "audio.png", Url: "/images/audio.png"}
		case types.TypeUnknownFile:
			b.Cover = types.ImageInfo{NameInArchive: "unknown.png", Url: "/images/unknown.png"}
		}
		list = append(list, &b)
	}
	return list, err
}

// todo： 根据扫描完成的书籍数据，清理本地数据库当中不存在的书籍