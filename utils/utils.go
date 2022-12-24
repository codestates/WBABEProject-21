package utils

import (
	"fmt"
	"log"
	"time"
)

func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func MongoTime() time.Time {
	currentTime := time.Now()
	_, offset := currentTime.Zone()
	mongoTime := currentTime.Add(time.Second * time.Duration(offset))
	return mongoTime
}
