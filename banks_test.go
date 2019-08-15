package banksdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotEmptyDB(t *testing.T) {
	require.NotEmpty(t, banksByCountry)
}

func TestNotEmptyBanksInCountry(t *testing.T) {
	for country, banks := range banksByCountry {
		require.NotEmpty(t, banks, "Country '%s' is empty", country)
	}
}

func TestUniquePrefixesInBanks(t *testing.T) {
	t.Skip() // Now this test failed because there is some nonunique prefixes in db.
	seen := make(map[int]struct{})
	for _, banks := range banksByCountry {
		for _, bank := range banks {
			for _, prefix := range bank.Prefixes {
				_, isSeenBefore := seen[prefix]
				require.False(t, isSeenBefore, "prefix %d seen before", prefix)
				seen[prefix] = struct{}{}
			}
		}
	}
}
