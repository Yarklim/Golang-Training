package base_rest

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var (
	mtx   = sync.Mutex{} // Создаю мьютекс для блокировки критических секций
	money = 1000
	bank  = 0
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	str := "Default handler"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err)
	} else {
		fmt.Println("HTTP запрос обработан корректно")
	}
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body) // получаю тело запроса в виде []byte
	if err != nil {
		fmt.Println("Fail to read HTTP body:", err)
		return
	}

	httpReqBodyStr := string(httpReqBody) // преобразую []byte в строку

	paymentAmount, err := strconv.Atoi(httpReqBodyStr) // преобразую строку в число
	if err != nil {
		fmt.Println("Fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock() // Блокирую критическую секцию обработки глобальной переменной
	if money-paymentAmount >= 0 {
		money -= paymentAmount

		fmt.Println("Оплата прошла успешно, остаток в кошельке:", money)
	}
	mtx.Unlock() // Разблокирую критическую секцию
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body) // получаю тело запроса в виде []byte
	if err != nil {
		fmt.Println("Fail to read HTTP body:", err)
		return
	}

	httpReqBodyStr := string(httpReqBody) // преобразую []byte в строку

	saveAmount, err := strconv.Atoi(httpReqBodyStr) // преобразую строку в число
	if err != nil {
		fmt.Println("Fail to convert HTTP body to integer:", err)
		return
	}

	mtx.Lock() // Блокирую критическую секцию обработки глобальной переменной
	if money >= saveAmount {
		// взять деньги из кошелька и положить в банк
		money -= saveAmount
		bank += saveAmount

		fmt.Println("Остаток средств в кошельке:", money)
		fmt.Println("Сумма на счету в банке:", bank)
	}
	mtx.Unlock() // Разблокирую критическую секцию
}

func BaseRest() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Запускаю HTTP сервер")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("HTTP server error:", err)
	}
	fmt.Println("Программа закончила свое выполнение")
}
