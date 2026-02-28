package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"concurrency/miner"
	"concurrency/postman"
)

func main() {
	var coal atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("----->>>>Рабочий день шахтеров окончен!")
		minerCancel()
	}()

	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("----->>>>Рабочий день почтальонов окончен!")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 6)
	mailTransferPoint := postman.PostmanPool(postmanContext, 4)

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range coalTransferPoint {
			coal.Add(int64(val))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for val := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, val)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Coal summary:", coal.Load())

	mtx.Lock()
	fmt.Println("Mails summary:", len(mails))
	mtx.Unlock()
}
