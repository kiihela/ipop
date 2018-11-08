package ipop

import (
	"github.com/gobuffalo/pop"
)

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
	// when the inner function returns, regardless. This can be useful for tests, etc...
	Rollback(fn func(tx Connection)) error
	// Q creates a new "empty" query for the current connection.
	Q() Query
	// TruncateAll truncates all data from the datasource
	TruncateAll() error
}

type ConnectionAdapter struct {
	conn *pop.Connection
}

func NewConnectionAdapter(c *pop.Connection) *ConnectionAdapter {
	return &ConnectionAdapter{conn: c}
}

func (c *ConnectionAdapter) String() string {
	return c.String()
}

func (c *ConnectionAdapter) URL() string {
	return c.URL()
}

func (c *ConnectionAdapter) MigrationURL() string {
	return c.MigrationURL()
}

func (c *ConnectionAdapter) MigrationTableName() string {
	return c.MigrationTableName()
}

func (c *ConnectionAdapter) Open() error {
	return c.Open()
}

func (c *ConnectionAdapter) Close() error {
	return c.Close()
}

func (c *ConnectionAdapter) Transaction(fn func(tx Connection) error) error {
	return c.Transaction(fn)
}

func (c *ConnectionAdapter) NewTransaction() (Connection, error) {
	conn, err := c.NewTransaction()

	return Connection(conn), err
}

func (c *ConnectionAdapter) Rollback(fn func(tx Connection)) error {
	return c.Rollback(fn)
}

func (c *ConnectionAdapter) Q() Query {
	q := c.Q()

	return Query(q)
}

func (c *ConnectionAdapter) TruncateAll() error {
	return c.TruncateAll()
}
