package utils

import "time"

func GetCurrentTimeInSeconds() int {
	return int(time.Now().Unix())
}

func GetCurrentTimeInNanoSeconds() int {
	return int(time.Now().UnixNano())
}
