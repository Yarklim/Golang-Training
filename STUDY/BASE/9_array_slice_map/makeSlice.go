package arrayslicemap

import "fmt"

func MakeSlice() {
	// Создание пустого slice, с последующим добавлением элементов
	// Если надо определить длину сразу, то - make([]int, 0, 5), где 5 - это длина массива (будет массив из пяти 0)
	// Длина массива (len) будет равна 0, а вместимость (cap) равна 5
	intSlice := make([]int, 0)

	intSlice = append(intSlice, 5, 10, 8, 22)

	fmt.Println(intSlice)
}
