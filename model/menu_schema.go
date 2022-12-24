package model

import (
	"fmt"
	"pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (m *Menu) GetMenuPaging(page, limit int) {
	fmt.Println("GetMenuPaging")
}
