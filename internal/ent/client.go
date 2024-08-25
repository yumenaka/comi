// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/yumenaka/comigo/internal/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/yumenaka/comigo/internal/ent/book"
	"github.com/yumenaka/comigo/internal/ent/singlepageinfo"
	"github.com/yumenaka/comigo/internal/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Book is the client for interacting with the Book builders.
	Book *BookClient
	// SinglePageInfo is the client for interacting with the SinglePageInfo builders.
	SinglePageInfo *SinglePageInfoClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Book = NewBookClient(c.config)
	c.SinglePageInfo = NewSinglePageInfoClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Book:           NewBookClient(cfg),
		SinglePageInfo: NewSinglePageInfoClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Book:           NewBookClient(cfg),
		SinglePageInfo: NewSinglePageInfoClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Book.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Book.Use(hooks...)
	c.SinglePageInfo.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Book.Intercept(interceptors...)
	c.SinglePageInfo.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BookMutation:
		return c.Book.mutate(ctx, m)
	case *SinglePageInfoMutation:
		return c.SinglePageInfo.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BookClient is a client for the Book schema.
type BookClient struct {
	config
}

// NewBookClient returns a client for the Book from the given config.
func NewBookClient(c config) *BookClient {
	return &BookClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `book.Hooks(f(g(h())))`.
func (c *BookClient) Use(hooks ...Hook) {
	c.hooks.Book = append(c.hooks.Book, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `book.Intercept(f(g(h())))`.
func (c *BookClient) Intercept(interceptors ...Interceptor) {
	c.inters.Book = append(c.inters.Book, interceptors...)
}

// Create returns a builder for creating a Book entity.
func (c *BookClient) Create() *BookCreate {
	mutation := newBookMutation(c.config, OpCreate)
	return &BookCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Book entities.
func (c *BookClient) CreateBulk(builders ...*BookCreate) *BookCreateBulk {
	return &BookCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *BookClient) MapCreateBulk(slice any, setFunc func(*BookCreate, int)) *BookCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &BookCreateBulk{err: fmt.Errorf("calling to BookClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*BookCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &BookCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Book.
func (c *BookClient) Update() *BookUpdate {
	mutation := newBookMutation(c.config, OpUpdate)
	return &BookUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BookClient) UpdateOne(b *Book) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBook(b))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BookClient) UpdateOneID(id int) *BookUpdateOne {
	mutation := newBookMutation(c.config, OpUpdateOne, withBookID(id))
	return &BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Book.
func (c *BookClient) Delete() *BookDelete {
	mutation := newBookMutation(c.config, OpDelete)
	return &BookDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BookClient) DeleteOne(b *Book) *BookDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BookClient) DeleteOneID(id int) *BookDeleteOne {
	builder := c.Delete().Where(book.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BookDeleteOne{builder}
}

// Query returns a query builder for Book.
func (c *BookClient) Query() *BookQuery {
	return &BookQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBook},
		inters: c.Interceptors(),
	}
}

// Get returns a Book entity by its id.
func (c *BookClient) Get(ctx context.Context, id int) (*Book, error) {
	return c.Query().Where(book.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BookClient) GetX(ctx context.Context, id int) *Book {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPageInfos queries the PageInfos edge of a Book.
func (c *BookClient) QueryPageInfos(b *Book) *SinglePageInfoQuery {
	query := (&SinglePageInfoClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(book.Table, book.FieldID, id),
			sqlgraph.To(singlepageinfo.Table, singlepageinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, book.PageInfosTable, book.PageInfosColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BookClient) Hooks() []Hook {
	return c.hooks.Book
}

// Interceptors returns the client interceptors.
func (c *BookClient) Interceptors() []Interceptor {
	return c.inters.Book
}

func (c *BookClient) mutate(ctx context.Context, m *BookMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BookCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BookUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BookUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BookDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Book mutation op: %q", m.Op())
	}
}

// SinglePageInfoClient is a client for the SinglePageInfo schema.
type SinglePageInfoClient struct {
	config
}

// NewSinglePageInfoClient returns a client for the SinglePageInfo from the given config.
func NewSinglePageInfoClient(c config) *SinglePageInfoClient {
	return &SinglePageInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `singlepageinfo.Hooks(f(g(h())))`.
func (c *SinglePageInfoClient) Use(hooks ...Hook) {
	c.hooks.SinglePageInfo = append(c.hooks.SinglePageInfo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `singlepageinfo.Intercept(f(g(h())))`.
func (c *SinglePageInfoClient) Intercept(interceptors ...Interceptor) {
	c.inters.SinglePageInfo = append(c.inters.SinglePageInfo, interceptors...)
}

// Create returns a builder for creating a SinglePageInfo entity.
func (c *SinglePageInfoClient) Create() *SinglePageInfoCreate {
	mutation := newSinglePageInfoMutation(c.config, OpCreate)
	return &SinglePageInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SinglePageInfo entities.
func (c *SinglePageInfoClient) CreateBulk(builders ...*SinglePageInfoCreate) *SinglePageInfoCreateBulk {
	return &SinglePageInfoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SinglePageInfoClient) MapCreateBulk(slice any, setFunc func(*SinglePageInfoCreate, int)) *SinglePageInfoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SinglePageInfoCreateBulk{err: fmt.Errorf("calling to SinglePageInfoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SinglePageInfoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SinglePageInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SinglePageInfo.
func (c *SinglePageInfoClient) Update() *SinglePageInfoUpdate {
	mutation := newSinglePageInfoMutation(c.config, OpUpdate)
	return &SinglePageInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SinglePageInfoClient) UpdateOne(spi *SinglePageInfo) *SinglePageInfoUpdateOne {
	mutation := newSinglePageInfoMutation(c.config, OpUpdateOne, withSinglePageInfo(spi))
	return &SinglePageInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SinglePageInfoClient) UpdateOneID(id int) *SinglePageInfoUpdateOne {
	mutation := newSinglePageInfoMutation(c.config, OpUpdateOne, withSinglePageInfoID(id))
	return &SinglePageInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SinglePageInfo.
func (c *SinglePageInfoClient) Delete() *SinglePageInfoDelete {
	mutation := newSinglePageInfoMutation(c.config, OpDelete)
	return &SinglePageInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SinglePageInfoClient) DeleteOne(spi *SinglePageInfo) *SinglePageInfoDeleteOne {
	return c.DeleteOneID(spi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SinglePageInfoClient) DeleteOneID(id int) *SinglePageInfoDeleteOne {
	builder := c.Delete().Where(singlepageinfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SinglePageInfoDeleteOne{builder}
}

// Query returns a query builder for SinglePageInfo.
func (c *SinglePageInfoClient) Query() *SinglePageInfoQuery {
	return &SinglePageInfoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSinglePageInfo},
		inters: c.Interceptors(),
	}
}

// Get returns a SinglePageInfo entity by its id.
func (c *SinglePageInfoClient) Get(ctx context.Context, id int) (*SinglePageInfo, error) {
	return c.Query().Where(singlepageinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SinglePageInfoClient) GetX(ctx context.Context, id int) *SinglePageInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SinglePageInfoClient) Hooks() []Hook {
	return c.hooks.SinglePageInfo
}

// Interceptors returns the client interceptors.
func (c *SinglePageInfoClient) Interceptors() []Interceptor {
	return c.inters.SinglePageInfo
}

func (c *SinglePageInfoClient) mutate(ctx context.Context, m *SinglePageInfoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SinglePageInfoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SinglePageInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SinglePageInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SinglePageInfoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown SinglePageInfo mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Book, SinglePageInfo, User []ent.Hook
	}
	inters struct {
		Book, SinglePageInfo, User []ent.Interceptor
	}
)
