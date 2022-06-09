// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yumenaka/comi/ent/predicate"
	"github.com/yumenaka/comi/ent/singlepageinfo"
)

// SinglePageInfoUpdate is the builder for updating SinglePageInfo entities.
type SinglePageInfoUpdate struct {
	config
	hooks    []Hook
	mutation *SinglePageInfoMutation
}

// Where appends a list predicates to the SinglePageInfoUpdate builder.
func (spiu *SinglePageInfoUpdate) Where(ps ...predicate.SinglePageInfo) *SinglePageInfoUpdate {
	spiu.mutation.Where(ps...)
	return spiu
}

// SetBookID sets the "BookID" field.
func (spiu *SinglePageInfoUpdate) SetBookID(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetBookID(s)
	return spiu
}

// SetPageNum sets the "PageNum" field.
func (spiu *SinglePageInfoUpdate) SetPageNum(i int) *SinglePageInfoUpdate {
	spiu.mutation.ResetPageNum()
	spiu.mutation.SetPageNum(i)
	return spiu
}

// AddPageNum adds i to the "PageNum" field.
func (spiu *SinglePageInfoUpdate) AddPageNum(i int) *SinglePageInfoUpdate {
	spiu.mutation.AddPageNum(i)
	return spiu
}

// SetNameInArchive sets the "NameInArchive" field.
func (spiu *SinglePageInfoUpdate) SetNameInArchive(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetNameInArchive(s)
	return spiu
}

// SetURL sets the "Url" field.
func (spiu *SinglePageInfoUpdate) SetURL(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetURL(s)
	return spiu
}

// SetBlurHash sets the "BlurHash" field.
func (spiu *SinglePageInfoUpdate) SetBlurHash(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetBlurHash(s)
	return spiu
}

// SetHeight sets the "Height" field.
func (spiu *SinglePageInfoUpdate) SetHeight(i int) *SinglePageInfoUpdate {
	spiu.mutation.ResetHeight()
	spiu.mutation.SetHeight(i)
	return spiu
}

// AddHeight adds i to the "Height" field.
func (spiu *SinglePageInfoUpdate) AddHeight(i int) *SinglePageInfoUpdate {
	spiu.mutation.AddHeight(i)
	return spiu
}

// SetWidth sets the "Width" field.
func (spiu *SinglePageInfoUpdate) SetWidth(i int) *SinglePageInfoUpdate {
	spiu.mutation.ResetWidth()
	spiu.mutation.SetWidth(i)
	return spiu
}

// AddWidth adds i to the "Width" field.
func (spiu *SinglePageInfoUpdate) AddWidth(i int) *SinglePageInfoUpdate {
	spiu.mutation.AddWidth(i)
	return spiu
}

// SetModeTime sets the "ModeTime" field.
func (spiu *SinglePageInfoUpdate) SetModeTime(t time.Time) *SinglePageInfoUpdate {
	spiu.mutation.SetModeTime(t)
	return spiu
}

// SetNillableModeTime sets the "ModeTime" field if the given value is not nil.
func (spiu *SinglePageInfoUpdate) SetNillableModeTime(t *time.Time) *SinglePageInfoUpdate {
	if t != nil {
		spiu.SetModeTime(*t)
	}
	return spiu
}

// SetFileSize sets the "FileSize" field.
func (spiu *SinglePageInfoUpdate) SetFileSize(i int64) *SinglePageInfoUpdate {
	spiu.mutation.ResetFileSize()
	spiu.mutation.SetFileSize(i)
	return spiu
}

// AddFileSize adds i to the "FileSize" field.
func (spiu *SinglePageInfoUpdate) AddFileSize(i int64) *SinglePageInfoUpdate {
	spiu.mutation.AddFileSize(i)
	return spiu
}

// SetRealImageFilePATH sets the "RealImageFilePATH" field.
func (spiu *SinglePageInfoUpdate) SetRealImageFilePATH(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetRealImageFilePATH(s)
	return spiu
}

// SetImgType sets the "ImgType" field.
func (spiu *SinglePageInfoUpdate) SetImgType(s string) *SinglePageInfoUpdate {
	spiu.mutation.SetImgType(s)
	return spiu
}

// Mutation returns the SinglePageInfoMutation object of the builder.
func (spiu *SinglePageInfoUpdate) Mutation() *SinglePageInfoMutation {
	return spiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (spiu *SinglePageInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spiu.hooks) == 0 {
		affected, err = spiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SinglePageInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spiu.mutation = mutation
			affected, err = spiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spiu.hooks) - 1; i >= 0; i-- {
			if spiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (spiu *SinglePageInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := spiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (spiu *SinglePageInfoUpdate) Exec(ctx context.Context) error {
	_, err := spiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spiu *SinglePageInfoUpdate) ExecX(ctx context.Context) {
	if err := spiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spiu *SinglePageInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   singlepageinfo.Table,
			Columns: singlepageinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: singlepageinfo.FieldID,
			},
		},
	}
	if ps := spiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spiu.mutation.BookID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldBookID,
		})
	}
	if value, ok := spiu.mutation.PageNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldPageNum,
		})
	}
	if value, ok := spiu.mutation.AddedPageNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldPageNum,
		})
	}
	if value, ok := spiu.mutation.NameInArchive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldNameInArchive,
		})
	}
	if value, ok := spiu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldURL,
		})
	}
	if value, ok := spiu.mutation.BlurHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldBlurHash,
		})
	}
	if value, ok := spiu.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldHeight,
		})
	}
	if value, ok := spiu.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldHeight,
		})
	}
	if value, ok := spiu.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldWidth,
		})
	}
	if value, ok := spiu.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldWidth,
		})
	}
	if value, ok := spiu.mutation.ModeTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: singlepageinfo.FieldModeTime,
		})
	}
	if value, ok := spiu.mutation.FileSize(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: singlepageinfo.FieldFileSize,
		})
	}
	if value, ok := spiu.mutation.AddedFileSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: singlepageinfo.FieldFileSize,
		})
	}
	if value, ok := spiu.mutation.RealImageFilePATH(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldRealImageFilePATH,
		})
	}
	if value, ok := spiu.mutation.ImgType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldImgType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, spiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{singlepageinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SinglePageInfoUpdateOne is the builder for updating a single SinglePageInfo entity.
type SinglePageInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SinglePageInfoMutation
}

// SetBookID sets the "BookID" field.
func (spiuo *SinglePageInfoUpdateOne) SetBookID(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetBookID(s)
	return spiuo
}

// SetPageNum sets the "PageNum" field.
func (spiuo *SinglePageInfoUpdateOne) SetPageNum(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.ResetPageNum()
	spiuo.mutation.SetPageNum(i)
	return spiuo
}

// AddPageNum adds i to the "PageNum" field.
func (spiuo *SinglePageInfoUpdateOne) AddPageNum(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.AddPageNum(i)
	return spiuo
}

// SetNameInArchive sets the "NameInArchive" field.
func (spiuo *SinglePageInfoUpdateOne) SetNameInArchive(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetNameInArchive(s)
	return spiuo
}

// SetURL sets the "Url" field.
func (spiuo *SinglePageInfoUpdateOne) SetURL(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetURL(s)
	return spiuo
}

// SetBlurHash sets the "BlurHash" field.
func (spiuo *SinglePageInfoUpdateOne) SetBlurHash(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetBlurHash(s)
	return spiuo
}

// SetHeight sets the "Height" field.
func (spiuo *SinglePageInfoUpdateOne) SetHeight(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.ResetHeight()
	spiuo.mutation.SetHeight(i)
	return spiuo
}

// AddHeight adds i to the "Height" field.
func (spiuo *SinglePageInfoUpdateOne) AddHeight(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.AddHeight(i)
	return spiuo
}

// SetWidth sets the "Width" field.
func (spiuo *SinglePageInfoUpdateOne) SetWidth(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.ResetWidth()
	spiuo.mutation.SetWidth(i)
	return spiuo
}

// AddWidth adds i to the "Width" field.
func (spiuo *SinglePageInfoUpdateOne) AddWidth(i int) *SinglePageInfoUpdateOne {
	spiuo.mutation.AddWidth(i)
	return spiuo
}

// SetModeTime sets the "ModeTime" field.
func (spiuo *SinglePageInfoUpdateOne) SetModeTime(t time.Time) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetModeTime(t)
	return spiuo
}

// SetNillableModeTime sets the "ModeTime" field if the given value is not nil.
func (spiuo *SinglePageInfoUpdateOne) SetNillableModeTime(t *time.Time) *SinglePageInfoUpdateOne {
	if t != nil {
		spiuo.SetModeTime(*t)
	}
	return spiuo
}

// SetFileSize sets the "FileSize" field.
func (spiuo *SinglePageInfoUpdateOne) SetFileSize(i int64) *SinglePageInfoUpdateOne {
	spiuo.mutation.ResetFileSize()
	spiuo.mutation.SetFileSize(i)
	return spiuo
}

// AddFileSize adds i to the "FileSize" field.
func (spiuo *SinglePageInfoUpdateOne) AddFileSize(i int64) *SinglePageInfoUpdateOne {
	spiuo.mutation.AddFileSize(i)
	return spiuo
}

// SetRealImageFilePATH sets the "RealImageFilePATH" field.
func (spiuo *SinglePageInfoUpdateOne) SetRealImageFilePATH(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetRealImageFilePATH(s)
	return spiuo
}

// SetImgType sets the "ImgType" field.
func (spiuo *SinglePageInfoUpdateOne) SetImgType(s string) *SinglePageInfoUpdateOne {
	spiuo.mutation.SetImgType(s)
	return spiuo
}

// Mutation returns the SinglePageInfoMutation object of the builder.
func (spiuo *SinglePageInfoUpdateOne) Mutation() *SinglePageInfoMutation {
	return spiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (spiuo *SinglePageInfoUpdateOne) Select(field string, fields ...string) *SinglePageInfoUpdateOne {
	spiuo.fields = append([]string{field}, fields...)
	return spiuo
}

// Save executes the query and returns the updated SinglePageInfo entity.
func (spiuo *SinglePageInfoUpdateOne) Save(ctx context.Context) (*SinglePageInfo, error) {
	var (
		err  error
		node *SinglePageInfo
	)
	if len(spiuo.hooks) == 0 {
		node, err = spiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SinglePageInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spiuo.mutation = mutation
			node, err = spiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(spiuo.hooks) - 1; i >= 0; i-- {
			if spiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (spiuo *SinglePageInfoUpdateOne) SaveX(ctx context.Context) *SinglePageInfo {
	node, err := spiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (spiuo *SinglePageInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := spiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spiuo *SinglePageInfoUpdateOne) ExecX(ctx context.Context) {
	if err := spiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (spiuo *SinglePageInfoUpdateOne) sqlSave(ctx context.Context) (_node *SinglePageInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   singlepageinfo.Table,
			Columns: singlepageinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: singlepageinfo.FieldID,
			},
		},
	}
	id, ok := spiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SinglePageInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := spiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, singlepageinfo.FieldID)
		for _, f := range fields {
			if !singlepageinfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != singlepageinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := spiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := spiuo.mutation.BookID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldBookID,
		})
	}
	if value, ok := spiuo.mutation.PageNum(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldPageNum,
		})
	}
	if value, ok := spiuo.mutation.AddedPageNum(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldPageNum,
		})
	}
	if value, ok := spiuo.mutation.NameInArchive(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldNameInArchive,
		})
	}
	if value, ok := spiuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldURL,
		})
	}
	if value, ok := spiuo.mutation.BlurHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldBlurHash,
		})
	}
	if value, ok := spiuo.mutation.Height(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldHeight,
		})
	}
	if value, ok := spiuo.mutation.AddedHeight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldHeight,
		})
	}
	if value, ok := spiuo.mutation.Width(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldWidth,
		})
	}
	if value, ok := spiuo.mutation.AddedWidth(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: singlepageinfo.FieldWidth,
		})
	}
	if value, ok := spiuo.mutation.ModeTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: singlepageinfo.FieldModeTime,
		})
	}
	if value, ok := spiuo.mutation.FileSize(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: singlepageinfo.FieldFileSize,
		})
	}
	if value, ok := spiuo.mutation.AddedFileSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: singlepageinfo.FieldFileSize,
		})
	}
	if value, ok := spiuo.mutation.RealImageFilePATH(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldRealImageFilePATH,
		})
	}
	if value, ok := spiuo.mutation.ImgType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: singlepageinfo.FieldImgType,
		})
	}
	_node = &SinglePageInfo{config: spiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, spiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{singlepageinfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}