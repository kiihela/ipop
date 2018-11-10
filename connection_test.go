package ipop

import (
	"log"
	"testing"

	"github.com/gobuffalo/pop"
	"github.com/stretchr/testify/assert"

	"github.com/dnnrly/ipop/testdata/models"
)

var (
	Debug = false
	Color = false

	db Connection
)

func init() {
	pop.Debug = Debug
	pop.Color = Color
	connection, err := pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}

	connection.MigrateReset("testdata/migrations")
	
	db = NewConnectionAdapter(connection)
}

func TestConnectionAdapter_SaveAndFind(t *testing.T) {
	user := models.User{
		Name: "Bob",
	}

	err := db.Save(&user)
	assert.NoError(t, err)

	found := models.User{}
	err = db.Find(&found, user.ID)
	assert.NoError(t, err)
	assert.False(t, found.CreatedAt.IsZero())
}

func TestConnectionAdapter_Callback(t *testing.T) {
	called := false

	err := db.Transaction(func (tx Connection) error {
		called = true
		return nil
	})

	assert.NoError(t, err)
	assert.True(t, called)
}
