package unit

import "math/big"

const (
	MWei = 1e12
)

// MWei returns the value in MWei units.
func (u Unit) MWei() *big.Float {
	return u.convertTo(MWei)
}

// NewMWei create instance of Unit with convert MWei to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewMWei(value *big.Float) *Unit {
	return convertFrom(value, MWei)
}
