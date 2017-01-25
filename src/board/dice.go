package board

import (
	"math/rand"
)

func DiceRoll() uint {
	return uint(rand.Intn(6) + 1)
}
