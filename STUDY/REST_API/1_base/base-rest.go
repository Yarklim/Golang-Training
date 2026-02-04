package base_rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
========= Query Params =========
fooParam := r.URL.Query().Get("foo")
booParam := r.URL.Query().Get("boo")
*/

var (
	money = 1000
	bank  = 0
	mtx   = sync.Mutex{} // Создаю мьютекс для блокировки критических секций
)

// ====== Структура для хранения истории покупок (JSON) ======
type PaymentInfo struct {
	ProductName string `json:"productName"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
	Amount      int    `json:"amount"`
	Customer    string `json:"customer"`
	CreatedAt   time.Time
}

var paymentHistory = make([]PaymentInfo, 0)

func (p PaymentInfo) Validate() bool {
	if p.Amount == 0 || p.Price == 0 || p.Quantity == 0 || p.ProductName == "" || p.Customer == "" {
		return false
	}

	return true
}

func (p PaymentInfo) Println() {
	fmt.Println("Product:", p.ProductName)
	fmt.Println("Product price:", p.Price)
	fmt.Println("Quantity:", p.Quantity)
	fmt.Println("Amount:", p.Amount)
	fmt.Println("Customer:", p.Customer)
}

// ====================================================

// ============= HTTP Response =============
type HttpResponse struct {
	Balance     int         `json:"balance"`
	PaymentInfo PaymentInfo `json:"paymentInfo"`
}

// =========================================

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
	// ======== Для общего понимания: ========
	// =============== Хедеры ================
	for key, value := range r.Header {
		fmt.Println("key:", key, "-- value:", value)
	}
	// =============== Методы ================
	fmt.Println("HTTP method:", r.Method)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed) // 405 Method Not Allowed
	}
	// =======================================

	var payment PaymentInfo
	//! ======== Внимание: в стандартной библиотеке json нет проверки на отсутствие поля, и будут переданы дефолтные значения типов! ========
	//! ======== Обязательно делать валидацию Body =========
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payment.CreatedAt = time.Now()
	payment.Println()

	mtx.Lock() // Блокирую критическую секцию обработки глобальной переменной
	if money-payment.Amount >= 0 {
		money -= payment.Amount

		httpResponse := HttpResponse{
			Balance:     money,
			PaymentInfo: payment,
		}

		b, err := json.Marshal(httpResponse)
		if err != nil {
			fmt.Println("err:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(b); err != nil {
			fmt.Println("err:", err)
			return
		}

		fmt.Println("Оплата прошла успешно, остаток в кошельке: " + strconv.Itoa(money))
		fmt.Println("Payment info:", payment)
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

	// Вариант условного ветвления при котором переменная (err) исчезнет после выполнения условия
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("HTTP server error:", err)
	}
	fmt.Println("Программа закончила свое выполнение")
}
