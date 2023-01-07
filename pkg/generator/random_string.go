package generator

import (
	"math/rand"
	"time"
)

// GenerateRandomString generate a random string given a lenght
func GenerateRandomString(length int) string {
	var letters = []rune("testuserndjdismaluehsmsldjd8a6egenidhnalsoduenkdoshsyabdlieee")
	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
