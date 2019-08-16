package banksdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindBankEmptyCreditCard(t *testing.T) {
	bank := FindBank("")
	require.Nil(t, bank)
}
