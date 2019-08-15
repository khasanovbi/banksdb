package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type countriesFileParams struct {
	Package   string
	Countries []string
}

type banksFileParams struct {
	Package           string
	CountryBanksSlice []CountryBanks
}

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

var (
	countriesTpl = template.Must(template.New("countriesTpl").Funcs(funcMap).Parse(
		`/*
* CODE GENERATED AUTOMATICALLY WITH github.com/khasanovbi/banksdb/internal/gen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package {{.Package}}

// Country represent country code.
type Country string

// Following constants represent country codes of known banks in db.
const (
{{range $country := .Countries}}	{{$country | ToUpper}} Country = "{{$country}}"
{{end}})
`))

	banksTpl = template.Must(template.New("banksTpl").Funcs(funcMap).Parse(
		`/*
* CODE GENERATED AUTOMATICALLY WITH github.com/khasanovbi/banksdb/internal/gen
* THIS FILE MUST NOT BE EDITED BY HAND
 */

package {{.Package}}

var banksByCountry = map[Country][]Bank{
{{range $countryBanks := .CountryBanksSlice}}	{{$countryBanks.Country | ToUpper}}: {
{{range $bank := $countryBanks.Banks}}		{
			Name:       "{{$bank.Name}}",
			Country:    "{{$bank.Country}}",
			LocalTitle: "{{$bank.LocalTitle}}",
			EngTitle:   "{{$bank.EngTitle}}",
			URL:        "{{$bank.URL}}",
			Color:      "{{$bank.Color}}",
			Prefixes: []int{{"{"}}{{range $i, $prefix := $bank.Prefixes}}
				{{$prefix}},{{end}}
			},
		},
{{end}}	},
{{end}}}
`))
)

func getPackageName(outputPath string) string {
	return filepath.Base(filepath.Dir(outputPath))
}

// GenerateCountriesFile generate go file with countries.
func GenerateCountriesFile(outputPath string, countries []string) {
	log.Printf("generate countries file: path='%s'", outputPath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	err = countriesTpl.Execute(
		outputFile,
		countriesFileParams{
			Package:   getPackageName(outputPath),
			Countries: countries,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file generated")
}

// GenerateBanksFile generate go file with country to bank mapping.
func GenerateBanksFile(outputPath string, countryBanksSlice []CountryBanks) {
	log.Printf("generate banks file: path='%s'", outputPath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	err = banksTpl.Execute(
		outputFile,
		banksFileParams{
			Package:           getPackageName(outputPath),
			CountryBanksSlice: countryBanksSlice,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file generated")
}
