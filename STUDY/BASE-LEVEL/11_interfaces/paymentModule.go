package interfaces

import "maps"

// Интерфейс
type PaymentMethod interface {
	Pay(sum int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

// Конструктор
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принимает: 1. Описание оплаты; 2. Сумма оплаты.
// Возвращает: ID проведенной операции.
func (payment *PaymentModule) Pay(description string, sum int) int {
	id := payment.paymentMethod.Pay(sum)

	info := PaymentInfo{
		Description: description,
		Sum:         sum,
		Cancelled:   false,
	}

	payment.paymentsInfo[id] = info

	return id
}

// Метод CancelPayment()
// Принимает: 1. ID операции.
// Возвращает: статус операции.
func (payment *PaymentModule) CancelPayment(id int) {
	info, ok := payment.paymentsInfo[id]
	if !ok {
		return
	}

	payment.paymentMethod.Cancel(id)

	info.Cancelled = true

	payment.paymentsInfo[id] = info
}

// Метод GetPaymentInfo()
// Принимает: 1. ID операции.
// Возвращает: информацию о платеже.
func (payment *PaymentModule) GetPaymentInfo(id int) PaymentInfo {
	info, ok := payment.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод GetAllPaymentsInfo()
// Принимает: ничего не принимает.
// Возвращает: информацию о всех платежах.
func (payment *PaymentModule) GetAllPaymentsInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(payment.paymentsInfo))

	maps.Copy(tempMap, payment.paymentsInfo)

	return tempMap
}
