package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomRune() rune {
	f := rand.Float64()
	var min, max rune

	switch {
	case f < float64(1)/3:
		min, max = 'a', 'z'
	case f < float64(2)/3:
		min, max = 'A', 'Z'
	default:
		min, max = '0', '9'
	}

	delta := int(max - min)
	i := rand.Intn(delta + 1)

	return min + rune(i)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const passwordLen = 16
	passwordRunes := make([]rune, passwordLen)

	for i := range passwordRunes {
		passwordRunes[i] = randomRune()
	}

	fmt.Println(string(passwordRunes))
}
