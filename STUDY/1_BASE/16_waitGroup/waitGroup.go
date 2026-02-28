package waitgroup_study

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнес газету", text, "в", i, "раз")
		time.Sleep(300 * time.Millisecond)
	}
}

func WaitGroupTest() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	postman(wg, "News")

	wg.Add(1)
	postman(wg, "Sport")

	wg.Add(1)
	postman(wg, "Films")

	wg.Wait()

	fmt.Println("WaitGroupTest ended")
}
