package model

import (
	"pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	Menu         primitive.ObjectID `bson:"menu"`
	Status       int                `bson:"status"` // 1=주문접수 || 2=조리중(취소불가) || 3=완료 || 4=주문취소
	Paid         string             `bson:"paid"`
	User         primitive.ObjectID `bson:"user"`
	ReceiptionAt time.Time          `bson:"receiptionAt"`
}

var Morder *Menu
var Corder *mongo.Collection = config.SelectCol(config.DB, "order")
