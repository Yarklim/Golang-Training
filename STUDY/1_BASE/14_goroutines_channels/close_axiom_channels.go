package goroutineschannels

import (
	"fmt"
	"math/rand"
	"time"
)

func ClosedChannels() {
	// Создаю новый открытый канал
	channel1 := make(chan int)

	// Закрываю канал
	// Если попытаться закрыть закрытый канал - будет паника
	close(channel1)

	// Чтение из закрытого канала - значение по умолчанию
	v1, ok := <-channel1
	fmt.Println(ok) // Если false, значит прочитанное значение - значение по умолчанию
	fmt.Println(v1) // 0 - значение по умолчанию типа int

	// Запись в закрытый канал - паника

	// var nilChannel chan int // nil канал
	// Если попытаться закрыть nil канал - будет паника
	// Чтение из nil канала - block
	// Запись в nil канала - block

	// Пример:
	transferPoint := make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)
		fmt.Println("Iterations:", iterations)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}

		close(transferPoint)
	}()

	count := 0
	// for {
	// 	val, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("Channel close")
	// 		break
	// 	}
	// 	count += val

	// 	fmt.Println("count:", count)
	// }

	for val := range transferPoint {
		count += val
		fmt.Println("count:", count)
	}

	fmt.Println("Final count:", count)
}
