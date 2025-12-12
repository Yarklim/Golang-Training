package paymentmethods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (payPal PayPal) Pay(sum int) int {
	fmt.Println("Оплата с помощью PayPal")
	fmt.Println("Сумма:", sum, "USD")

	// Рандомный ID
	return rand.Int()
}

func (payPal PayPal) Cancel(id int) {
	fmt.Println("Отмена операции! ID:", id)
}
