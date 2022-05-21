// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yumenaka/comi/ent/predicate"
	"github.com/yumenaka/comi/ent/singlepageinfo"
)

// SinglePageInfoQuery is the builder for querying SinglePageInfo entities.
type SinglePageInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SinglePageInfo
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SinglePageInfoQuery builder.
func (spiq *SinglePageInfoQuery) Where(ps ...predicate.SinglePageInfo) *SinglePageInfoQuery {
	spiq.predicates = append(spiq.predicates, ps...)
	return spiq
}

// Limit adds a limit step to the query.
func (spiq *SinglePageInfoQuery) Limit(limit int) *SinglePageInfoQuery {
	spiq.limit = &limit
	return spiq
}

// Offset adds an offset step to the query.
func (spiq *SinglePageInfoQuery) Offset(offset int) *SinglePageInfoQuery {
	spiq.offset = &offset
	return spiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spiq *SinglePageInfoQuery) Unique(unique bool) *SinglePageInfoQuery {
	spiq.unique = &unique
	return spiq
}

// Order adds an order step to the query.
func (spiq *SinglePageInfoQuery) Order(o ...OrderFunc) *SinglePageInfoQuery {
	spiq.order = append(spiq.order, o...)
	return spiq
}

// First returns the first SinglePageInfo entity from the query.
// Returns a *NotFoundError when no SinglePageInfo was found.
func (spiq *SinglePageInfoQuery) First(ctx context.Context) (*SinglePageInfo, error) {
	nodes, err := spiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{singlepageinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) FirstX(ctx context.Context) *SinglePageInfo {
	node, err := spiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SinglePageInfo ID from the query.
// Returns a *NotFoundError when no SinglePageInfo ID was found.
func (spiq *SinglePageInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{singlepageinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := spiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SinglePageInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SinglePageInfo entity is found.
// Returns a *NotFoundError when no SinglePageInfo entities are found.
func (spiq *SinglePageInfoQuery) Only(ctx context.Context) (*SinglePageInfo, error) {
	nodes, err := spiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{singlepageinfo.Label}
	default:
		return nil, &NotSingularError{singlepageinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) OnlyX(ctx context.Context) *SinglePageInfo {
	node, err := spiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SinglePageInfo ID in the query.
// Returns a *NotSingularError when more than one SinglePageInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (spiq *SinglePageInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = &NotSingularError{singlepageinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := spiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SinglePageInfos.
func (spiq *SinglePageInfoQuery) All(ctx context.Context) ([]*SinglePageInfo, error) {
	if err := spiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return spiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) AllX(ctx context.Context) []*SinglePageInfo {
	nodes, err := spiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SinglePageInfo IDs.
func (spiq *SinglePageInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := spiq.Select(singlepageinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := spiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spiq *SinglePageInfoQuery) Count(ctx context.Context) (int, error) {
	if err := spiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return spiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) CountX(ctx context.Context) int {
	count, err := spiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spiq *SinglePageInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := spiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return spiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (spiq *SinglePageInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := spiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SinglePageInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spiq *SinglePageInfoQuery) Clone() *SinglePageInfoQuery {
	if spiq == nil {
		return nil
	}
	return &SinglePageInfoQuery{
		config:     spiq.config,
		limit:      spiq.limit,
		offset:     spiq.offset,
		order:      append([]OrderFunc{}, spiq.order...),
		predicates: append([]predicate.SinglePageInfo{}, spiq.predicates...),
		// clone intermediate query.
		sql:    spiq.sql.Clone(),
		path:   spiq.path,
		unique: spiq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BookID string `json:"BookID,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SinglePageInfo.Query().
//		GroupBy(singlepageinfo.FieldBookID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (spiq *SinglePageInfoQuery) GroupBy(field string, fields ...string) *SinglePageInfoGroupBy {
	group := &SinglePageInfoGroupBy{config: spiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := spiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return spiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BookID string `json:"BookID,omitempty"`
//	}
//
//	client.SinglePageInfo.Query().
//		Select(singlepageinfo.FieldBookID).
//		Scan(ctx, &v)
//
func (spiq *SinglePageInfoQuery) Select(fields ...string) *SinglePageInfoSelect {
	spiq.fields = append(spiq.fields, fields...)
	return &SinglePageInfoSelect{SinglePageInfoQuery: spiq}
}

func (spiq *SinglePageInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range spiq.fields {
		if !singlepageinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spiq.path != nil {
		prev, err := spiq.path(ctx)
		if err != nil {
			return err
		}
		spiq.sql = prev
	}
	return nil
}

func (spiq *SinglePageInfoQuery) sqlAll(ctx context.Context) ([]*SinglePageInfo, error) {
	var (
		nodes   = []*SinglePageInfo{}
		withFKs = spiq.withFKs
		_spec   = spiq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, singlepageinfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SinglePageInfo{config: spiq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, spiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (spiq *SinglePageInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spiq.querySpec()
	_spec.Node.Columns = spiq.fields
	if len(spiq.fields) > 0 {
		_spec.Unique = spiq.unique != nil && *spiq.unique
	}
	return sqlgraph.CountNodes(ctx, spiq.driver, _spec)
}

func (spiq *SinglePageInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := spiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (spiq *SinglePageInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   singlepageinfo.Table,
			Columns: singlepageinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: singlepageinfo.FieldID,
			},
		},
		From:   spiq.sql,
		Unique: true,
	}
	if unique := spiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := spiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, singlepageinfo.FieldID)
		for i := range fields {
			if fields[i] != singlepageinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spiq *SinglePageInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spiq.driver.Dialect())
	t1 := builder.Table(singlepageinfo.Table)
	columns := spiq.fields
	if len(columns) == 0 {
		columns = singlepageinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spiq.sql != nil {
		selector = spiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spiq.unique != nil && *spiq.unique {
		selector.Distinct()
	}
	for _, p := range spiq.predicates {
		p(selector)
	}
	for _, p := range spiq.order {
		p(selector)
	}
	if offset := spiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SinglePageInfoGroupBy is the group-by builder for SinglePageInfo entities.
type SinglePageInfoGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spigb *SinglePageInfoGroupBy) Aggregate(fns ...AggregateFunc) *SinglePageInfoGroupBy {
	spigb.fns = append(spigb.fns, fns...)
	return spigb
}

// Scan applies the group-by query and scans the result into the given value.
func (spigb *SinglePageInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := spigb.path(ctx)
	if err != nil {
		return err
	}
	spigb.sql = query
	return spigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := spigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(spigb.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := spigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := spigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = spigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) StringX(ctx context.Context) string {
	v, err := spigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(spigb.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := spigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := spigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = spigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) IntX(ctx context.Context) int {
	v, err := spigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(spigb.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := spigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := spigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = spigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) Float64X(ctx context.Context) float64 {
	v, err := spigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(spigb.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := spigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := spigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (spigb *SinglePageInfoGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = spigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (spigb *SinglePageInfoGroupBy) BoolX(ctx context.Context) bool {
	v, err := spigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (spigb *SinglePageInfoGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range spigb.fields {
		if !singlepageinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := spigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (spigb *SinglePageInfoGroupBy) sqlQuery() *sql.Selector {
	selector := spigb.sql.Select()
	aggregation := make([]string, 0, len(spigb.fns))
	for _, fn := range spigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(spigb.fields)+len(spigb.fns))
		for _, f := range spigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(spigb.fields...)...)
}

// SinglePageInfoSelect is the builder for selecting fields of SinglePageInfo entities.
type SinglePageInfoSelect struct {
	*SinglePageInfoQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (spis *SinglePageInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if err := spis.prepareQuery(ctx); err != nil {
		return err
	}
	spis.sql = spis.SinglePageInfoQuery.sqlQuery(ctx)
	return spis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (spis *SinglePageInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := spis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(spis.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := spis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (spis *SinglePageInfoSelect) StringsX(ctx context.Context) []string {
	v, err := spis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = spis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (spis *SinglePageInfoSelect) StringX(ctx context.Context) string {
	v, err := spis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(spis.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := spis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (spis *SinglePageInfoSelect) IntsX(ctx context.Context) []int {
	v, err := spis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = spis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (spis *SinglePageInfoSelect) IntX(ctx context.Context) int {
	v, err := spis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(spis.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := spis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (spis *SinglePageInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := spis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = spis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (spis *SinglePageInfoSelect) Float64X(ctx context.Context) float64 {
	v, err := spis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(spis.fields) > 1 {
		return nil, errors.New("ent: SinglePageInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := spis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (spis *SinglePageInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := spis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (spis *SinglePageInfoSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = spis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{singlepageinfo.Label}
	default:
		err = fmt.Errorf("ent: SinglePageInfoSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (spis *SinglePageInfoSelect) BoolX(ctx context.Context) bool {
	v, err := spis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (spis *SinglePageInfoSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := spis.sql.Query()
	if err := spis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
