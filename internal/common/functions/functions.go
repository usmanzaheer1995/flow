package functions

import (
	"log"
	"strconv"
)

func IntToString(key int) string {
	value := strconv.Itoa(key)
	return value
}

func StringToInt(key string) int {
	if key == "" {
		return 0
	}
	value, err := strconv.Atoi(key)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
