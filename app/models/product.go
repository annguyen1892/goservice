package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Product Represent a product
// the properties in mongodb document
type Product struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID int32         `bson:"product_id"`
	UserID    int32         `bson:"user_id"`
	CreatedAt string        `bson:"created_at"`
	Microtime float64       `bson:"microtime"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
