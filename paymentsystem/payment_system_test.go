//nolint: gomnd
package paymentsystem

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestFindPaymentSystem(t *testing.T) {
	tests := map[PaymentSystem]string{
		AmericanExpress:         "341414688673814",
		Dankort:                 "5019613744152545",
		DinersClubInternational: "3872481490310852",
		Discover:                "6011988461284820",
		InterPayment:            "6365364106904019758",
		InstaPayment:            "6397686307314238",
		JCB:                     "3582616798872373",
		LankaPay:                "3571113212345544",
		Maestro:                 "5080741549588144561",
		MasterCard:              "2550053850150029",
		Mir:                     "2204941877211882",
		NPSPridnestrovie:        "6054740386428539",
		RuPay:                   "6522133919284495",
		Troy:                    "9792376700578340",
		TUnion:                  "3152259470993676486",
		UATP:                    "132435418436821",
		UnionPay:                "6221269639999729",
		Verve:                   "5061725767660126",
		Visa:                    "4607322866767830",
	}
	for ps, creditCard := range tests {
		ps := ps
		creditCard := creditCard

		t.Run(string(ps), func(t *testing.T) {
			actualPaymentSystem := FindPaymentSystem(creditCard)
			require.NotNil(t, actualPaymentSystem)
			require.EqualValues(t, ps, *actualPaymentSystem)
		})
	}
}

func TestEmptyCreditCardFindPaymentSystem(t *testing.T) {
	actualPaymentSystem := FindPaymentSystem("")
	require.Nil(t, actualPaymentSystem)
}

type RadixDBTestSuite struct {
	suite.Suite
}

func (suite *RadixDBTestSuite) TestNotLongestPrefix() {
	db := newRadixDB()
	EqualByPrefixPs := PaymentSystem("EqualByPrefixPs")
	EqualByLengthAndPrefixPs := PaymentSystem("EqualByLengthAndPrefixPs")
	creditCard := "12"
	err := db.InitFromMap(map[PaymentSystem][]paymentSystemInfo{
		EqualByLengthAndPrefixPs: {
			{prefixes: []int{1}, lengthChecker: &exactLengthChecker{Exact: 2}},
		},
		EqualByPrefixPs: {
			{prefixes: []int{12}, lengthChecker: &exactLengthChecker{Exact: 5}},
		},
	})

	suite.Require().NoError(err)

	actualPs := db.FindPaymentSystem(creditCard)
	suite.Require().NotNil(actualPs)
	suite.Require().Equal(EqualByLengthAndPrefixPs, *actualPs)
	actualPs = db.FindPaymentSystemByPrefix(creditCard)
	suite.Require().NotNil(actualPs)
	suite.Require().Equal(EqualByPrefixPs, *actualPs)
}

func (suite *RadixDBTestSuite) TestInitErrorAtSamePrefix() {
	db := newRadixDB()
	ps1 := PaymentSystem("ps1")
	ps2 := PaymentSystem("ps2")
	err := db.InitFromMap(map[PaymentSystem][]paymentSystemInfo{
		ps1: {
			{prefixes: []int{1}, lengthChecker: &exactLengthChecker{Exact: 2}},
		},
		ps2: {
			{prefixes: []int{1}, lengthChecker: &exactLengthChecker{Exact: 5}},
		},
	})
	suite.Require().True(errors.Is(err, errNonUniquePrefix))
}

func TestRadixDBTestSuite(t *testing.T) {
	suite.Run(t, new(RadixDBTestSuite))
}
