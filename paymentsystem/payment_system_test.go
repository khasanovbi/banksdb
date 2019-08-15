package paymentsystem_test

import (
	"testing"

	"github.com/khasanovbi/banksdb/v2/paymentsystem"
	"github.com/stretchr/testify/require"
)

func TestFindPaymentSystem(t *testing.T) {
	tests := map[paymentsystem.PaymentSystem]string{
		paymentsystem.AmericanExpress:         "341414688673814",
		paymentsystem.Dankort:                 "5019613744152545",
		paymentsystem.DinersClubInternational: "3872481490310852",
		paymentsystem.Discover:                "6011988461284820",
		paymentsystem.InterPayment:            "6365364106904019758",
		paymentsystem.InstaPayment:            "6397686307314238",
		paymentsystem.JCB:                     "3582616798872373",
		paymentsystem.LankaPay:                "3571113212345544",
		paymentsystem.Maestro:                 "5080741549588144561",
		paymentsystem.MasterCard:              "2550053850150029",
		paymentsystem.Mir:                     "2204941877211882",
		paymentsystem.NPSPridnestrovie:        "6054740386428539",
		paymentsystem.RuPay:                   "6522133919284495",
		paymentsystem.Troy:                    "9792376700578340",
		paymentsystem.TUnion:                  "3152259470993676486",
		paymentsystem.UATP:                    "132435418436821",
		paymentsystem.UnionPay:                "6221269639999729",
		paymentsystem.Verve:                   "5061725767660126",
		paymentsystem.Visa:                    "4607322866767830",
	}
	for ps, creditCard := range tests {
		ps := ps
		creditCard := creditCard
		t.Run(string(ps), func(t *testing.T) {
			actualPaymentSystem := paymentsystem.FindPaymentSystem(creditCard)
			require.NotNil(t, actualPaymentSystem)
			require.EqualValues(t, ps, *actualPaymentSystem)
		})
	}
}

func TestEmptyCreditCardFindPaymentSystem(t *testing.T) {
	actualPaymentSystem := paymentsystem.FindPaymentSystem("")
	require.Nil(t, actualPaymentSystem)
}
