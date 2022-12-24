package model

import (
	"context"
	"errors"
	"pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Menu struct {
	Name        string              `bson:"name"`
	Is_active   bool                `bson:"is_active"`
	Amount      int                 `bson:"amount"`
	MadeIn      string              `bson:"madein"`
	Price_won   int                 `bson:"price_won"`
	Today_menu  bool                `bson:"today_menu"`
	Spicy_level int                 `bson:"spicy_level"`
	CreatedAt   time.Time           `bson:"createdAt"`
	UpdateAt    primitive.Timestamp `bson:"updateAt"`
}

var Mmenu *Menu
var Cmenu *mongo.Collection = config.SelectCol(config.DB, "menu")

func (m *Menu) GetMenuPaging(limit, offset int) ([]Menu, error) {
	var result []Menu
	limit_int64 := int64(limit)
	offset_int64 := int64(offset)
	Opt := options.Find().SetSort(bson.D{{"createdAt", -1}}).SetLimit(limit_int64).SetSkip(offset_int64)
	cur, err := Cmenu.Find(context.TODO(), bson.M{}, Opt)
	if err != nil {
		return nil, errors.New("FIND ERR")
	}
	cur.Decode(&result)
	cur.All(context.TODO(), &result)

	return result, nil
}
