package banksdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindBankEmptyCreditCard(t *testing.T) {
	bank := FindBank("")
	require.Nil(t, bank)
}

func TestFindBankInvalidCreditCard(t *testing.T) {
	bank := FindBank("no-digits")
	require.Nil(t, bank)
}
