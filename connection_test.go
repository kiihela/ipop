package ipop

import (
	"log"

	"github.com/gobuffalo/pop"
)

var (
	Debug = false
	Color = false

	db *pop.Connection
)

func init() {
	pop.Debug = Debug
	pop.Color = Color
	connection, err := pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}

	connection.
		db = connection
}

func ExampleSimpleFindSingle() {

}
