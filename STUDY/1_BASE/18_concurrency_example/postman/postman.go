package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string,
	num int,
	mail string,
) {
	defer wg.Done()

	for {
		fmt.Println("Я почтальон номер:", num, ", взял письмо")

		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер:", num, ", мой рабочий день закончен!")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Я почтальон номер:", num, ", донес письмо до почты:", mail)
		}

		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер:", num, ", мой рабочий день закончен!")
			return
		case transferPoint <- mail:
			fmt.Println("Я почтальон номер:", num, ", передал письмо:", mail)
		}

		// select {
		// case <-ctx.Done():
		// 	fmt.Println("Я почтальон номер:", num, ", мой рабочий день закончен!")
		// 	return
		// default:
		// 	fmt.Println("Я почтальон номер:", num, ", взял письмо")
		// 	time.Sleep(1 * time.Second)
		// 	fmt.Println("Я почтальон номер:", num, ", донес письмо до почты:", mail)

		// 	transferPoint <- mail

		// 	fmt.Println("Я почтальон номер:", num, ", передал письмо:", mail)
		// }
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Семейный привет",
		2: "Приглашение от друга",
		3: "Информация из автосервиса",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}

	return mail
}
