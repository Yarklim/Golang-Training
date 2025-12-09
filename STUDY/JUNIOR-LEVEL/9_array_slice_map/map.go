package arrayslicemap

import "fmt"

func MapArray() {
	// weather := make(map[int]int, 10) // можно создавать как и в slice

	weather := map[int]int{
		11: +3,
		12: +6,
		13: +9,
		14: -4,
		15: +1,
	}

	weather[20] = +5 // добавление ключа-значения

	c, ok := weather[30] // проверка на валидность значения ключа

	// Иттерация по map (порядок элементов не сохраняется)
	for k, v := range weather {
		weather[k] += 1
		fmt.Println(k, v) // здесь копия массива!
	}

	fmt.Println(weather[12])
	fmt.Println(weather[20])
	fmt.Println(weather[10]) // если ключа нет, то выведет 0 (значение по умолчанию типа int)
	fmt.Println(c, ok)       // выведет 0 false
}
