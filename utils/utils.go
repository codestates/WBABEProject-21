package utils

import (
	"fmt"
	"log"
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
