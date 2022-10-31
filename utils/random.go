package utils

import (
	"math/rand"
	"strings"
	"time"
)

// Package init.
func init() {
	// Generate different random seed by time.
	rand.Seed(time.Now().UnixNano())
	// fmt.Println("random init")
}

func RandomArray(length, n int) (arr []int) {
	arr = make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}

// RandomString generates a random string with chars from ' ' to '~'
func RandomString(length int) string {
	var sb strings.Builder
	k := '~' - ' '
	for i := 0; i < length; i++ {
		ch := ' ' + rand.Int31n(k)
		sb.WriteByte(byte(ch))
	}
	return sb.String()
}
