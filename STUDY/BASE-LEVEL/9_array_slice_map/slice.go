package arrayslicemap

import "fmt"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func SliceArray() {
	usersArray := []User{
		{
			Name:    "Yar",
			Rating:  9.5,
			Premium: true,
		},
		{
			Name:    "John",
			Rating:  5.5,
			Premium: false,
		},
		{
			Name:    "Marta",
			Rating:  7.5,
			Premium: true,
		},
	}

	fmt.Println("Users Array before:", usersArray)

	// Добавление элемента в массив - append(arr, el)
	usersArray = append(
		usersArray,
		User{
			Name:    "Anna",
			Rating:  8.5,
			Premium: true,
		})

	fmt.Println("len:", len(usersArray)) // Длина массива
	fmt.Println("cap:", cap(usersArray)) // Вместимость массива

	// Классический перебор массива
	// for i := 0; i < len(usersArray); i++ {
	// 	difference := 10.0 - usersArray[i].Rating

	// 	if usersArray[i].Premium {
	// 		if difference >= 1 {
	// 			usersArray[i].Rating += 1
	// 		} else {
	// 			usersArray[i].Rating = 10
	// 		}
	// 	}
	// }

	// Продвинутый перебор массива (в данном случае создается копия массива!)
	// Если надо менять исходный массив, то необходимо по индексу обращаться к элементу исходного массива и менять его.
	for _, user := range usersArray {
		fmt.Println(user.Name)
	}

	for i, user := range usersArray {
		difference := 10.0 - usersArray[i].Rating

		if user.Premium {
			if difference >= 1 {
				usersArray[i].Rating += 1
			} else {
				usersArray[i].Rating = 10
			}
		}
	}

	fmt.Println("Users Array after:", usersArray)
}
