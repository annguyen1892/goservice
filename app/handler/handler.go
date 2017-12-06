package handler

import (
	// "fmt"
	"fmt"
	// "github.com/getsentry/raven-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math"
	"strconv"
	"strings"
	. "talaria-recentlyviewed-go/app/models"
	"time"
)

// Handler ...
type Handler struct {
	Addrs          string
	Timeout        int
	Database       string
	Username       string
	Password       string
	Replicasetname string
	Readpreference string
}

var db *mgo.Database
var mode_read mgo.Mode

// COLLECTION
const (
	COLLECTION = "recently_viewed"
)

// Connect mongodb
func (m *Handler) Connect() {
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:          []string{m.Addrs},
		Database:       m.Database,
		Username:       m.Username,
		Password:       m.Password,
		Timeout:        5 * time.Second,
		ReplicaSetName: m.Replicasetname,
	})
	if err != nil {
		// raven.CaptureErrorAndWait(err, nil)
		log.Fatal(err)
	}

	switch m.Readpreference {
	case "Secondary":
		mode_read = mgo.Secondary
	case "PrimaryPreferred":
		mode_read = mgo.PrimaryPreferred
	case "Primary":
		mode_read = mgo.Primary
	case "Nearest":
		mode_read = mgo.Nearest
	default:
		mode_read = mgo.SecondaryPreferred
	}

	session.SetMode(mode_read, true)
	db = session.DB(m.Database)
	fmt.Println("connected!")
}

// GetList ...
func (m *Handler) GetList(user_id int, limit int, page int) (Response, error) {
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	var products []Product
	total, error := db.C(COLLECTION).Find(bson.M{"user_id": user_id}).Count()
	if error != nil {
		log.Fatal(error)
	}
	float_number := float64(total) / float64(limit)
	last_page := int(math.Ceil(float_number))
	paging := Paging{Total: total, PerPage: limit, CurrentPage: page, LastPage: last_page}
	err := db.C(COLLECTION).Find(bson.M{"user_id": user_id}).Sort("-$natural").Skip(offset).Limit(limit).All(&products)

	if len(products) == 0 {
		products = []Product{}
	}

	return Response{Data: products, Paging: paging}, err
}

// CreateProducts
func (m *Handler) CreateProducts(user_id int, product_ids string) Create {
	pids := strings.Split(product_ids, ",")
	for _, pid := range pids {
		product_id, error := strconv.Atoi(pid)
		if error != nil {
			log.Fatal(error)
		}
		created_at := time.Now().Format("2006-01-02 15:04:05")
		microtime := time.Now().UnixNano() / int64(time.Millisecond)
		product := bson.M{"product_id": int32(product_id), "user_id": int32(user_id), "created_at": created_at, "microtime": float64(microtime)}
		selecter := bson.M{"product_id": int32(product_id), "user_id": int32(user_id)}

		_, err := db.C(COLLECTION).Upsert(selecter, product)
		if err != nil {
			log.Fatal(error)
		}
	}

	return Create{Result: "success"}
}
