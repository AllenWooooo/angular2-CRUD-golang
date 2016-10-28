package controller

import (
	"net/http"

	"../users"
	mgo "gopkg.in/mgo.v2"
)

func Init(mux *http.ServeMux, session *mgo.Session) {
	db := session.DB("crud_demo")

	users.Init(mux, db)
}
