package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SinglePageInfo holds the schema definition for the SinglePageInfo entity.
type SinglePageInfo struct {
	ent.Schema
}

// Fields of the SinglePageInfo.
func (SinglePageInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("BookID"),
		field.Int("PageNum"),
		field.String("NameInArchive"),
		field.String("Url"),
		field.String("BlurHash"),
		field.Int("Height"),
		field.Int("Width"),
		field.Time("ModeTime").Default(time.Now),
		field.Float("FileSize"),
		field.String("RealImageFilePATH"),
		field.String("ImgType"),
	}
}

// Edges of the SinglePageInfo.
func (SinglePageInfo) Edges() []ent.Edge {
	return nil
}
