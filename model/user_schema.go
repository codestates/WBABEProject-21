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

type User struct {
	Id           primitive.ObjectID `bson:"_id", omitempty`
	/*
	유저의 경우에 벤더인지 고객인지에 대해서 구분할 수 필드가 필요해 보입니다.
	그 값을 이용해서, 벤더인 경우에만 메뉴를 생성할 수 있도록 제어도 가능합니다.
	*/
	Nickname     string             `bson:"nickname"`
	RegisteredAt time.Time          `bson:"registered_at, omitempty"`
}

var Muser *User
var Cuser *mongo.Collection = config.SelectCol(config.DB, "user")

func (u *User) CreateUser(nickname string) (User, error) {
	var userInput User
	userInput.Id = primitive.NewObjectID()
	userInput.Nickname = nickname
	time := utils.MongoTime()
	userInput.RegisteredAt = time
	fmt.Println(userInput)

	_, err := Cuser.InsertOne(context.TODO(), userInput)
	if err != nil {
		return userInput, errors.New("CREATE USER FAILED")
	}

	user, err := u.FindUserByNickname(nickname)
	if err != nil {
		return userInput, errors.New("CREATE USER FAILED")
	}

	return user, nil
}

/*
기존 함수와 연관성이 떨어지는 부분이라면, 이렇게 따로 분리해둔 점 좋습니다.
*/
func (u *User) FindUserByNickname(nickname string) (User, error) {
	var user User
	filter := bson.D{{"nickname", nickname}}
	Cuser.FindOne(context.TODO(), filter).Decode(&user)
	return user, nil
}
