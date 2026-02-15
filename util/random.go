package util

import (
	"math/rand"
)

func RandomMoney() int64 {
	return rand.Int63n(1000)
}

var currencies = []string{"USD", "EUR", "GBP"}

func RandomCurrency() string {
	return currencies[rand.Intn(len(currencies))]
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
