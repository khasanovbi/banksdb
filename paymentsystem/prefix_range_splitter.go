package paymentsystem

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var errDifferentPrefixRangeLengths = errors.New("different prefix range lengths")

func findCommonPrefixLength(first string, second string) int {
	for commonPrefixLength := 0; commonPrefixLength < len(first); commonPrefixLength++ {
		if first[commonPrefixLength] != second[commonPrefixLength] {
			return commonPrefixLength
		}
	}

	return len(first)
}

func isOnlyChar(str string, char uint8) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != char {
			return false
		}
	}

	return true
}

func isOnlyZeros(str string) bool {
	return isOnlyChar(str, '0')
}

func isOnlyNines(str string) bool {
	return isOnlyChar(str, '9')
}

func appendChars(str string, char uint8, totalSize int) string {
	if count := totalSize - len(str); count > 0 {
		str += strings.Repeat(string(char), count)
	}

	return str
}

func appendNines(str string, totalSize int) string {
	return appendChars(str, '9', totalSize)
}

func appendZeros(str string, totalSize int) string {
	return appendChars(str, '0', totalSize)
}

func splitPrefixRangeStr(fromStr string, toStr string) (prefixes []string) {
	commonPrefixLength := findCommonPrefixLength(fromStr, toStr)
	commonPrefix := fromStr[:commonPrefixLength]

	fromStr = fromStr[commonPrefixLength:]
	toStr = toStr[commonPrefixLength:]

	if len(fromStr) == 0 {
		return []string{commonPrefix}
	}

	startChar := fromStr[0]
	endChar := toStr[0]

	if len(fromStr) > 1 && !isOnlyZeros(fromStr[1:]) {
		startChar++

		prefixes = append(prefixes, splitPrefixRangeStr(fromStr, appendNines(fromStr[:1], len(fromStr)))...)
	}

	if len(toStr) > 1 && !isOnlyNines(toStr[1:]) {
		endChar--

		prefixes = append(prefixes, splitPrefixRangeStr(appendZeros(toStr[:1], len(toStr)), toStr)...)
	}

	for char := startChar; char <= endChar; char++ {
		prefixes = append(prefixes, string(char))
		// prevent overflow?
		if char == math.MaxUint8 {
			break
		}
	}

	for i := range prefixes {
		prefixes[i] = commonPrefix + prefixes[i]
	}

	return prefixes
}

func splitPrefixRange(prefixRange prefixRange) ([]int, error) {
	fromStr := strconv.Itoa(prefixRange.from)
	toStr := strconv.Itoa(prefixRange.to)

	if len(fromStr) != len(toStr) {
		return nil, fmt.Errorf("%w, from='%s', to='%s'", errDifferentPrefixRangeLengths, fromStr, toStr)
	}

	strPrefixes := splitPrefixRangeStr(fromStr, toStr)

	result := make([]int, 0, len(strPrefixes))

	for _, strPrefix := range strPrefixes {
		prefix, err := strconv.Atoi(strPrefix)
		// Shouldn't be never not nil.
		if err != nil {
			return nil, fmt.Errorf("prefix to int convert error: %w", err)
		}

		result = append(result, prefix)
	}

	return result, nil
}
