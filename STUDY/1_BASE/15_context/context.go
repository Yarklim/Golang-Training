package context_test

import (
	"context"
	"fmt"
	"time"
)

// родительский контекст
func fooContext(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo has ended!", num)
			return
		default:
			fmt.Println("Foo:", num)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// дочерний контекст
func booContext(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo has ended!", num)
			return
		default:
			fmt.Println("Boo:", num)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func ContextTest() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)

	go fooContext(parentContext, 1)
	go fooContext(parentContext, 2)
	go fooContext(parentContext, 3)

	go booContext(childContext, 1)
	go booContext(childContext, 2)
	go booContext(childContext, 3)

	time.Sleep(1 * time.Second)
	// parentCancel() // Если отменить родительский контекст, то завершится и дочерний контекст
	// Отмена дочернего контекста
	childCancel()

	time.Sleep(2 * time.Second)
	// Отмена родительского контекста
	parentCancel()

	time.Sleep(3 * time.Second)
	fmt.Println("Main has ended")
}
