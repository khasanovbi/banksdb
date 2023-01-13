package banksdb_test

import (
	"fmt"

	"github.com/khasanovbi/banksdb/v2"
)

func ExampleFindBank() {
	creditCard := "5275940000000000"
	bank := banksdb.FindBank(creditCard)
	fmt.Printf("Bank info: %#v\n", bank)

	// output:
	// Bank info: &banksdb.Bank{Name:"citibank", Country:"ru", LocalTitle:"Ситибанк", EngTitle:"Citibank", URL:"https://www.citibank.ru/", Color:"#0088cf", Prefixes:[]int{419349, 427760, 427761, 520306, 527594}}
}

func ExampleBanksDB_FindBank() { //nolint:nosnakecase
	// Create BanksDB only for Canadian and USA banks.
	db := banksdb.NewBanksDB(banksdb.CA, banksdb.US)
	bank := db.FindBank("5290994338557863")
	fmt.Printf("Bank info: %#v\n", bank)

	// output:
	// Bank info: &banksdb.Bank{Name:"ally", Country:"us", LocalTitle:"Ally", EngTitle:"Ally", URL:"https://www.ally.com", Color:"#650360", Prefixes:[]int{529099, 557552}}
}
