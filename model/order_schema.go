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
)

type Order struct {
	Id           primitive.ObjectID   `bson:"_id, omitempty"`
	Menu         []primitive.ObjectID `bson:"menu"`
	Status       int                  `bson:"status"` // 1=주문접수 || 2=조리중(취소불가) || 3=완료 || 4=주문취소
	Paid         int                  `bson:"paid, omitempty"`
	User         primitive.ObjectID   `bson:"user"`
	ReceiptionAt time.Time            `bson:"receiptionAt"`
}

var Morder *Order
var Corder *mongo.Collection = config.SelectCol(config.DB, "order")

func (o *Order) UpdateStatus(input Order, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"status": input.Status}
	_, err := Corder.UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) CreateOrder(input Order, discount_rate int) error {
	var OrderForm Order

	tp := 0
	// 유저 인풋의 Menu ID 값을 통해 Menu Table의 Price_Won 조회
	for _, menuId := range input.Menu {
		var menu Menu
		Cmenu.FindOne(context.TODO(), bson.M{"_id": menuId}).Decode(&menu)
		tp += menu.Price_won
	}
	// tp(총 지불액)
	// discount_rate(할인율) 5%
	// => 할인 적용 지불액(OrderForm.Paid)
	var dr float64
	dr = float64(discount_rate) / float64(100)
	discount_price := float64(tp) * dr
	fmt.Println(tp)
	fmt.Println(discount_rate)
	// 조회된 값에서, 할인율을 적용해 최종 order.Paid값을 산출하시오.
	OrderForm.Menu = input.Menu
	currentTime := utils.MongoTime()
	OrderForm.ReceiptionAt = currentTime
	OrderForm.Status = input.Status
	OrderForm.Paid = tp - int(discount_price)
	OrderForm.Id = primitive.NewObjectID()
	OrderForm.User = input.User

	_, err := Corder.InsertOne(context.TODO(), OrderForm)
	if err != nil {
		fmt.Println(err)
		return errors.New("CREATE ORDER FAILED")
	}
	return nil
}

func (o *Order) FindOrderCountByUserId(userId primitive.ObjectID) (int64, error) {
	filter := bson.D{{"user", userId}}
	count, err := Corder.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, errors.New("FIND COUND FAILED")
	}

	return count, nil
}
