# Banks DB

[![Build Status](https://travis-ci.org/khasanovbi/banksdb.svg?branch=master)](https://travis-ci.org/khasanovbi/banksdb)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/khasanovbi/banksdb)
[![Go Report Card](https://goreportcard.com/badge/github.com/khasanovbi/banksdb?style=flat-square)](https://goreportcard.com/report/github.com/khasanovbi/banksdb)
[![Release](https://img.shields.io/github/release/khasanovbi/banksdb.svg?style=flat-square)](https://github.com/khasanovbi/banksdb/releases/latest)

Community driven database to get bank info (name, brand, color, etc.) by bankcard prefix (BIN)

> This is golang port of [ramoona's banks-db](https://github.com/ramoona/banks-db).

### Install

```
go get -u github.com/KhasanovBI/banksdb/banksdb
```

### Usage

Below is an example which shows some common use cases for banksdb:

```go
package main

import (
	"fmt"

	"github.com/khasanovbi/banksdb/banksdb"
)

func main() {
	for _, creditCard := range []string{"5275940000000000", "4111111111111111"} {
		bank := banksdb.FindBank(creditCard)
		paymentSystem := banksdb.FindPaymentSystem(creditCard)
		fmt.Printf("CreditCard: %s\n", creditCard)
		fmt.Printf("Bank info: %#v\n", bank)
		if paymentSystem != nil {
			fmt.Printf("Payment system: %s\n", *paymentSystem)
		}
		fmt.Println()
	}
}

```

Output:
```
CreditCard: 5275940000000000
Bank info: &banksdb.Bank{Name:"citibank", Country:"ru", LocalTitle:"Ситибанк", EngTitle:"Citibank", URL:"https://www.citibank.ru/", Color:"#0088cf", Prefixes:[]int{419349, 427760, 427761, 520306, 527594}}
Payment system: mastercard

CreditCard: 4111111111111111
Bank info: (*banksdb.Bank)(nil)
Payment system: visa
```