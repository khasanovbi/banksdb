//nolint:forbidigo
package banksdb_test

import (
	"fmt"

	"github.com/khasanovbi/banksdb/v2"
)

func ExampleFindBank() {
	creditCard := "5275940000000000"
	bank := banksdb.FindBank(creditCard)
	fmt.Printf("Bank info: %#v\n", bank)
}

func ExampleBanksDB_findBank() {
	// Create BanksDB only for Canadian and USA banks.
	db := banksdb.NewBanksDB(banksdb.CA, banksdb.US)
	bank := db.FindBank("5290994338557863")
	fmt.Printf("Bank info: %#v\n", bank)
}
