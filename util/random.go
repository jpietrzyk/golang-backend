package util

import (
	"math/rand"
	"strings"
	"time"
)

const chars = "abcdefghijklmnopqrstuwyxz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(chars)

	for i := 0; i < n; i++ {
		c := chars[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUserName() string {
	return RandomString(8)
}

func RandomUserEmail() string {
	domains := []string{"com", "org", "uk", "pl"}

	return RandomString(6) + "@" + RandomString(4) + "." + domains[rand.Intn(len(domains))]
}
