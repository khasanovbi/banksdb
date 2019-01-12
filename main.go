package main

import (
	"github.com/khasanovbi/banks_db/gen"
	"log"
)

const GenPath = "banks_db/banks.go"

func main() {
	banks := gen.ParseBanks()
	log.Printf("the banks are parsed: count=%d", len(banks))
	gen.GenerateFile(GenPath, banks)
}
