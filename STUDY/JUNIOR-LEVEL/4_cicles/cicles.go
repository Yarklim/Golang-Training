package cicles

import (
	"fmt"
	"math/rand"
	"time"
)

func Cicles() {
	score := 0

	for i := 1; i <= 5; i++ {
		score++
		randomNum := rand.Intn(5)

		if score == randomNum {
			fmt.Println("Yes!")
		} else {
			fmt.Println("No(")
		}

		time.Sleep(500 * time.Millisecond)
	}

	// Бесконечный цикл
	for {
		fmt.Println("Start")

		if rand.Intn(5) == 1 {
			fmt.Println("End")

			break
		}

	}
}
