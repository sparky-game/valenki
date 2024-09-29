package common

import (
  "math/rand"
)

func rollDie() int {
  return rand.Intn(6) + 1
}

func RollDice(n int) int {
  sum := 0
  for i := 0; i < n; i++ {
    sum += rollDie()
  }
  return sum
}
