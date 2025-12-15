package goroutineschannels

import (
	"fmt"
	"time"
)

func provider(transferPoint chan int, id int) {
	fmt.Println("Start provider:", id)
	time.Sleep(1 * time.Second)
	fmt.Println("End provider", id)

	transferPoint <- 10 // Передача данных в канал

	fmt.Println("Provider #", id, "transferred count")
}

func MainGoroutines() {
	count := 0

	// transferPoint := make(chan int) // Небуферезированный Канал
	transferPoint := make(chan int, 3) // Буферезированный Канал

	initTime := time.Now()

	go provider(transferPoint, 1) // Горутина 1 и передача канала в аргумент
	go provider(transferPoint, 2) // Горутина 2
	go provider(transferPoint, 3) // Горутина 3

	count += <-transferPoint // Запись данных из канала в переменную
	count += <-transferPoint // Запись данных из канала в переменную
	count += <-transferPoint // Запись данных из канала в переменную

	fmt.Println("Final count", count)
	fmt.Println("Time wasted", time.Since(initTime))
}
