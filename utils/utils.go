package utils

import (
	"fmt"
	"log"
	"time"
)

/*
유틸성을 지니는 코드들의 경우 따로 분리해주신점 정말 좋습니다.
공통적으로 많이 사용되는 부분들은 이렇게 유틸로 따로 분리하면 여러 곳에서 가져다 사용하기에 용이하고, 테스트를 작성하는 것도 쉬워집니다.
*/

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

/*
현재의 시간을 가져오는 것이므로 GetCurrentTime 같은 네이밍이 더 좋아보입니다.
MongoTime이라는 네이밍은 직관적이지 않습니다.
*/
func MongoTime() time.Time {
	currentTime := time.Now()
	_, offset := currentTime.Zone()
	mongoTime := currentTime.Add(time.Second * time.Duration(offset))
	return mongoTime
}
