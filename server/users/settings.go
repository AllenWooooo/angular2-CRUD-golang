package users

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

func Init(mux *http.ServeMux, db *mgo.Database) {
	c := db.C("users")

	index := mgo.Index{
		Key:        []string{"name", "balance"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	NewUserController(mux, NewUserService(c))
}
