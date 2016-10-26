package users

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type Error struct {
	Error string `json:"error"`
}

type UserController struct {
	userService UserService
}

func NewUserController(mux *http.ServeMux, s UserService) *UserController {
	ctrl := UserController{s}

	mux.Handle("/users", ctrl)
	mux.Handle("/users/", ctrl)

	return &ctrl
}

func (u UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := http.StatusMethodNotAllowed
	var data interface{}

	defer func(c int) {
		log.Println(r.URL, "-", r.Method, "-", code, r.RemoteAddr)
	}(code)

	if r.URL.Path == "/users" {
		switch r.Method {
		case "GET":
			code, data = u.List(w, r)
		case "POST":
			code, data = u.Add(w, r)
		default:
			return
		}
	} else {
		switch r.Method {
		case "GET":
			code, data = u.Get(w, r)
		case "PUT":
			code, data = u.Update(w, r)
		case "DELETE":
			code, data = u.Delete(w, r)
		default:
			return
		}
	}

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Println("Failed to write data:", err)
		code = http.StatusInternalServerError
	}
}

func (u UserController) List(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	return http.StatusOK, u.userService.List()
}

func (u UserController) Add(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	if r.Body == nil {
		return http.StatusBadRequest, Error{"no payload"}
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return http.StatusBadRequest, Error{"can't parse json payload"}
	}

	if user.Name == "" {
		return 422, Error{"please provide 'name'"}
	}

	u.userService.Add(&user)
	return http.StatusCreated, u.userService.List()
}

func (u UserController) Get(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	if !bson.IsObjectIdHex(id) {
		return http.StatusBadRequest, Error{"id should be bson.ObjectId"}
	}

	objectID := bson.ObjectIdHex(id)
	user, err := u.userService.Get(objectID)

	if err != nil {
		return http.StatusNotFound, Error{err.Error()}
	}

	return http.StatusOK, user
}

func (u UserController) Update(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	if r.Body == nil {
		return http.StatusBadRequest, Error{"no payload"}
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return http.StatusBadRequest, Error{"can't parse json payload"}
	}

	if user.Name == "" {
		return 422, Error{"please provide 'name'"}
	}

	err = u.userService.Update(&user)

	if err != nil {
		return http.StatusNotFound, Error{err.Error()}
	}

	return http.StatusOK, u.userService.List()
}

func (u UserController) Delete(w http.ResponseWriter, r *http.Request) (int, interface{}) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	if !bson.IsObjectIdHex(id) {
		return http.StatusBadRequest, Error{"id should be bson.ObjectId"}
	}

	objectID := bson.ObjectIdHex(id)
	err := u.userService.Delete(objectID)

	if err != nil {
		return http.StatusNotFound, Error{err.Error()}
	}

	return http.StatusOK, u.userService.List()
}
