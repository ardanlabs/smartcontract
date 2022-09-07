package unit

import (
	"math/big"
)

const (
	KWei = 1e15
)

// KWei returns the value in KWei units.
func (u Unit) KWei() *big.Float {
	return u.convertTo(KWei)
}

// NewKWei create instance of Unit with convert KWei to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewKWei(value *big.Float) *Unit {
	return convertFrom(value, KWei)
}
