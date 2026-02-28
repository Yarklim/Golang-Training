package errors_study

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, sum int) (int, error) {
	if user.Ballance-sum < 0 {
		// Использую fmt.Errorf и %d для подстановки числа
		return 0, fmt.Errorf("недостаточно средств, чтобы оплатить товар на сумму $%d", sum)
	}

	user.Ballance -= sum

	return sum, nil
}

func Errors() {
	user := User{
		Name:     "John",
		Ballance: 10,
	}

	pp.Println("User before:", user)
	sum, res := Pay(&user, 5)
	pp.Println("User after:", user)

	if res != nil {
		fmt.Println("Оплаты не было! Причина:", res.Error())
	} else {
		fmt.Printf("Была произведена оплата на сумму $%d", sum)
	}
}
