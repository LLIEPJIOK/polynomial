package polynomial

import "fmt"

type ErrUnknownOp struct {
	op string
}

func NewErrUnknownOp(op string) error {
	return ErrUnknownOp{
		op: op,
	}
}

func (e ErrUnknownOp) Error() string {
	return fmt.Sprintf("unknown operation %q", e.op)
}
