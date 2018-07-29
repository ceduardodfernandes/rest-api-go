package dao

import (
	"log"
	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WidgetsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	USERS = "users"
	WIDGETS = "widgets"
)

func (m *WidgetsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *WidgetsDAO) FindAllUsers() ([]User, error) {
	var users []User
	err := db.C(USERS).Find(bson.M{}).All(&users)
	return users, err
}

func (m *WidgetsDAO) FindUserById(id string) (User, error) {
	var user User
	err := db.C(USERS).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *WidgetsDAO) FindAllWidgets() ([]Widget, error) {
	var widgets []Widget
	err := db.C(WIDGETS).Find(bson.M{}).All(&widgets)
	return widgets, err
}

func (m *WidgetsDAO) FindWidgetById(id string) (Widget, error) {
	var widget Widget
	err := db.C(WIDGETS).FindId(bson.ObjectIdHex(id)).One(&widget)
	return widget, err
}

func (m *WidgetsDAO) InsertWidget(widget Widget) error {
	err := db.C(WIDGETS).Insert(&widget)
	return err
}

func (m *WidgetsDAO) UpdateWidget(widget Widget) error {
	err := db.C(WIDGETS).UpdateId(widget.ID, &widget)
	return err
}

func (m *WidgetsDAO) RemoveWidget(widget Widget) error {
	err := db.C(WIDGETS).Remove(&widget)
	return err
}
