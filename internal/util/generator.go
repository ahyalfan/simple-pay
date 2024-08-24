package util

import (
	"math/rand"
	"time"
)

func GeneratorRandomString(n int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charsets := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	latters := make([]rune, n)
	for i := range latters {
		latters[i] = charsets[random.Intn(len(charsets))]
	}
	return string(latters)
}
func GeneratorRandomNumber(n int) string {
	random := rand.New(rand.NewSource(time.Now().Unix()))
	charsets := []rune("0123456789")
	latters := make([]rune, n)
	for i := range latters {
		latters[i] = charsets[random.Intn(len(charsets))]
	}
	return string(latters)
}
