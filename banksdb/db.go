package banksdb

import "strconv"

// BanksDB is an interface representing the ability to find bank info by creditCard number.
type BanksDB interface {
	FindBank(creditCard string) *Bank
}

type banksDBImpl struct {
	prefixToBank map[int]*Bank
}

func (b *banksDBImpl) FindBank(creditCard string) *Bank {
	// NOTE: Start in reverse order to make less lookups
	for _, prefixLength := range []int{6, 5} {
		prefix, err := strconv.Atoi(creditCard[:prefixLength])
		if err != nil {
			return nil
		}
		if bank, ok := b.prefixToBank[prefix]; ok {
			return bank
		}
	}
	return nil
}

func (b *banksDBImpl) addBanksToDB(banks []Bank) {
	for i := range banks {
		bank := &banks[i]
		for _, prefix := range bank.Prefixes {
			b.prefixToBank[prefix] = bank
		}
	}
}

// NewBanksDB creates BanksDB for given countries.
func NewBanksDB(countries ...Country) BanksDB {
	banksDB := &banksDBImpl{
		prefixToBank: make(map[int]*Bank),
	}
	for _, country := range countries {
		banks := banksByCountry[country]
		banksDB.addBanksToDB(banks)
	}
	return banksDB
}
