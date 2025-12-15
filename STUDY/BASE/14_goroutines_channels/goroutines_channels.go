package goroutineschannels

import (
	"fmt"
	"time"
)

func provider(transferPoint chan int, id int) {
	fmt.Println("Start", id)
	time.Sleep(1 * time.Second)
	fmt.Println("End", id)

	transferPoint <- 10
}

func MainGoroutines() {
	// count := 0

	transferPoint := make(chan int) // Канал

	go provider(transferPoint, 1) // Горутина 1 и передача канала в аргумент
	go provider(transferPoint, 2) // Горутина 2
	go provider(transferPoint, 3) // Горутина 3
}
