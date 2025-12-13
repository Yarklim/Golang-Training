package userinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserInputs() {
	scanner := bufio.NewScanner(os.Stdin) // сканирование потока данных пользовательского ввода из консоли

	fmt.Print("Введите команду: ")

	ok := scanner.Scan()
	if !ok {
		fmt.Println("Ошибка ввода")
		return
	}

	// можно ещё и так проверить
	// if ok := scanner.Scan(); !ok {
	// 	fmt.Println("Ошибка ввода")
	// 	return
	// }

	text := scanner.Text()
	fields := strings.Fields(text) // создаем slice из слов ввода

	if len(fields) == 0 {
		fmt.Println("Вы ничего не ввели!")
		return
	}

	fmt.Println("Fields:", fields)
	fmt.Println("Первое слово:", fields[0])
}
