package arrayslicemap

import "fmt"

func MapArray() {
	// weather := make(map[string]int, 10) // можно создавать как и в slice

	weather := map[string]int{
		"Mon": +3,
		"Tue": +6,
		"Wed": +9,
		"Thu": -4,
		"Fri": +1,
		"Sat": +1,
		"Sun": +1,
	}

	weather["Tue"] = +5 // добавление/изменение ключа-значения

	c, ok := weather["Gok"] // проверка на валидность значения ключа

	// Иттерация по map (порядок элементов не сохраняется)
	for k, v := range weather {
		weather[k] += 1
		fmt.Println(k, v) // здесь копия массива!
	}

	fmt.Println(weather["Sat"])
	fmt.Println(weather["Tue"])
	fmt.Println(weather["Gok"]) // если ключа нет, то выведет 0 (значение по умолчанию типа int)
	fmt.Println(c, ok)          // выведет 0 false
}
