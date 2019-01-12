package banks_db

import "strconv"

type Bank struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	LocalTitle string `json:"localTitle"`
	EngTitle   string `json:"engTitle"`
	URL        string `json:"url"`
	Color      string `json:"color"`
	Prefixes   []int  `json:"prefixes"`
}

func FindBank(creditCard string) *Bank {
	for _, prefixLength := range []int{5, 6} {
		prefix, err := strconv.Atoi(creditCard[:prefixLength])
		if err != nil {
			return nil
		}
		if bank, ok := prefixToBank[prefix]; ok {
			return bank
		}
	}
	return nil
}
