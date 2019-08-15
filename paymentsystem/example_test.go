package paymentsystem_test

import (
	"fmt"

	"github.com/khasanovbi/banksdb/v2/paymentsystem"
)

func ExampleFindPaymentSystem() {
	creditCard := "5275940000000000"
	paymentSystem := paymentsystem.FindPaymentSystem(creditCard)
	if paymentSystem != nil {
		fmt.Printf("Payment system: %s\n", *paymentSystem)
	}
}

func ExampleFindPaymentSystemByPrefix() {
	creditCard := "527594"
	paymentSystem := paymentsystem.FindPaymentSystemByPrefix(creditCard)
	if paymentSystem != nil {
		fmt.Printf("Payment system: %s\n", *paymentSystem)
	}
}
