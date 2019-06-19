package main

import (
	"log"
	"sort"

	"github.com/khasanovbi/banks_db/internal/gen"
)

const (
	baseDir       = "banksdb"
	countriesPath = baseDir + "/countries.go"
	banksPath     = baseDir + "/banks.go"
)

func calculateBanksCount(countryBanksSlice []gen.CountryBanks) int {
	sum := 0
	for _, countryBanks := range countryBanksSlice {
		sum += len(countryBanks.Banks)
	}
	return sum
}

func getCountries(banksByCountry []gen.CountryBanks) []string {
	countries := make([]string, 0, len(banksByCountry))
	for _, countryBanks := range banksByCountry {
		countries = append(countries, countryBanks.Country)
	}
	sort.Strings(countries)
	return countries
}

func main() {
	countryBanksSlice := gen.ParseBanks()
	log.Printf(
		"the banks are parsed: countriesCount=%d, banksCount=%d",
		len(countryBanksSlice),
		calculateBanksCount(countryBanksSlice),
	)
	gen.GenerateCountriesFile(countriesPath, getCountries(countryBanksSlice))
	gen.GenerateBanksFile(banksPath, countryBanksSlice)
}
