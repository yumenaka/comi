// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yumenaka/comigo/internal/ent/predicate"
	"github.com/yumenaka/comigo/internal/ent/singlepageinfo"
)

// SinglePageInfoDelete is the builder for deleting a SinglePageInfo entity.
type SinglePageInfoDelete struct {
	config
	hooks    []Hook
	mutation *SinglePageInfoMutation
}

// Where appends a list predicates to the SinglePageInfoDelete builder.
func (spid *SinglePageInfoDelete) Where(ps ...predicate.SinglePageInfo) *SinglePageInfoDelete {
	spid.mutation.Where(ps...)
	return spid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spid *SinglePageInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, spid.sqlExec, spid.mutation, spid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (spid *SinglePageInfoDelete) ExecX(ctx context.Context) int {
	n, err := spid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spid *SinglePageInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(singlepageinfo.Table, sqlgraph.NewFieldSpec(singlepageinfo.FieldID, field.TypeInt))
	if ps := spid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, spid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	spid.mutation.done = true
	return affected, err
}

// SinglePageInfoDeleteOne is the builder for deleting a single SinglePageInfo entity.
type SinglePageInfoDeleteOne struct {
	spid *SinglePageInfoDelete
}

// Where appends a list predicates to the SinglePageInfoDelete builder.
func (spido *SinglePageInfoDeleteOne) Where(ps ...predicate.SinglePageInfo) *SinglePageInfoDeleteOne {
	spido.spid.mutation.Where(ps...)
	return spido
}

// Exec executes the deletion query.
func (spido *SinglePageInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := spido.spid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{singlepageinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spido *SinglePageInfoDeleteOne) ExecX(ctx context.Context) {
	if err := spido.Exec(ctx); err != nil {
		panic(err)
	}
}
