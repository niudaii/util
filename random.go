package util

import (
	"fmt"
	"math/rand"
	"time"
)

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func RandomIp() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("101.%d.%d.%d", 1+rand.Intn(254), 1+rand.Intn(254), 1+rand.Intn(254))
}
