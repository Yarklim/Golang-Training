package main

import (
	"fmt"
	"reflect"
)

func main() {
	textType := "Hello, Yar!" // Присвоение значения в переменную
	intType := 10             // Целое число нельзя изменить на дробное число
	floatType := 0.5          // Дробное число нельзя изменить на целое число
	boolType := true

	floatTypeResult := 2 * floatType  // Результат будет целым числом, но тип будет float64
	intResult := int(floatTypeResult) // Приведение к типу int

	resDevision := intType / 3 // Результат будет 3, а не 3.33333. Если делитель и делимое целые числа, то результат всегда будет целым числом (целочисленное деление)
	// Чтобы результат деления был правильным, делимое должно быть типа float64

	remainderFromDivision := intType % 3 // Остаток от деления. Здесь равен 1.

	fmt.Println(textType) // Вывод с переносом
	fmt.Println(intType)  // Вывод в одну строку
	fmt.Println(floatType)
	fmt.Println(boolType)
	fmt.Println(floatTypeResult)
	fmt.Println("Тип floatTypeResult:", reflect.TypeOf(floatTypeResult))
	fmt.Println("Тип intResult:", reflect.TypeOf(intResult))
	fmt.Println(resDevision)
	fmt.Println(remainderFromDivision)
}
