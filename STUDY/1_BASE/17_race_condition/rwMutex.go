package racecondition_study

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0

var rwMtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		rwMtx.Lock()
		likes++
		rwMtx.Unlock()
	}
}

func getLikes(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		rwMtx.RLock()
		_ = likes
		rwMtx.RUnlock()
	}
}

func Likes() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i < 100; i++ {
		wg.Add(1)
		go getLikes(wg)
	}

	wg.Wait()

	fmt.Println("Время выполнения:", time.Since(initTime))
}
