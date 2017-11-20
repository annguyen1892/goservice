package handler

import (
	"github.com/annguyen1892/goservice/app/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

type Handler struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "recently_viewed"
)

func (m *Handler) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *Handler) FindAll() ([]Products, error) {
	// var products []Products
	// err := db.C(COLLECTION).Find(b.son.M{}).All(&products)
	// return products, err
}
