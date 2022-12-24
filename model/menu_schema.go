package model

import (
	"context"
	"errors"
	"fmt"
	"pkg/config"
	"pkg/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Menu struct {
	Name        string    `bson:"name"`
	Is_active   bool      `bson:"is_active"`
	Amount      int       `bson:"amount"`
	MadeIn      string    `bson:"madein"`
	Price_won   int       `bson:"price_won"`
	Today_menu  bool      `bson:"today_menu"`
	Spicy_level int       `bson:"spicy_level"`
	CreatedAt   time.Time `bson:"createdAt"`
	UpdateAt    time.Time `bson:"updateAt"`
}

var Mmenu *Menu
var Cmenu *mongo.Collection = config.SelectCol(config.DB, "menu")

func (m *Menu) CreateMenu(input Menu) error {
	currentTime := utils.MongoTime()
	input.CreatedAt = currentTime
	input.UpdateAt = currentTime
	_, err := Cmenu.InsertOne(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (m *Menu) UpdateMenu(input Menu, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	currentTime := utils.MongoTime()
	update := bson.M{
		"name":        input.Name,
		"is_active":   input.Is_active,
		"amount":      input.Amount,
		"madein":      input.MadeIn,
		"price_won":   input.Price_won,
		"today_menu":  input.Today_menu,
		"spicy_level": input.Spicy_level,
		"updateAt":    currentTime,
	}
	_, err := Cmenu.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (m *Menu) DeleteMenu(id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	_, err := Cmenu.DeleteOne(context.TODO(), bson.M{"_id": objId})
	if err != nil {
		return err
	}
	return nil
}

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
