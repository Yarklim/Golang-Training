package racecondition_study

import (
	"fmt"
	"sync"
	// "sync/atomic"
)

// var number int = 0

// var number atomic.Int64 // Вариант с atomic (тяжело для производительности)

var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		// number++
		// number.Add(1) // Вариант с atomic
		mtx.Lock() // блокируем до начала критической секции (где может быть гонка)
		slice = append(slice, i)
		mtx.Unlock() // разблокируем после выполнения критической секции
	}
}

func Race() {
	wg := &sync.WaitGroup{}

	wg.Add(10)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()

	// fmt.Println("Number:", number)
	// fmt.Println("Number:", number.Load()) // Вариант с atomic
	mtx.Lock()
	fmt.Println("Slice len:", len(slice))
	mtx.Unlock()
}
