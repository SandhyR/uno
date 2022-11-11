package utils

import (
	"math/rand"
	"time"
)

func RandomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func RandomNumber() int {
	min := 1
	max := 9
	return rand.Intn(max-min) + min
}
