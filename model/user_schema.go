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

func (u *User) FindUserByNickname(nickname string) (User, error) {
	var user User
	filter := bson.D{{"nickname", nickname}}
	Cuser.FindOne(context.TODO(), filter).Decode(&user)
	return user, nil
}
