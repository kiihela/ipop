package ipop

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/mock"
)

type MockQuery struct {
	mock.Mock
}

func (m *MockQuery) BelongsTo(model interface{}) Query {
	args := m.Called(model)
	return args.Get(0).(Query)
}
func (m *MockQuery) BelongsToAs(model interface{}, as string) Query {
	args := m.Called(model, as)
	return args.Get(0).(Query)
}
func (m *MockQuery) BelongsToThrough(bt, thru interface{}) Query {
	args := m.Called(bt, thru)
	return args.Get(0).(Query)
}
func (m *MockQuery) Exec() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockQuery) ExecWithCount() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}
func (m *MockQuery) Find(model interface{}, id interface{}) error {
	args := m.Called(model, id)
	return args.Error(0)
}
func (m *MockQuery) First(model interface{}) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockQuery) Last(model interface{}) error {
	args := m.Called(model)
	return args.Error(0)
}
func (m *MockQuery) All(models interface{}) error {
	args := m.Called(models)
	return args.Error(0)
}
func (m *MockQuery) Exists(model interface{}) (bool, error) {
	args := m.Called(model)
	return args.Bool(0), args.Error(1)
}
func (m *MockQuery) Count(model interface{}) (int, error) {
	args := m.Called(model)
	return args.Int(0), args.Error(1)
}
func (m *MockQuery) CountByField(model interface{}, field string) (int, error) {
	args := m.Called(model, field)
	return args.Int(0), args.Error(1)
}
func (m *MockQuery) Select(fields ...string) Query {
	args := m.Called(fields)
	return args.Get(0).(Query)
}
func (m *MockQuery) Paginate(page int, perPage int) Query {
	args := m.Called(page, perPage)
	return args.Get(0).(Query)
}
func (m *MockQuery) PaginateFromParams(params pop.PaginationParams) Query {
	args := m.Called(params)
	return args.Get(0).(Query)
}
func (m *MockQuery) Clone(targetQ Query) {
	m.Called(targetQ)
}
func (m *MockQuery) RawQuery(stmt string, argsIn ...interface{}) Query {
	args := m.Called(stmt, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) Eager(fields ...string) Query {
	args := m.Called(fields)
	return args.Get(0).(Query)
}
func (m *MockQuery) Where(stmt string, argsIn ...interface{}) Query {
	args := m.Called(stmt, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) Order(stmt string) Query {
	args := m.Called(stmt)
	return args.Get(0).(Query)
}
func (m *MockQuery) Limit(limit int) Query {
	args := m.Called(limit)
	return args.Get(0).(Query)
}
func (m *MockQuery) ToSQL(model *pop.Model, addColumns ...string) (string, []interface{}) {
	args := m.Called(model, addColumns)
	return args.String(0), args.Get(1).([]interface{})
}
func (m *MockQuery) GroupBy(field string, fields ...string) Query {
	args := m.Called(field, fields)
	return args.Get(0).(Query)
}
func (m *MockQuery) Having(condition string, argsIn ...interface{}) Query {
	args := m.Called(condition, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) Join(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) LeftJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) RightJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) LeftOuterJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) RightOuterJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) LeftInnerJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) RightInnerJoin(table string, on string, argsIn ...interface{}) Query {
	args := m.Called(table, on, argsIn)
	return args.Get(0).(Query)
}
func (m *MockQuery) Scope(sf pop.ScopeFunc) Query {
	args := m.Called(sf)
	return args.Get(0).(Query)
}
