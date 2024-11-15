package utils

import (
	"math/rand"
	"strconv"
)

func GenerateGameID() string {
	randomNum := rand.Intn(900000) + 100000

	id := strconv.Itoa(randomNum)

	return id
}
