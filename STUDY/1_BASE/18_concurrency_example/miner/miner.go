package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	num int,
	power int,
) {
	defer wg.Done()

	for {
		fmt.Println("Я шахтер номер:", num, ", начал добывать уголь")
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", num, ", мой рабочий день закончен!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я шахтер номер:", num, ", добыл уголь:", power)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", num, ", мой рабочий день закончен!")
			return
		case transferPoint <- power:
			fmt.Println("Я шахтер номер:", num, ", передал уголь:", power)
		}

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Я шахтер номер:", num, ", мой рабочий день закончен!")
		// 	return
		// default:
		// 	fmt.Println("Я шахтер номер:", num, ", начал добывать уголь")
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Println("Я шахтер номер:", num, ", добыл уголь:", power)

		// 	transferPoint <- power

		// 	fmt.Println("Я шахтер номер:", num, ", передал уголь:", power)
		// }
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
