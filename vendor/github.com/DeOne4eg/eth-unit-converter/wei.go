package unit

import "math/big"

const (
	Wei = 1e18
)

func NewWei(value *big.Int) *Unit {
	return &Unit{Value: value}
}

func (u Unit) Wei() *big.Int {
	return u.Value
}
