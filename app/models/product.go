package models

// Product Represent a product
// the properties in mongodb document
type Product struct {
	ProductID int32  `bson:"product_id" json:"product_id"`
	UserID    int32  `bson:"user_id"    json:"user_id"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	Microtime int64  `bson:"microtime"  json:"microtime"`
}

type Paging struct {
	Total       int `bson:"total" json:"total"`
	PerPage     int `bson:"per_page" json:"per_page"`
	CurrentPage int `bson:"current_page" json:"current_page"`
	LastPage    int `bson:"last_page" json:"last_page"`
}

type Response struct {
	Paging Paging    `bson:"paging" json:"paging"`
	Data   []Product `bson:"data" json:"data"`
}

type Create struct {
	Result string `bson:"result" json:"result"`
}
