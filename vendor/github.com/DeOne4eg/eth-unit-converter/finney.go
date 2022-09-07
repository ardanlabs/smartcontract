package unit

import "math/big"

const (
	Finney = 1e3
)

// Finney returns the value in Finney units.
func (u Unit) Finney() *big.Float {
	return u.convertTo(Finney)
}

// NewFinney create instance of Unit with convert Finney to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewFinney(value *big.Float) *Unit {
	return convertFrom(value, Finney)
}
