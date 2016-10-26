package main

import (
	"log"
	"net/http"

	"./controller"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../static/")))

	controller.Init(mux, session)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
