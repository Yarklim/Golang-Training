package base_rest

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var (
	mtx   = sync.Mutex{}   // Создаю мьютекс для блокировки критических секций
	money = atomic.Int64{} // Так как payHandler выполняется в отдельной горутине, чтобы не было гонки переменную оборачиваю в атомик
	bank  = atomic.Int64{}
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
	if money.Load()-int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))

		fmt.Println("Оплата прошла успешно, остаток в кошельке:", money.Load())
	} else {
		fmt.Println("Оплата не прошла - недостаточно средств на счету")
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
	if money.Load() >= int64(saveAmount) {
		// взять деньги из кошелька и положить в банк
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))

		fmt.Println("Остаток средств в кошельке:", money.Load())
		fmt.Println("Сумма на счету в банке:", bank.Load())
	} else {
		fmt.Println("Недостаточно средств в кошельке")
	}
	mtx.Unlock() // Разблокирую критическую секцию
}

func BaseRest() {
	money.Add(1000)

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
