package ipop

import (
	"github.com/gobuffalo/pop"
)

// Query ...
type Query interface {
	// BelongsTo adds a "where" clause based on the "ID" of the
	// "model" passed into it.
	BelongsTo(model interface{}) Query
	// BelongsToAs adds a "where" clause based on the "ID" of the
	// "model" passed into it, using an alias.
	BelongsToAs(model interface{}, as string) Query
	// BelongsToThrough adds a "where" clause that connects the "bt" model
	// through the associated "thru" model.
	BelongsToThrough(bt, thru interface{}) Query

	// Exec runs the given query.
	Exec() error
	// ExecWithCount runs the given query, and returns the amount of
	// affected rows.
	ExecWithCount() (int, error)

	// Find the first record of the model in the database with a particular id.
	//
	//	q.Find(&User{}, 1)
	Find(model interface{}, id interface{}) error
	// First record of the model in the database that matches the query.
	//
	//	q.Where("name = ?", "mark").First(&User{})
	First(model interface{}) error
	// Last record of the model in the database that matches the query.
	//
	//	q.Where("name = ?", "mark").Last(&User{})
	Last(model interface{}) error
	// All retrieves all of the records in the database that match the query.
	//
	//	q.Where("name = ?", "mark").All(&[]User{})
	All(models interface{}) error
	// Exists returns true/false if a record exists in the database that matches
	// the query.
	//
	// 	q.Where("name = ?", "mark").Exists(&User{})
	Exists(model interface{}) (bool, error)
	// Count the number of records in the database.
	//
	//	q.Where("name = ?", "mark").Count(&User{})
	Count(model interface{}) (int, error)
	// CountByField counts the number of records in the database, for a given field.
	//
	//	q.Where("sex = ?", "f").Count(&User{}, "name")
	CountByField(model interface{}, field string) (int, error)
	// Select allows to query only fields passed as parameter.
	// c.Select("field1", "field2").All(&model)
	// => SELECT field1, field2 FROM models
	Select(fields ...string) Query

	// Paginate records returned from the database.
	//
	//	q = q.Paginate(2, 15)
	//	q.All(&[]User{})
	//	q.Paginator
	Paginate(page int, perPage int) Query
	// PaginateFromParams paginates records returned from the database.
	//
	//	q = q.PaginateFromParams(req.URL.Query())
	//	q.All(&[]User{})
	//	q.Paginator
	PaginateFromParams(params pop.PaginationParams) Query

	// Clone will fill targetQ query with the connection used in q, if
	// targetQ is not empty, Clone will override all the fields.
	Clone(targetQ Query)
	// RawQuery will override the query building feature of Pop and will use
	// whatever query you want to execute against the `Connection`. You can continue
	// to use the `?` argument syntax.
	//
	//	q.RawQuery("select * from foo where id = ?", 1)
	RawQuery(stmt string, args ...interface{}) Query
	// Eager will enable load associations of the model.
	// by defaults loads all the associations on the model,
	// but can take a variadic list of associations to load.
	//
	// 	q.Eager().Find(model, 1) // will load all associations for model.
	// 	q.Eager("Books").Find(model, 1) // will load only Book association for model.
	Eager(fields ...string) Query
	// Where will append a where clause to the query. You may use `?` in place of
	// arguments.
	//
	// 	q.Where("id = ?", 1)
	// 	q.Where("id in (?)", 1, 2, 3)
	Where(stmt string, args ...interface{}) Query
	// Order will append an order clause to the query.
	//
	// 	q.Order("name desc")
	Order(stmt string) Query
	// Limit will add a limit clause to the query.
	Limit(limit int) Query
	// ToSQL will generate SQL and the appropriate arguments for that SQL
	// from the `Model` passed in.
	ToSQL(model *pop.Model, addColumns ...string) (string, []interface{})

	// GroupBy will append a GROUP BY clause to the query
	GroupBy(field string, fields ...string) Query

	// Having will append a HAVING clause to the query
	Having(condition string, args ...interface{}) Query

	// Join will append a JOIN clause to the query
	Join(table string, on string, args ...interface{}) Query
	// LeftJoin will append a LEFT JOIN clause to the query
	LeftJoin(table string, on string, args ...interface{}) Query
	// RightJoin will append a RIGHT JOIN clause to the query
	RightJoin(table string, on string, args ...interface{}) Query
	// LeftOuterJoin will append a LEFT OUTER JOIN clause to the query
	LeftOuterJoin(table string, on string, args ...interface{}) Query
	// RightOuterJoin will append a RIGHT OUTER JOIN clause to the query
	RightOuterJoin(table string, on string, args ...interface{}) Query
	// LeftInnerJoin will append a LEFT INNER JOIN clause to the query
	LeftInnerJoin(table string, on string, args ...interface{}) Query
	// RightInnerJoin will append a RIGHT INNER JOIN clause to the query
	RightInnerJoin(table string, on string, args ...interface{}) Query

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
	Scope(sf pop.ScopeFunc) Query
}
