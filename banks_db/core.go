package banks_db

var banksDB = &banksDBImpl{
	prefixToBank: make(map[int]*Bank),
}

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
	return banksDB.FindBank(creditCard)
}

func init() {
	for _, banks := range banksByCountry {
		banksDB.addBanksToDB(banks)
	}
}
