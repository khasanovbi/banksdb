package main

import (
	"flag"
	"log"
	"path"
	"path/filepath"
	"sort"
)

var baseDir = "../.."

func init() {
	flag.StringVar(&baseDir, "base-dir", "", "repo base directory")
}

func calculateBanksCount(countryBanksSlice []CountryBanks) int {
	sum := 0
	for _, countryBanks := range countryBanksSlice {
		sum += len(countryBanks.Banks)
	}

	return sum
}

func getCountries(banksByCountry []CountryBanks) []string {
	countries := make([]string, 0, len(banksByCountry))
	for _, countryBanks := range banksByCountry {
		countries = append(countries, countryBanks.Country)
	}

	sort.Strings(countries)

	return countries
}

func main() {
	flag.Parse()

	countryBanksSlice := ParseBanks()
	log.Printf(
		"the banks are parsed: countriesCount=%d, banksCount=%d",
		len(countryBanksSlice),
		calculateBanksCount(countryBanksSlice),
	)

	countriesPath, err := filepath.Abs(path.Join(baseDir, "countries.go"))
	if err != nil {
		log.Fatal(err)
	}

	banksPath, err := filepath.Abs(path.Join(baseDir, "banks.go"))
	if err != nil {
		log.Fatal(err)
	}

	GenerateCountriesFile(countriesPath, getCountries(countryBanksSlice))
	GenerateBanksFile(banksPath, countryBanksSlice)
}
