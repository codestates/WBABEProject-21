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

type Order struct {
	Id           primitive.ObjectID   `bson:"_id, omitempty"`
	Menu         []primitive.ObjectID `bson:"menu"`
	/*
	Enum을 활용해보는 것은 어떨까요? int형으로 저장하고 주석을 남겨두면 어떤 값인지는 확인할 수 있지만
	숫자로 저장되기에 무슨 상태인지 확인하려면 코드를 다시 확인해야하는 번거로움이 있습니다.
	String 형식의 Enum을 활용해보시는 것을 추천 드립니다.
	*/
	Status       int                  `bson:"status"` // 1=주문접수 || 2=조리중(취소불가) || 3=완료 || 4=주문취소
	Paid         int                  `bson:"paid, omitempty"`
	User         primitive.ObjectID   `bson:"user"`
	/*
	주문의 경우에도 Menu와 같이 created_at, updated_at 필드를 가지는 것이 좋아보입니다.
	거의 모든 경우에 위의 두 값은 필수적으로 들어가야 디버깅에 좋습니다.
	*/
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
	/*
	라우터 부분에서도 코멘트를 드렸지만, 인풋값으로 받기보단 nested한 구성을 통해서 메뉴의 아이디를 가져오는 편이 좋아보입니다.
	이렇게 하는 경우 어쨌든 Database에 접속하고 데이터를 가져오는 시간만큼 지연이 됩니다. 
	*/
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

func (o *Order) GetOrderByUser(user_id string, limit, offset int) ([]Order, error) {
	var result []Order
	limit_int64 := int64(limit)
	offset_int64 := int64(offset)
	user_id_toObj, _ := primitive.ObjectIDFromHex(user_id)
	fmt.Println(user_id_toObj)

	Opt := options.Find().SetSort(bson.D{{"createdAt", -1}}).SetLimit(limit_int64).SetSkip(offset_int64)

	cur, err := Corder.Find(context.TODO(), bson.D{{"user", user_id_toObj}}, Opt)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("FIND ERR")
	}
	cur.Decode(&result)
	cur.All(context.TODO(), &result)

	return result, nil
}

func (o *Order) UpdateOrder(order Order, newMenu []primitive.ObjectID) error {
	var tp int
	// 유저 인풋의 Menu ID 값을 통해 Menu Table의 Price_Won 조회
	for _, menuId := range newMenu {
		var menu Menu
		Cmenu.FindOne(context.TODO(), bson.M{"_id": menuId}).Decode(&menu)
		tp += menu.Price_won
	}

	update := bson.M{
		"menu": newMenu,
		"paid": tp,
	}

	_, err := Corder.UpdateOne(context.TODO(), bson.M{"_id": order.Id}, bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}
