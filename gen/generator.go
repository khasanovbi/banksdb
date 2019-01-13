package gen

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type CountriesFileParams struct {
	Package   string
	Countries []string
}

type BanksFileParams struct {
	Package           string
	CountryBanksSlice []CountryBanks
}

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
}

var (
	countriesTpl = template.Must(template.New("countriesTpl").Funcs(funcMap).Parse(
		`package {{.Package}}

type Country string

const (
{{range $country := .Countries}}	{{$country | ToUpper}} Country = "{{$country}}"
{{end}})
`))

	banksTpl = template.Must(template.New("banksTpl").Funcs(funcMap).Parse(
		`package {{.Package}}

var banksByCountry = map[Country][]Bank{
{{range $countryBanks := .CountryBanksSlice}}	{{$countryBanks.Country | ToUpper}}: {
{{range $bank := $countryBanks.Banks}}		{
			Name:       "{{$bank.Name}}",
			Country:    "{{$bank.Country}}",
			LocalTitle: "{{$bank.LocalTitle}}",
			EngTitle:   "{{$bank.EngTitle}}",
			URL:        "{{$bank.URL}}",
			Color:      "{{$bank.Color}}",
			Prefixes:   []int{{"{"}}{{range $i, $prefix := $bank.Prefixes}}{{if $i}}, {{end}}{{$prefix}}{{end}}},
		},
{{end}}	},
{{end}}}
`))
)

func getPackageName(outputPath string) string {
	return filepath.Dir(outputPath)
}

func GenerateCountriesFile(outputPath string, countries []string) {
	log.Printf("generate countries file: path='%s'", outputPath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	err = countriesTpl.Execute(
		outputFile,
		CountriesFileParams{
			Package:   getPackageName(outputPath),
			Countries: countries,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file generated")
}

func GenerateBanksFile(outputPath string, countryBanksSlice []CountryBanks) {
	log.Printf("generate banks file: path='%s'", outputPath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	err = banksTpl.Execute(
		outputFile,
		BanksFileParams{
			Package:           getPackageName(outputPath),
			CountryBanksSlice: countryBanksSlice,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file generated")
}
