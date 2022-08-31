package utils

import (
	"math/rand"
	"strings"
	"time"
)

const Alphabes = "abcdefghijklmnpqrstuvwxyz"
const phones = "0123456789"
const composedConstants = "abcdefghijklmnpqrstuvwxyz0123456789$#@!%*&)"

func init() {
	rand.Seed(time.Now().Unix())
}

func RandomeString(n int, givenstring string) string {
	var sb strings.Builder
	k := len(givenstring)
	for i := 0; i < n; i++ {
		c := givenstring[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
func RandomUserName() string {
	return RandomeString(8, Alphabes)
}

func RandomPassword() string {
	return RandomeString(9, composedConstants)
}
func RandomeEmail() string {
	email := RandomeString(5, Alphabes) + "@gmail.com"
	return email
}
func RandomePhoneNumber() string {
	phoneNumber := "+251" + RandomeString(9, phones)
	return phoneNumber
}
