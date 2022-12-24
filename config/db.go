package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var host *Config = GetConfig("./config/config.toml")

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(host.DB.Host))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("✅ Database Connected")
	return client
}

var DB *mongo.Client = ConnectDB()

func SelectCol(client *mongo.Client, colName string) *mongo.Collection {
	col := client.Database(host.DB.Database).Collection(colName)
	return col
}
