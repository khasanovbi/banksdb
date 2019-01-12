package gen

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type Params struct {
	Package string
	Banks   []Bank
}

var (
	banksTpl = template.Must(template.New("banksTpl").Parse(
		`package {{.Package}}

var prefixToBank = make(map[int]*Bank)
var banks = []Bank{
{{range $bank := .Banks}}	{
		Name:       "{{$bank.Name}}",
		Country:    "{{$bank.Country}}",
		LocalTitle: "{{$bank.LocalTitle}}",
		EngTitle:   "{{$bank.EngTitle}}",
		URL:        "{{$bank.URL}}",
		Color:      "{{$bank.Color}}",
		Prefixes:   []int{{"{"}}{{range $i, $prefix := $bank.Prefixes}}{{if $i}}, {{end}}{{$prefix}}{{end}}},
	},
{{end}}}

func init() {
	for i := range banks {
		bank := &banks[i]
		for _, prefix := range bank.Prefixes {
			prefixToBank[prefix] = bank
		}
	}
}
`))
)

func getPackageName(outputPath string) string {
	return filepath.Dir(outputPath)
}

func GenerateFile(outputPath string, banks []Bank) {
	log.Printf("generate file: path='%s'", outputPath)
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}
	err = banksTpl.Execute(
		outputFile,
		Params{
			Package: getPackageName(outputPath),
			Banks:   banks,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("file generated")
}
