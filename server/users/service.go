package users

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name" bson:"name"`
	Balance float64       `json:"balance" bson:"balance"`
}

type UserService interface {
	List() []*User
	Add(*User)
	Get(bson.ObjectId) (*User, error)
	Update(*User) error
	Delete(bson.ObjectId) error
}

type userService struct {
	c *mgo.Collection
}

func NewUserService(c *mgo.Collection) UserService {
	return &userService{c}
}

func (s userService) List() []*User {
	users := []*User{}
	err := s.c.Find(nil).All(&users)

	if err != nil {
		panic(err)
	}

	return users
}

func (s userService) Add(user *User) {
	user.ID = bson.NewObjectId()

	s.c.Insert(user)
}

func (s userService) Get(id bson.ObjectId) (*User, error) {
	var user *User

	if err := s.c.FindId(id).One(&user); err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s userService) Update(user *User) error {
	if err := s.c.UpdateId(user.ID, user); err != nil {
		return errors.New("User not found")
	}

	return nil
}

func (s userService) Delete(id bson.ObjectId) error {
	if err := s.c.RemoveId(id); err != nil {
		return errors.New("User not found")
	}

	return nil
}
