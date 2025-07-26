package ipop

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/pop/v6"
	"log"
	"net/url"
	"testing"

	"github.com/kiihela/ipop/testdata/models"
	"github.com/stretchr/testify/assert"
)

var (
	Debug = false
	Color = false

	popConn *pop.Connection
	db      Connection
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

	popConn = connection
	db = NewConnectionAdapter(connection)
}

func ExampleNewConnectionAdapter() {
	popConnection, _ := pop.Connect("test")
	adapted := NewConnectionAdapter(popConnection)
	adapted.String() // Use it as you would *pop.Connection
}

func TestConnectionAdapter_Strings(t *testing.T) {
	assert.Equal(t, popConn.String(), db.String())
	assert.Equal(t, popConn.URL(), db.URL())
}

func TestConnectionAdapter_SaveAndUpdating(t *testing.T) {
	user := models.User{
		Name: "Bob",
	}

	assert.NoError(t, db.Save(&user))
	origUpdated := user.UpdatedAt

	found := models.User{}
	assert.NoError(t, db.Find(&found, user.ID))
	assert.False(t, found.CreatedAt.IsZero())

	updated := models.User{
		ID:   user.ID,
		Name: "Another name",
	}

	assert.NoError(t, db.Save(&updated))
	assert.False(t, found.UpdatedAt.IsZero())
	assert.NotEqual(t, origUpdated.String(), found.UpdatedAt.String())

	assert.NoError(t, db.Find(&found, user.ID))
	assert.Equal(t, "Another name", found.Name)

	updated.Name = "Yet another name"
	assert.NoError(t, db.Update(&updated))

	assert.NoError(t, db.Find(&found, user.ID))
	assert.Equal(t, "Yet another name", found.Name)

	assert.NoError(t, db.Destroy(&found))
}

func TestConnectionAdapter_CreateAndQueries(t *testing.T) {
	assert.NoError(t, db.TruncateAll())

	for i := 0; i < 100; i++ {
		u := models.User{
			Name: fmt.Sprintf("User #%d", i+1),
		}
		assert.NoError(t, db.Create(&u), "Could not create user %d", i)
	}

	var all []models.User
	assert.NoError(t, db.All(&all))
	assert.Equal(t, 100, len(all))

	count, err := db.Count(&models.User{})
	assert.NoError(t, err)
	assert.Equal(t, 100, count)

	var first models.User
	assert.NoError(t, db.First(&first))
	assert.Equal(t, all[0].ID, first.ID)

	var last models.User
	assert.NoError(t, db.Last(&last))
	assert.Equal(t, all[99].ID, last.ID)

	var page2 []models.User
	q1 := db.Paginate(2, 2)
	assert.NoError(t, q1.All(&page2))
	assert.Equal(t, "User #3", page2[0].Name)
	assert.Equal(t, "User #4", page2[1].Name)

	var page3 []models.User
	q2 := db.PaginateFromParams(url.Values{"page": {"3"}, "per_page": {"2"}})
	assert.NoError(t, q2.All(&page3))
	assert.Equal(t, "User #5", page3[0].Name)
	assert.Equal(t, "User #6", page3[1].Name)

	var where []models.User
	q3 := db.Where("Name = ?", "User #10")
	assert.NoError(t, q3.All(&where))
	assert.Equal(t, 1, len(where))

	var order []models.User
	q4 := db.Order("name desc")
	assert.NoError(t, q4.All(&order))
	assert.Equal(t, "User #99", order[0].Name)

	assert.NoError(t, db.TruncateAll())

	count, err = db.Count(&models.User{})
	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}

func TestConnectionAdapter_TransactionWorks(t *testing.T) {
	called := false

	tx, err := db.NewTransaction()
	assert.NoError(t, err)

	err = tx.Transaction(func(tx Connection) error {
		called = true
		return nil
	})

	assert.NoError(t, err)
	assert.True(t, called)
}

func TestConnectionAdapter_Verification(t *testing.T) {
	user := models.User{
		Name: "George",
	}

	verrs, _ := db.ValidateAndCreate(&user)
	assert.Equal(t, 0, len(verrs.Errors))
	origUpdated := user.UpdatedAt

	found := models.User{}
	assert.NoError(t, db.Find(&found, user.ID))
	assert.False(t, found.CreatedAt.IsZero())

	updated := models.User{
		ID:   user.ID,
		Name: "Not that person",
	}

	verrs, _ = db.ValidateAndSave(&updated)
	assert.Equal(t, 0, len(verrs.Errors))
	assert.False(t, found.UpdatedAt.IsZero())
	assert.NotEqual(t, origUpdated.String(), found.UpdatedAt.String())

	assert.NoError(t, db.Find(&found, user.ID))
	assert.Equal(t, "Not that person", found.Name)

	updated.Name = "Someone else"
	verrs, _ = db.ValidateAndUpdate(&updated)
	assert.Equal(t, 0, len(verrs.Errors))

	assert.NoError(t, db.Find(&found, user.ID))
	assert.Equal(t, "Someone else", found.Name)

	assert.NoError(t, db.Destroy(&found))
}

func TestConnectionAdapter_TransactionFails(t *testing.T) {
	called := false

	tx, err := db.NewTransaction()
	assert.NoError(t, err)

	err = tx.Transaction(func(tx Connection) error {
		called = true
		return errors.New("ooops")
	})

	assert.Error(t, err)
	assert.True(t, called)
}

func TestConnectionAdapter_Callback(t *testing.T) {
	called := false

	tx, err := db.NewTransaction()
	assert.NoError(t, err)

	err = tx.Rollback(func(tx Connection) {
		called = true
	})

	assert.NoError(t, err)
	assert.True(t, called)
}

func TestConnectionAdapter_Connections(t *testing.T) {
	err := db.Open()
	assert.NoError(t, err)

	err = db.Close()
	assert.NoError(t, err)
}
