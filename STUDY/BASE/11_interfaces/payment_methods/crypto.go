package paymentmethods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (crypto Crypto) Pay(sum int) int {
	fmt.Println("Оплата криптой")
	fmt.Println("Сумма:", sum, "USDT")

	// Рандомный ID
	return rand.Int()
}

func (crypto Crypto) Cancel(id int) {
	fmt.Println("Отмена крипто-операции! ID:", id)
}
