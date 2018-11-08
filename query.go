package ipop

import (
	"github.com/gobuffalo/pop"
)

// Query ...
type Query interface {
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
}
