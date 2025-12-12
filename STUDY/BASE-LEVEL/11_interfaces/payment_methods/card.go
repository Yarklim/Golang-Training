package paymentmethods

import (
	"fmt"
	"math/rand"
)

type Card struct{}

func NewCard() Card {
	return Card{}
}

func (card Card) Pay(sum int) int {
	fmt.Println("Оплата картой")
	fmt.Println("Сумма:", sum, "USD")

	// Рандомный ID
	return rand.Int()
}

func (card Card) Cancel(id int) {
	fmt.Println("Отмена операции! ID:", id)
}
