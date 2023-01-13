package banksdb_test

import (
	"testing"

	"github.com/khasanovbi/banksdb/v2"
	"github.com/stretchr/testify/require"
)

func TestFindBankEmptyCreditCard(t *testing.T) {
	t.Parallel()

	bank := banksdb.FindBank("")
	require.Nil(t, bank)
}

func TestFindBankInvalidCreditCard(t *testing.T) {
	t.Parallel()

	bank := banksdb.FindBank("no-digits")
	require.Nil(t, bank)
}
