package paymentsystem_test

import (
	"regexp"
	"testing"

	"github.com/khasanovbi/banksdb/v2/paymentsystem"
)

type singleRecordRePaymentSystemDB struct {
	psFullLengthRe *regexp.Regexp
	psPrefixRe     *regexp.Regexp
	ps             *paymentsystem.PaymentSystem
}

func (s *singleRecordRePaymentSystemDB) FindPaymentSystem(
	creditCard string,
) (paymentSystem *paymentsystem.PaymentSystem) {
	if s.psFullLengthRe.MatchString(creditCard) {
		return s.ps
	}

	return nil
}

func (s *singleRecordRePaymentSystemDB) FindPaymentSystemByPrefix(
	creditCardPrefix string,
) (paymentSystem *paymentsystem.PaymentSystem) {
	if s.psPrefixRe.MatchString(creditCardPrefix) {
		return s.ps
	}

	return nil
}

func newPaymentSystemDBRe() paymentsystem.DB {
	ps := paymentsystem.UnionPay

	return &singleRecordRePaymentSystemDB{
		psFullLengthRe: regexp.MustCompile(`^622126\d{10}$`),
		psPrefixRe:     regexp.MustCompile(`^622126\d*$`),
		ps:             &ps,
	}
}

func BenchmarkFindPaymentSystem(b *testing.B) {
	// Use 2 level in radix card. 6 - Maestro, 622126 - UnionPay
	paymentSystemDBRadix := paymentsystem.NewDB()
	// Use simple regexp that cover one case
	paymentSystemDBRe := newPaymentSystemDBRe()

	paymentSystemDBs := map[string]paymentsystem.DB{
		"Radix": paymentSystemDBRadix,
		"Re":    paymentSystemDBRe,
	}
	creditCard := "6221269639999729"

	for name, paymentSystemDB := range paymentSystemDBs {
		paymentSystemDB := paymentSystemDB

		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				paymentSystemDB.FindPaymentSystem(creditCard)
			}
		})
	}
}

func BenchmarkFindPaymentSystemByPrefix(b *testing.B) {
	// Use 2 level in radix card. 6 - Maestro, 622126 - UnionPay
	paymentSystemDBRadix := paymentsystem.NewDB()
	// Use simple regexp that cover one case
	paymentSystemDBRe := newPaymentSystemDBRe()

	paymentSystemDBs := map[string]paymentsystem.DB{
		"Radix": paymentSystemDBRadix,
		"Re":    paymentSystemDBRe,
	}
	creditCard := "6221269639999729"

	for name, paymentSystemDB := range paymentSystemDBs {
		paymentSystemDB := paymentSystemDB

		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				paymentSystemDB.FindPaymentSystemByPrefix(creditCard)
			}
		})
	}
}
