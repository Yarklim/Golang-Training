package arrayslicemap

import "fmt"

// ========= Статический массив (применяется редко) ==========
func StaticArray() {
	staticArr := [5]int{7, 55, 8, 94, 22}
	staticArrIdx3 := staticArr[3]

	fmt.Println(staticArr)

	staticArr[1] += 5

	fmt.Println(staticArr)
	fmt.Println(staticArrIdx3)
}
