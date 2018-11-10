package ipop

import (
	"log"
	"testing"

	"github.com/dnnrly/ipop/testdata/models"
	"github.com/gobuffalo/pop"
	"github.com/stretchr/testify/assert"
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

	migrator, err := pop.NewFileMigrator("testdata/migrations", connection)
	if err != nil {
		log.Panic(err)
	}

	err = migrator.Reset()
	if err != nil {
		log.Panic(err)
	}

	db = NewConnectionAdapter(connection)
}

func TestConnectionAdapter_SaveAndFindAndUpdate(t *testing.T) {
	user := models.User{
		Name: "Bob",
	}

	err := db.Save(&user)
	assert.NoError(t, err)
	origUpdated := user.UpdatedAt

	found := models.User{}
	err = db.Find(&found, user.ID)
	assert.NoError(t, err)
	assert.False(t, found.CreatedAt.IsZero())

	updated := models.User{
		ID:   user.ID,
		Name: "Another name",
	}

	err = db.Update(&updated)
	assert.NoError(t, err)
	assert.False(t, found.UpdatedAt.IsZero())
	assert.NotEqual(t, origUpdated.String(), found.UpdatedAt.String())
}

func TestConnectionAdapter_CreateAndFirst(t *testing.T) {
	user := models.User{
		Name: "Alice",
	}

	err := db.Create(&user)
	assert.NoError(t, err)

	found := models.User{}
	err = db.First(&found)
	assert.NoError(t, err)
}

func TestConnectionAdapter_ErrorCallback(t *testing.T) {
	called := false

	err := db.Transaction(func(tx Connection) error {
		called = true
		return nil
	})

	assert.NoError(t, err)
	assert.True(t, called)
}

func TestConnectionAdapter_Callback(t *testing.T) {
	called := false

	err := db.Rollback(func(tx Connection) {
		called = true
	})

	assert.NoError(t, err)
	assert.True(t, called)
}
