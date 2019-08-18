package paymentsystem

import (
	"errors"
	"strconv"

	"github.com/armon/go-radix"
	"golang.org/x/xerrors"
)

// PaymentSystem represent system used to settle financial transactions.
type PaymentSystem string

type prefixRange struct {
	from int
	to   int
}

type paymentSystemInfo struct {
	prefixes      []int
	prefixRanges  []prefixRange
	lengthChecker lengthChecker
}

var db DB

type radixValue struct {
	paymentSystem PaymentSystem
	lengthChecker lengthChecker
}

// DB represent interface to find payment system by credit card number or by its prefix.
type DB interface {
	FindPaymentSystem(creditCard string) (paymentSystem *PaymentSystem)
	FindPaymentSystemByPrefix(creditCardPrefix string) (paymentSystem *PaymentSystem)
}

func getPaymentSystemFromValue(v interface{}, creditCardLength int, ignoreLengthCheck bool) *PaymentSystem {
	value := v.(*radixValue)
	if ignoreLengthCheck || value.lengthChecker.CheckLength(creditCardLength) {
		return &value.paymentSystem
	}
	return nil
}

var errNonUniquePrefix = errors.New("non unique prefix")

type radixDB struct {
	tree *radix.Tree
}

func (r *radixDB) findPaymentSystem(creditCard string, ignoreLengthCheck bool) (paymentSystem *PaymentSystem) {
	creditCardLength := len(creditCard)
	prefix, value, ok := r.tree.LongestPrefix(creditCard)
	if !ok {
		return nil
	}
	// Optimization to check the longest prefix first
	paymentSystem = getPaymentSystemFromValue(value, creditCardLength, ignoreLengthCheck)
	if paymentSystem != nil {
		return paymentSystem
	}
	r.tree.WalkPath(prefix, func(s string, v interface{}) bool {
		currentPaymentSystem := getPaymentSystemFromValue(v, creditCardLength, ignoreLengthCheck)
		if currentPaymentSystem != nil {
			paymentSystem = currentPaymentSystem
		}
		return false
	})
	return
}

func (r *radixDB) FindPaymentSystem(creditCard string) (paymentSystem *PaymentSystem) {
	return r.findPaymentSystem(creditCard, false)
}

func (r *radixDB) FindPaymentSystemByPrefix(creditCardPrefix string) (paymentSystem *PaymentSystem) {
	return r.findPaymentSystem(creditCardPrefix, true)
}

func (r *radixDB) InitFromMap(rawPaymentSystems map[PaymentSystem][]paymentSystemInfo) error {
	for paymentSystem, paymentSystemParams := range rawPaymentSystems {
		for i := range paymentSystemParams {
			paymentSystemParam := paymentSystemParams[i]
			prefixes := make([]int, 0, len(paymentSystemParam.prefixRanges)+len(paymentSystemParam.prefixes))
			prefixes = append(prefixes, paymentSystemParam.prefixes...)
			for _, prefixRange := range paymentSystemParam.prefixRanges {
				rangePrefixes, err := splitPrefixRange(prefixRange)
				if err != nil {
					return xerrors.Errorf("prefix range split error: %w", err)
				}
				prefixes = append(prefixes, rangePrefixes...)
			}
			for _, prefix := range prefixes {
				newValue := &radixValue{paymentSystem: paymentSystem, lengthChecker: paymentSystemParam.lengthChecker}
				oldValue, isUpdated := r.tree.Insert(strconv.Itoa(prefix), newValue)
				if isUpdated {
					oldPaymentSystem := oldValue.(*radixValue).paymentSystem
					return xerrors.Errorf(
						"prefix=%d, old=%s, new=%s: %w",
						prefix,
						oldPaymentSystem,
						newValue.paymentSystem,
						errNonUniquePrefix,
					)
				}
			}
		}
	}
	return nil
}

func newRadixDB() *radixDB {
	return &radixDB{tree: radix.New()}
}

// NewDB creates instance of payment system DB.
func NewDB() DB {
	db := newRadixDB()
	err := db.InitFromMap(rawPaymentSystems)
	if err != nil {
		panic(err)
	}
	return db
}

// FindPaymentSystem returns payment system of given credit card.
func FindPaymentSystem(creditCard string) (paymentSystem *PaymentSystem) {
	return db.FindPaymentSystem(creditCard)
}

// FindPaymentSystemByPrefix returns payment system by credit card prefix.
// Similar to FindPaymentSystem, but finds the payment system with the longest prefix, ignoring the length of the card.
func FindPaymentSystemByPrefix(creditCard string) (paymentSystem *PaymentSystem) {
	return db.FindPaymentSystemByPrefix(creditCard)
}

func init() {
	db = NewDB()
}
