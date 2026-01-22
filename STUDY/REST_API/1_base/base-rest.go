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
		w.WriteHeader(http.StatusInternalServerError) // Статус код
		msg := "Fail to read HTTP body: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
		return
	}

	httpReqBodyStr := string(httpReqBody) // преобразую []byte в строку

	paymentAmount, err := strconv.Atoi(httpReqBodyStr) // преобразую строку в число
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Статус код
		msg := "Fail to convert HTTP body to integer: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock() // Блокирую критическую секцию обработки глобальной переменной
	if money-paymentAmount >= 0 {
		money -= paymentAmount

		msg := "Оплата прошла успешно, остаток в кошельке: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err = w.Write([]byte(msg)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
	}
	mtx.Unlock() // Разблокирую критическую секцию
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpReqBody, err := io.ReadAll(r.Body) // получаю тело запроса в виде []byte
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Статус код
		msg := "Fail to read HTTP body: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
		return
	}

	httpReqBodyStr := string(httpReqBody) // преобразую []byte в строку

	saveAmount, err := strconv.Atoi(httpReqBodyStr) // преобразую строку в число
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Статус код
		msg := "Fail to convert HTTP body to integer: " + err.Error()
		fmt.Println(msg)
		_, err = w.Write([]byte(msg)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
		return
	}

	mtx.Lock() // Блокирую критическую секцию обработки глобальной переменной
	if money >= saveAmount {
		// взять деньги из кошелька и положить в банк
		money -= saveAmount
		bank += saveAmount

		msgPay := "Остаток средств в кошельке: " + strconv.Itoa(money)
		fmt.Println(msgPay)
		_, err = w.Write([]byte(msgPay)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}

		msgBank := "Сумма на счету в банке: " + strconv.Itoa(bank)
		fmt.Println(msgBank)
		_, err = w.Write([]byte(msgBank)) // Тело ответа
		if err != nil {
			fmt.Println("Fail to write HTTP response:", err)
		}
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
