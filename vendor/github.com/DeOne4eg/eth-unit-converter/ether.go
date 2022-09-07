package unit

import (
	"math/big"
)

const (
	Ether = 1
)

// Ether returns the value in Ether units.
func (u Unit) Ether() float64 {
	return float64(u.Value.Int64()) / Wei
}

// NewEther create instance of Unit with convert Ether to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewEther(value *big.Float) *Unit {
	return convertFrom(value, Ether)
}
