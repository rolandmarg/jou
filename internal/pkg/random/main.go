package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns random positive integer
func RandomInt() int {
	return rand.Int()
}

// RandomInt64 returns random positive 64 bit integer
func RandomInt64() int64 {
	return rand.Int63()
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

// RandomString returns random string
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RandomStrings returns slice of random strings
func RandomStrings(strLen, sliceLen int) []string {
	s := make([]string, sliceLen)
	for i := range s {
		s[i] = RandomString(strLen)
	}

	return s
}
