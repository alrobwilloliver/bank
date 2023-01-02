package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// create a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return int64(rand.Intn(1000000))
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, GBP}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
