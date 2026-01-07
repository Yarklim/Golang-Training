package base_rest

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	str := "Default handler"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("HTTP запрос обработан корректно")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Pay handler"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время оплаты произошла ошибка:", err.Error())
	} else {
		fmt.Println("Оплата произведена!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Cancel handler"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время отмены оплаты произошла ошибка:", err.Error())
	} else {
		fmt.Println("Отмена оплаты произведена!")
	}
}

func BaseRest() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	fmt.Println("Запускаю HTTP сервер")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка сервера:", err.Error())
	}
	fmt.Println("Программа закончила свое выполнение")
}
