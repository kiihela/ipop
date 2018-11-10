package ipop

import (
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type popCBErr func(tx *pop.Connection) error
type popCB func(tx *pop.Connection)
type cbErr func(tx Connection) error
type cb func(tx Connection)

func cbConvert(cfn cb) popCB {
	return func(tx *pop.Connection) {
		ctx := NewConnectionAdapter(tx)
		cfn(ctx)
	}
}

func cbConvertErr(cfn cbErr) popCBErr {
	return func(tx *pop.Connection) error {
		ctx := NewConnectionAdapter(tx)
		return cfn(ctx)
	}
}

// Connection represents a Buffalo pop connection to a database
type Connection interface {
	String() string
	// URL returns the datasource connection string
	URL() string
	// MigrationURL returns the datasource connection string used for running the migrations
	MigrationURL() string
	// MigrationTableName returns the name of the table to track migrations
	MigrationTableName() string
	// Open creates a new datasource connection
	Open() error
	// Close destroys an active datasource connection
	Close() error
	// Transaction will start a new transaction on the connection. If the inner function
	// returns an error then the transaction will be rolled back, otherwise the transaction
	// will automatically commit at the end.
	Transaction(fn func(tx Connection) error) error
	// NewTransaction starts a new transaction on the connection
	NewTransaction() (Connection, error)
	// Rollback will open a new transaction and automatically rollback that transaction
	// when the inner function returns, regardless. This can be useful for tests, etc.
	Rollback(fn func(tx Connection)) error
	// Q creates a new "empty" query for the current connection.
	Q() *pop.Query
	// TruncateAll truncates all data from the datasource
	TruncateAll() error

	// BelongsTo adds a "where" clause based on the "ID" of the
	// "model" passed into it.
	BelongsTo(model interface{}) *pop.Query
	// BelongsToAs adds a "where" clause based on the "ID" of the
	// "model" passed into it using an alias.
	BelongsToAs(model interface{}, as string) *pop.Query
	// BelongsToThrough adds a "where" clause that connects the "bt" model
	// through the associated "thru" model.
	BelongsToThrough(bt, thru interface{}) *pop.Query

	// Reload fetch fresh data for a given model, using its ID.
	Reload(model interface{}) error
	// ValidateAndSave applies validation rules on the given entry, then save it
	// if the validation succeed, excluding the given columns.
	ValidateAndSave(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	// Save wraps the Create and Update methods. It executes a Create if no ID is provided with the entry;
	// or issues an Update otherwise.
	Save(model interface{}, excludeColumns ...string) error
	// ValidateAndCreate applies validation rules on the given entry, then creates it
	// if the validation succeed, excluding the given columns.
	ValidateAndCreate(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	// Create add a new given entry to the database, excluding the given columns.
	// It updates `created_at` and `updated_at` columns automatically.
	Create(model interface{}, excludeColumns ...string) error
	// ValidateAndUpdate applies validation rules on the given entry, then update it
	// if the validation succeed, excluding the given columns.
	ValidateAndUpdate(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	// Update writes changes from an entry to the database, excluding the given columns.
	// It updates the `updated_at` column automatically.
	Update(model interface{}, excludeColumns ...string) error
	// Destroy deletes a given entry from the database
	Destroy(model interface{}) error

	// Find the first record of the model in the database with a particular id.
	//
	//	c.Find(&User{}, 1)
	Find(model interface{}, id interface{}) error
	// First record of the model in the database that matches the query.
	//
	//	c.First(&User{})
	First(model interface{}) error
	// Last record of the model in the database that matches the query.
	//
	//	c.Last(&User{})
	Last(model interface{}) error
	// All retrieves all of the records in the database that match the query.
	//
	//	c.All(&[]User{})
	All(models interface{}) error
	// Load loads all association or the fields specified in params for
	// an already loaded model.
	//
	// tx.First(&u)
	// tx.Load(&u)
	Load(model interface{}, fields ...string) error
	// Count the number of records in the database.
	//
	//	c.Count(&User{})
	Count(model interface{}) (int, error)
	// Select allows to query only fields passed as parameter.
	// c.conn.Select("field1", "field2").All(&model)
	// => SELECT field1, field2 FROM models
	Select(fields ...string) *pop.Query

	// MigrateUp is deprecated, and will be removed in a future version. Use FileMigrator#Up instead.
	MigrateUp(path string) error
	// MigrateDown is deprecated, and will be removed in a future version. Use FileMigrator#Down instead.
	MigrateDown(path string, step int) error
	// MigrateStatus is deprecated, and will be removed in a future version. Use FileMigrator#Status instead.
	MigrateStatus(path string) error
	// MigrateReset is deprecated, and will be removed in a future version. Use FileMigrator#Reset instead.
	MigrateReset(path string) error

	// Paginate records returned from the database.
	//
	//	return c.conn.Paginate(2, 15)
	//	q.All(&[]User{})
	//	q.Paginator
	Paginate(page int, perPage int) *pop.Query
	// PaginateFromParams paginates records returned from the database.
	//
	//	return c.conn.PaginateFromParams(req.URL.Query())
	//	q.All(&[]User{})
	//	q.Paginator
	PaginateFromParams(params pop.PaginationParams) *pop.Query

	// RawQuery will override the query building feature of Pop and will use
	// whatever query you want to execute against the `Connection`. You can continue
	// to use the `?` argument syntax.
	//
	//	c.RawQuery("select * from foo where id = ?", 1)
	RawQuery(stmt string, args ...interface{}) *pop.Query
	// Eager will enable load associations of the model.
	// by defaults loads all the associations on the model,
	// but can take a variadic list of associations to load.
	//
	// 	c.Eager().Find(model, 1) // will load all associations for model.
	// 	c.Eager("Books").Find(model, 1) // will load only Book association for model.
	Eager(fields ...string) Connection
	// Where will append a where clause to the query. You may use `?` in place of
	// arguments.
	//
	// 	c.Where("id = ?", 1)
	// 	q.Where("id in (?)", 1, 2, 3)
	Where(stmt string, args ...interface{}) *pop.Query
	// Order will append an order clause to the query.
	//
	// 	c.Order("name desc")
	Order(stmt string) *pop.Query
	// Limit will add a limit clause to the query.
	Limit(limit int) *pop.Query

	// Scope the query by using a `ScopeFunc`
	//
	//	func ByName(name string) ScopeFunc {
	//		return func(q Query) Query {
	//			return q.Where("name = ?", name)
	//		}
	//	}
	//
	//	func WithDeleted(q *pop.Query) *pop.Query {
	//		return q.Where("deleted_at is null")
	//	}
	//
	//	c.Scope(ByName("mark)).Scope(WithDeleted).First(&User{})
	Scope(sf pop.ScopeFunc) *pop.Query
}

type ConnectionAdapter struct {
	conn *pop.Connection
}

func NewConnectionAdapter(c *pop.Connection) *ConnectionAdapter {
	return &ConnectionAdapter{conn: c}
}

func (c *ConnectionAdapter) String() string {
	return c.conn.String()
}

// URL returns the datasource connection string
func (c *ConnectionAdapter) URL() string {
	return c.conn.URL()
}

// MigrationURL returns the datasource connection string used for running the migrations
func (c *ConnectionAdapter) MigrationURL() string {
	return c.conn.MigrationURL()
}

// MigrationTableName returns the name of the table to track migrations
func (c *ConnectionAdapter) MigrationTableName() string {
	return c.conn.MigrationTableName()
}

// Open creates a new datasource connection
func (c *ConnectionAdapter) Open() error {
	return c.conn.Open()
}

// Close destroys an active datasource connection
func (c *ConnectionAdapter) Close() error {
	return c.conn.Close()
}

// Transaction will start a new transaction on the connection. If the inner function
// returns an error then the transaction will be rolled back, otherwise the transaction
// will automatically commit at the end.
func (c *ConnectionAdapter) Transaction(fn func(tx Connection) error) error {
	return c.conn.Transaction(cbConvertErr(fn))
}

// NewTransaction starts a new transaction on the connection
func (c *ConnectionAdapter) NewTransaction() (Connection, error) {
	conn, err := c.conn.NewTransaction()
	return NewConnectionAdapter(conn), err
}

// Rollback will open a new transaction and automatically rollback that transaction
// when the inner function returns, regardless. This can be useful for tests, etc.
func (c *ConnectionAdapter) Rollback(fn func(tx Connection)) error {
	return c.conn.Rollback(cbConvert(fn))
}

// Q creates a new "empty" query for the current connection.
func (c *ConnectionAdapter) Q() *pop.Query {
	return c.conn.Q()
}

// TruncateAll truncates all data from the datasource
func (c *ConnectionAdapter) TruncateAll() error {
	return c.conn.TruncateAll()
}

// BelongsTo adds a "where" clause based on the "ID" of the
// "model" passed into it.
func (c *ConnectionAdapter) BelongsTo(model interface{}) *pop.Query {
	return c.conn.BelongsTo(model)
}

// BelongsToAs adds a "where" clause based on the "ID" of the
// "model" passed into it using an alias.
func (c *ConnectionAdapter) BelongsToAs(model interface{}, as string) *pop.Query {
	return c.conn.BelongsToAs(model, as)
}

// BelongsToThrough adds a "where" clause that connects the "bt" model
// through the associated "thru" model.
func (c *ConnectionAdapter) BelongsToThrough(bt interface{}, thru interface{}) *pop.Query {
	return c.conn.BelongsToThrough(bt, thru)
}

// Reload fetch fresh data for a given model, using its ID.
func (c *ConnectionAdapter) Reload(model interface{}) error {
	return c.conn.Reload(model)
}

// ValidateAndSave applies validation rules on the given entry, then save it
// if the validation succeed, excluding the given columns.
func (c *ConnectionAdapter) ValidateAndSave(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	return c.conn.ValidateAndSave(model, excludeColumns...)
}

// Save wraps the Create and Update methods. It executes a Create if no ID is provided with the entry;
// or issues an Update otherwise.
func (c *ConnectionAdapter) Save(model interface{}, excludeColumns ...string) error {
	return c.conn.Save(model, excludeColumns...)
}

// ValidateAndCreate applies validation rules on the given entry, then creates it
// if the validation succeed, excluding the given columns.
func (c *ConnectionAdapter) ValidateAndCreate(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	return c.conn.ValidateAndCreate(model, excludeColumns...)
}

// Create add a new given entry to the database, excluding the given columns.
// It updates `created_at` and `updated_at` columns automatically.
func (c *ConnectionAdapter) Create(model interface{}, excludeColumns ...string) error {
	return c.conn.Create(model, excludeColumns...)
}

// ValidateAndUpdate applies validation rules on the given entry, then update it
// if the validation succeed, excluding the given columns.
func (c *ConnectionAdapter) ValidateAndUpdate(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	return c.conn.ValidateAndUpdate(model, excludeColumns...)
}

// Update writes changes from an entry to the database, excluding the given columns.
// It updates the `updated_at` column automatically.
func (c *ConnectionAdapter) Update(model interface{}, excludeColumns ...string) error {
	return c.conn.Update(model, excludeColumns...)
}

// Destroy deletes a given entry from the database
func (c *ConnectionAdapter) Destroy(model interface{}) error {
	return c.conn.Destroy(model)
}

// Find the first record of the model in the database with a particular id.
//
//	c.Find(&User{}, 1)
func (c *ConnectionAdapter) Find(model interface{}, id interface{}) error {
	return c.conn.Find(model, id)
}

// First record of the model in the database that matches the query.
//
//	c.First(&User{})
func (c *ConnectionAdapter) First(model interface{}) error {
	return c.conn.First(model)
}

// Last record of the model in the database that matches the query.
//
//	c.Last(&User{})
func (c *ConnectionAdapter) Last(model interface{}) error {
	return c.conn.Last(model)
}

// All retrieves all of the records in the database that match the query.
//
//	c.All(&[]User{})
func (c *ConnectionAdapter) All(models interface{}) error {
	return c.conn.All(models)
}

// Load loads all association or the fields specified in params for
// an already loaded model.
//
// tx.First(&u)
// tx.Load(&u)
func (c *ConnectionAdapter) Load(model interface{}, fields ...string) error {
	return c.conn.Load(model, fields...)
}

// Count the number of records in the database.
//
//	c.Count(&User{})
func (c *ConnectionAdapter) Count(model interface{}) (int, error) {
	return c.conn.Count(model)
}

// Select allows to query only fields passed as parameter.
// c.conn.Select("field1", "field2").All(&model)
// => SELECT field1, field2 FROM models
func (c *ConnectionAdapter) Select(fields ...string) *pop.Query {
	return c.conn.Select(fields...)
}

// MigrateUp is deprecated, and will be removed in a future version. Use FileMigrator#Up instead.
func (c *ConnectionAdapter) MigrateUp(path string) error {
	return c.conn.MigrateUp(path)
}

// MigrateDown is deprecated, and will be removed in a future version. Use FileMigrator#Down instead.
func (c *ConnectionAdapter) MigrateDown(path string, step int) error {
	return c.conn.MigrateDown(path, step)
}

// MigrateStatus is deprecated, and will be removed in a future version. Use FileMigrator#Status instead.
func (c *ConnectionAdapter) MigrateStatus(path string) error {
	return c.conn.MigrateStatus(path)
}

// MigrateReset is deprecated, and will be removed in a future version. Use FileMigrator#Reset instead.
func (c *ConnectionAdapter) MigrateReset(path string) error {
	return c.conn.MigrateReset(path)
}

// Paginate records returned from the database.
//
//	return c.conn.Paginate(2, 15)
//	q.All(&[]User{})
//	q.Paginator
func (c *ConnectionAdapter) Paginate(page int, perPage int) *pop.Query {
	return c.conn.Paginate(page, perPage)
}

// PaginateFromParams paginates records returned from the database.
//
//	return c.conn.PaginateFromParams(req.URL.Query())
//	q.All(&[]User{})
//	q.Paginator
func (c *ConnectionAdapter) PaginateFromParams(params pop.PaginationParams) *pop.Query {
	return c.conn.PaginateFromParams(params)
}

// RawQuery will override the query building feature of Pop and will use
// whatever query you want to execute against the `Connection`. You can continue
// to use the `?` argument syntax.
//
//	c.RawQuery("select * from foo where id = ?", 1)]
func (c *ConnectionAdapter) RawQuery(stmt string, args ...interface{}) *pop.Query {
	return c.conn.RawQuery(stmt, args...)
}

// Eager will enable load associations of the model.
// by defaults loads all the associations on the model,
// but can take a variadic list of associations to load.
//
// 	c.Eager().Find(model, 1) // will load all associations for model.
// 	c.Eager("Books").Find(model, 1) // will load only Book association for model.
func (c *ConnectionAdapter) Eager(fields ...string) Connection {
	popConn := c.conn.Eager(fields...)
	return NewConnectionAdapter(popConn)
}

// Where will append a where clause to the query. You may use `?` in place of
// arguments.
//
// 	c.Where("id = ?", 1)
// 	q.Where("id in (?)", 1, 2, 3)
func (c *ConnectionAdapter) Where(stmt string, args ...interface{}) *pop.Query {
	return c.conn.Where(stmt, args...)
}

// Order will append an order clause to the query.
//
// 	c.Order("name desc")
func (c *ConnectionAdapter) Order(stmt string) *pop.Query {
	return c.conn.Order(stmt)
}

// Limit will add a limit clause to the query.
func (c *ConnectionAdapter) Limit(limit int) *pop.Query {
	return c.conn.Limit(limit)
}

// Scope the query by using a `ScopeFunc`
//
//	func ByName(name string) ScopeFunc {
//		return func(q Query) Query {
//			return q.Where("name = ?", name)
//		}
//	}
//
//	func WithDeleted(q *pop.Query) *pop.Query {
//		return q.Where("deleted_at is null")
//	}
//
//	c.Scope(ByName("mark)).Scope(WithDeleted).First(&User{})
func (c *ConnectionAdapter) Scope(sf pop.ScopeFunc) *pop.Query {
	return c.conn.Scope(sf)
}
