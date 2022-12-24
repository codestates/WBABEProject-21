package model

import (
	"pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Nickname     string    `bson:"nickname"`
	RegisteredAt time.Time `bson:"registered_at"`
}

var Muser *User
var Cuser *mongo.Collection = config.SelectCol(config.DB, "user")
