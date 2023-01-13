package banksdb

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NewBanksDBTestSuite struct {
	suite.Suite
	country1 Country
	bank1    *Bank
	country2 Country
	bank2    *Bank
}

func (suite *NewBanksDBTestSuite) getSomeBankPrefix(bank *Bank) string {
	return strconv.Itoa(bank.Prefixes[0])
}

func (suite *NewBanksDBTestSuite) SetupSuite() {
	suite.country1 = RU
	suite.bank1 = &banksByCountry[suite.country1][0]
	suite.country2 = CN
	suite.bank2 = &banksByCountry[suite.country2][0]
}

func (suite *NewBanksDBTestSuite) TestEmptyNewBanksDBTestSuite() {
	emptyBanksDB := NewBanksDB()
	bank := emptyBanksDB.FindBank(suite.getSomeBankPrefix(suite.bank1))
	suite.Require().Nil(bank)
}

func (suite *NewBanksDBTestSuite) TestFindBankInSingleCountry() {
	banksDB := NewBanksDB(suite.country1)
	bank := banksDB.FindBank(suite.getSomeBankPrefix(suite.bank1))
	suite.Require().NotNil(bank)
	suite.Require().Equal(suite.bank1, bank)

	bank = banksDB.FindBank(suite.getSomeBankPrefix(suite.bank2))
	suite.Require().Nil(bank)
}

func TestNewBanksDBTestSuite(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(NewBanksDBTestSuite))
}
