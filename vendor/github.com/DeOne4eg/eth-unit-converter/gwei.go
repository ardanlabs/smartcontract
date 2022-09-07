package unit

import "math/big"

const (
	GWei = 1e9
)

// GWei returns the value in GWei units.
func (u Unit) GWei() *big.Float {
	return u.convertTo(GWei)
}

// NewGWei create instance of Unit with convert GWei to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewGWei(value *big.Float) *Unit {
	return convertFrom(value, GWei)
}
