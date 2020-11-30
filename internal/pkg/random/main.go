package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Int returns random positive integer
func Int() int {
	return rand.Int()
}

// Int64 returns random positive 64 bit integer
func Int64() int64 {
	return rand.Int63()
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

// String returns random string
func String(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Strings returns slice of random strings
func Strings(strLen, sliceLen int) []string {
	s := make([]string, sliceLen)
	for i := range s {
		s[i] = String(strLen)
	}

	return s
}

// Time generates random time.Time
func Time() time.Time {
	// https://stackoverflow.com/a/43497333
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
