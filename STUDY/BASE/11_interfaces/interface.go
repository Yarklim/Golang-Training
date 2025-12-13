package interfaces

import (
	paymentmethods "study/STUDY/BASE/11_interfaces/payment_methods"

	"github.com/k0kubun/pp"
)

func Interfaces() {
	method := paymentmethods.NewPayPal()

	paymentModule := NewPaymentModule(method)

	paymentModule.Pay("Notebook", 1550)
	idPhone := paymentModule.Pay("Phone", 1100)

	allInfo := paymentModule.GetAllPaymentsInfo()
	pp.Println("All purchases:", allInfo)

	phoheInfo := paymentModule.GetPaymentInfo(idPhone)
	pp.Println(phoheInfo)
}
