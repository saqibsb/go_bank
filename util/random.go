package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomInt will genaetate a random int between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// randomString will generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// randomOwner will generate a random Owner name
func RandomOwner() string {
	return RandomString(6)
}

// randomOwner will generate a random amount of Meney
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// randomOwner will generate a random amount of Meney
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}
