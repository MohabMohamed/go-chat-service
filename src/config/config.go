package config

import (
	"os"
	"strconv"
	"time"
)

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

var TimeInSecToPersist time.Duration

func Init() {
	timeInSeconds, _ := strconv.Atoi(GetEnv("PERSIST_TIME_IN_SEC", "5"))
	TimeInSecToPersist = time.Duration(time.Second * time.Duration(timeInSeconds))

}
