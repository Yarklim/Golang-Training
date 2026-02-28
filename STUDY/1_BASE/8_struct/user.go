package structures

import "fmt"

// Типизация struct
type User struct {
	name        string  // ""
	age         int     // 0
	phoneNumber string  // ""
	isClose     bool    // false
	rating      float64 // 0.0
}

// Конструктор
func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		fmt.Println("Name is invalide!")
		return User{}
	}
	if age <= 0 || age >= 150 {
		fmt.Println("Age is invalide!")
		return User{}
	}
	if phoneNumber == "" {
		fmt.Println("PhoneNumber is invalide!")
		return User{}
	}
	if rating < 0 || rating > 10 {
		fmt.Println("Rating is invalide!")
		return User{}
	}

	return User{
		name:        name,
		age:         age,
		phoneNumber: phoneNumber,
		isClose:     isClose,
		rating:      rating,
	}
}
