package controller

import "go.mongodb.org/mongo-driver/bson/primitive"

type ReqForm_CreateOrder struct {
	Menu     []primitive.ObjectID `bson:"menu"`
	Nickname string               `bson:"nickname"`
}

type RespForm_CreateOrder struct {
}
