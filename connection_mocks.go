package ipop

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/stretchr/testify/mock"
)

// MockConnection is a mock implementation of the Connection interface.
// You can embed this struct in your tests and override methods as needed.
type MockConnection struct {
	mock.Mock
	StringFunc             func() string
	URLFunc                func() string
	MigrationURLFunc       func() string
	MigrationTableNameFunc func() string
	OpenFunc               func() error
	CloseFunc              func() error
	TransactionFunc        func(fn func(tx Connection) error) error
	NewTransactionFunc     func() (Connection, error)
	RollbackFunc           func(fn func(tx Connection)) error
	QFunc                  func() *pop.Query
	TruncateAllFunc        func() error
	BelongsToFunc          func(model interface{}) *pop.Query
	BelongsToAsFunc        func(model interface{}, as string) *pop.Query
	BelongsToThroughFunc   func(bt, thru interface{}) *pop.Query
	ReloadFunc             func(model interface{}) error
	ValidateAndSaveFunc    func(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	SaveFunc               func(model interface{}, excludeColumns ...string) error
	ValidateAndCreateFunc  func(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	CreateFunc             func(model interface{}, excludeColumns ...string) error
	ValidateAndUpdateFunc  func(model interface{}, excludeColumns ...string) (*validate.Errors, error)
	UpdateFunc             func(model interface{}, excludeColumns ...string) error
	DestroyFunc            func(model interface{}) error
	FindFunc               func(model interface{}, id interface{}) error
	FirstFunc              func(model interface{}) error
	LastFunc               func(model interface{}) error
	AllFunc                func(models interface{}) error
	LoadFunc               func(model interface{}, fields ...string) error
	CountFunc              func(model interface{}) (int, error)
	SelectFunc             func(fields ...string) *pop.Query
	PaginateFunc           func(page int, perPage int) *pop.Query
	PaginateFromParamsFunc func(params pop.PaginationParams) *pop.Query
	RawQueryFunc           func(stmt string, args ...interface{}) *pop.Query
	EagerFunc              func(fields ...string) Connection
	WhereFunc              func(stmt string, args ...interface{}) *pop.Query
	OrderFunc              func(stmt string) *pop.Query
	LimitFunc              func(limit int) *pop.Query
	ScopeFunc              func(sf pop.ScopeFunc) *pop.Query
}

func (m *MockConnection) String() string {
	if m.StringFunc != nil {
		return m.StringFunc()
	}
	return "mock-connection"
}
func (m *MockConnection) URL() string {
	if m.URLFunc != nil {
		return m.URLFunc()
	}
	return "mock-url"
}
func (m *MockConnection) MigrationURL() string {
	if m.MigrationURLFunc != nil {
		return m.MigrationURLFunc()
	}
	return "mock-migration-url"
}
func (m *MockConnection) MigrationTableName() string {
	if m.MigrationTableNameFunc != nil {
		return m.MigrationTableNameFunc()
	}
	return "schema_migrations"
}
func (m *MockConnection) Open() error {
	if m.OpenFunc != nil {
		return m.OpenFunc()
	}
	return nil
}
func (m *MockConnection) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}
func (m *MockConnection) Transaction(fn func(tx Connection) error) error {
	if m.TransactionFunc != nil {
		return m.TransactionFunc(fn)
	}
	return nil
}
func (m *MockConnection) NewTransaction() (Connection, error) {
	if m.NewTransactionFunc != nil {
		return m.NewTransactionFunc()
	}
	return m, nil
}
func (m *MockConnection) Rollback(fn func(tx Connection)) error {
	if m.RollbackFunc != nil {
		return m.RollbackFunc(fn)
	}
	return nil
}
func (m *MockConnection) Q() *pop.Query {
	if m.QFunc != nil {
		return m.QFunc()
	}
	return &pop.Query{}
}
func (m *MockConnection) TruncateAll() error {
	if m.TruncateAllFunc != nil {
		return m.TruncateAllFunc()
	}
	return nil
}
func (m *MockConnection) BelongsTo(model interface{}) *pop.Query {
	if m.BelongsToFunc != nil {
		return m.BelongsToFunc(model)
	}
	return &pop.Query{}
}
func (m *MockConnection) BelongsToAs(model interface{}, as string) *pop.Query {
	if m.BelongsToAsFunc != nil {
		return m.BelongsToAsFunc(model, as)
	}
	return &pop.Query{}
}
func (m *MockConnection) BelongsToThrough(bt, thru interface{}) *pop.Query {
	if m.BelongsToThroughFunc != nil {
		return m.BelongsToThroughFunc(bt, thru)
	}
	return &pop.Query{}
}
func (m *MockConnection) Reload(model interface{}) error {
	if m.ReloadFunc != nil {
		return m.ReloadFunc(model)
	}
	return nil
}
func (m *MockConnection) ValidateAndSave(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	if m.ValidateAndSaveFunc != nil {
		return m.ValidateAndSaveFunc(model, excludeColumns...)
	}
	return nil, nil
}
func (m *MockConnection) Save(model interface{}, excludeColumns ...string) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(model, excludeColumns...)
	}
	return nil
}
func (m *MockConnection) ValidateAndCreate(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	if m.ValidateAndCreateFunc != nil {
		return m.ValidateAndCreateFunc(model, excludeColumns...)
	}
	return nil, nil
}
func (m *MockConnection) Create(model interface{}, excludeColumns ...string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(model, excludeColumns...)
	}
	return nil
}
func (m *MockConnection) ValidateAndUpdate(model interface{}, excludeColumns ...string) (*validate.Errors, error) {
	if m.ValidateAndUpdateFunc != nil {
		return m.ValidateAndUpdateFunc(model, excludeColumns...)
	}
	return nil, nil
}
func (m *MockConnection) Update(model interface{}, excludeColumns ...string) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(model, excludeColumns...)
	}
	return nil
}
func (m *MockConnection) Destroy(model interface{}) error {
	if m.DestroyFunc != nil {
		return m.DestroyFunc(model)
	}
	return nil
}
func (m *MockConnection) Find(model interface{}, id interface{}) error {
	if m.FindFunc != nil {
		return m.FindFunc(model, id)
	}
	return nil
}
func (m *MockConnection) First(model interface{}) error {
	if m.FirstFunc != nil {
		return m.FirstFunc(model)
	}
	return nil
}
func (m *MockConnection) Last(model interface{}) error {
	if m.LastFunc != nil {
		return m.LastFunc(model)
	}
	return nil
}
func (m *MockConnection) All(models interface{}) error {
	if m.AllFunc != nil {
		return m.AllFunc(models)
	}
	return nil
}
func (m *MockConnection) Load(model interface{}, fields ...string) error {
	if m.LoadFunc != nil {
		return m.LoadFunc(model, fields...)
	}
	return nil
}
func (m *MockConnection) Count(model interface{}) (int, error) {
	if m.CountFunc != nil {
		return m.CountFunc(model)
	}
	return 0, nil
}
func (m *MockConnection) Select(fields ...string) *pop.Query {
	if m.SelectFunc != nil {
		return m.SelectFunc(fields...)
	}
	return &pop.Query{}
}
func (m *MockConnection) Paginate(page int, perPage int) *pop.Query {
	if m.PaginateFunc != nil {
		return m.PaginateFunc(page, perPage)
	}
	return &pop.Query{}
}
func (m *MockConnection) PaginateFromParams(params pop.PaginationParams) *pop.Query {
	if m.PaginateFromParamsFunc != nil {
		return m.PaginateFromParamsFunc(params)
	}
	return &pop.Query{}
}
func (m *MockConnection) RawQuery(stmt string, args ...interface{}) *pop.Query {
	if m.RawQueryFunc != nil {
		return m.RawQueryFunc(stmt, args...)
	}
	return &pop.Query{}
}
func (m *MockConnection) Eager(fields ...string) Connection {
	if m.EagerFunc != nil {
		return m.EagerFunc(fields...)
	}
	return m
}
func (m *MockConnection) Where(stmt string, args ...interface{}) *pop.Query {
	if m.WhereFunc != nil {
		return m.WhereFunc(stmt, args...)
	}
	return &pop.Query{}
}
func (m *MockConnection) Order(stmt string) *pop.Query {
	if m.OrderFunc != nil {
		return m.OrderFunc(stmt)
	}
	return &pop.Query{}
}
func (m *MockConnection) Limit(limit int) *pop.Query {
	if m.LimitFunc != nil {
		return m.LimitFunc(limit)
	}
	return &pop.Query{}
}
func (m *MockConnection) Scope(sf pop.ScopeFunc) *pop.Query {
	if m.ScopeFunc != nil {
		return m.ScopeFunc(sf)
	}
	return &pop.Query{}
}
