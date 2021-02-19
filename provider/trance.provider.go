package provider

import (
	"time"
	"math/rand"
)

const length = 12
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewTranceId() string {
	tranceId := make([]byte, length)

	for i := range tranceId {
		tranceId[i] = letters[rand.Int63() % int64(len(letters))]
	}
	return string(tranceId)
}
