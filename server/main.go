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
	mux.Handle("/assets/", http.FileServer(http.Dir("../static/")))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/index.html")
	})

	controller.Init(mux, session)
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
