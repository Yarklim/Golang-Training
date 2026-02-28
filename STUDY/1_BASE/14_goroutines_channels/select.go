package goroutineschannels

import (
	"fmt"
	"strconv"
	"time"
)

func Select() {
	intChannel := make(chan int)
	strChannel := make(chan string)

	// Блокирующий канал на 1 сек
	go func() {
		i := 1
		for {
			intChannel <- i
			i++

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		i := 1
		for {
			strChannel <- "hi" + strconv.Itoa(i)
			i++
		}
	}()

	// Селект, читающий данные из каналов, как только они приходят
	for {
		select {
		case number := <-intChannel:
			fmt.Println("intChannel:", number)
		case str := <-strChannel:
			fmt.Println("strChannel:", str)
		default:
			fmt.Println("Никакой канал не готов")
		}
	}
}
