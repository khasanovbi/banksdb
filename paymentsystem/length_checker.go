package paymentsystem

type lengthChecker interface {
	CheckLength(length int) bool
}

type exactLengthChecker struct {
	Exact int
}

func (e *exactLengthChecker) CheckLength(length int) bool {
	return e.Exact == length
}

type rangeLengthChecker struct {
	From int
	To   int
}

func (r *rangeLengthChecker) CheckLength(length int) bool {
	return r.From <= length && length <= r.To
}

type oneOfLengthChecker []int

func (o oneOfLengthChecker) CheckLength(length int) bool {
	for _, l := range o {
		if length == l {
			return true
		}
	}

	return false
}
