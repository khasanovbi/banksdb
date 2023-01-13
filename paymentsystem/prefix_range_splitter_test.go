package paymentsystem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitPrefixRange(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		prefixRange      prefixRange
		expectedPrefixes []int
	}{
		"middle related prefix": {
			prefixRange{from: 10058, to: 10071},
			[]int{10058, 10059, 1006, 10070, 10071},
		},
		"same from to": {
			prefixRange{from: 10058, to: 10058},
			[]int{10058},
		},
		"boundary near prefix": {
			prefixRange{from: 10000, to: 10120},
			[]int{100, 1010, 1011, 10120},
		},
	}

	for testName, testParams := range tests {
		testParams := testParams

		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			result, err := splitPrefixRange(testParams.prefixRange)
			require.NoError(t, err)
			require.ElementsMatchf(t, testParams.expectedPrefixes, result, "%v", result)
		})
	}
}
