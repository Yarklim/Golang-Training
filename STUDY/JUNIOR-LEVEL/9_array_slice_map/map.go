package arrayslicemap

import "fmt"

func MapArray() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}

	c, ok := weather[30] // проверка на валидность значения ключа

	fmt.Println(weather[15])
	fmt.Println(weather[10]) // если ключа нет, то выведет 0 (значение по умолчанию типа int)
	fmt.Println(c, ok)       // выведет 0 false
}
