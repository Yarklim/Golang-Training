package arrayslicemap

import "fmt"

func MapArray() {
	// weatherMonth := make(map[int]int, 31) // можно создавать как и в slice

	weatherWeek := map[string]int{
		"Mon": +3,
		"Tue": +6,
		"Wed": +9,
		"Thu": -4,
		"Fri": +1,
		"Sat": +1,
		"Sun": +1,
	}

	weatherWeek["Tue"] = +5 // добавление/изменение ключа-значения

	c, ok := weatherWeek["Gok"] // проверка на валидность значения ключа

	if !ok {
		fmt.Println("Такого дня не существует")
	}

	// Иттерация по map (порядок элементов не сохраняется)
	for k, v := range weatherWeek {
		weatherWeek[k] += 1
		fmt.Println(k, v) // здесь копия массива!
	}

	fmt.Println(weatherWeek["Sat"])
	fmt.Println(weatherWeek["Tue"])
	fmt.Println(weatherWeek["Gok"]) // если ключа нет, то выведет 0 (значение по умолчанию типа int)
	fmt.Println(c, ok)              // выведет 0 false
}
