package util

import (
	"math/rand"
	"time"
)

const (
	FullCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringRandom(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = FullCharSet[seededRand.Intn(len(FullCharSet))]
	}
	return string(b)
}
