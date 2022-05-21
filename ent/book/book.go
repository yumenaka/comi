// Code generated by entc, DO NOT EDIT.

package book

import (
	"time"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldBookID holds the string denoting the bookid field in the database.
	FieldBookID = "book_id"
	// FieldFilePath holds the string denoting the filepath field in the database.
	FieldFilePath = "file_path"
	// FieldBookStorePath holds the string denoting the bookstorepath field in the database.
	FieldBookStorePath = "book_store_path"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldChildBookNum holds the string denoting the childbooknum field in the database.
	FieldChildBookNum = "child_book_num"
	// FieldDepth holds the string denoting the depth field in the database.
	FieldDepth = "depth"
	// FieldParentFolder holds the string denoting the parentfolder field in the database.
	FieldParentFolder = "parent_folder"
	// FieldAllPageNum holds the string denoting the allpagenum field in the database.
	FieldAllPageNum = "all_page_num"
	// FieldFileSize holds the string denoting the filesize field in the database.
	FieldFileSize = "file_size"
	// FieldAuthors holds the string denoting the authors field in the database.
	FieldAuthors = "authors"
	// FieldISBN holds the string denoting the isbn field in the database.
	FieldISBN = "isbn"
	// FieldPress holds the string denoting the press field in the database.
	FieldPress = "press"
	// FieldPublishedAt holds the string denoting the publishedat field in the database.
	FieldPublishedAt = "published_at"
	// FieldExtractPath holds the string denoting the extractpath field in the database.
	FieldExtractPath = "extract_path"
	// FieldModified holds the string denoting the modified field in the database.
	FieldModified = "modified"
	// FieldExtractNum holds the string denoting the extractnum field in the database.
	FieldExtractNum = "extract_num"
	// FieldInitComplete holds the string denoting the initcomplete field in the database.
	FieldInitComplete = "init_complete"
	// FieldReadPercent holds the string denoting the readpercent field in the database.
	FieldReadPercent = "read_percent"
	// FieldNonUTF8Zip holds the string denoting the nonutf8zip field in the database.
	FieldNonUTF8Zip = "non_utf8zip"
	// FieldZipTextEncoding holds the string denoting the ziptextencoding field in the database.
	FieldZipTextEncoding = "zip_text_encoding"
	// EdgePageInfos holds the string denoting the pageinfos edge name in mutations.
	EdgePageInfos = "PageInfos"
	// Table holds the table name of the book in the database.
	Table = "books"
	// PageInfosTable is the table that holds the PageInfos relation/edge.
	PageInfosTable = "single_page_infos"
	// PageInfosInverseTable is the table name for the SinglePageInfo entity.
	// It exists in this package in order to avoid circular dependency with the "singlepageinfo" package.
	PageInfosInverseTable = "single_page_infos"
	// PageInfosColumn is the table column denoting the PageInfos relation/edge.
	PageInfosColumn = "book_page_infos"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldBookID,
	FieldFilePath,
	FieldBookStorePath,
	FieldType,
	FieldChildBookNum,
	FieldDepth,
	FieldParentFolder,
	FieldAllPageNum,
	FieldFileSize,
	FieldAuthors,
	FieldISBN,
	FieldPress,
	FieldPublishedAt,
	FieldExtractPath,
	FieldModified,
	FieldExtractNum,
	FieldInitComplete,
	FieldReadPercent,
	FieldNonUTF8Zip,
	FieldZipTextEncoding,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ChildBookNumValidator is a validator for the "ChildBookNum" field. It is called by the builders before save.
	ChildBookNumValidator func(int) error
	// DepthValidator is a validator for the "Depth" field. It is called by the builders before save.
	DepthValidator func(int) error
	// AllPageNumValidator is a validator for the "AllPageNum" field. It is called by the builders before save.
	AllPageNumValidator func(int) error
	// DefaultModified holds the default value on creation for the "Modified" field.
	DefaultModified func() time.Time
)
