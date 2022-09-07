package unit

import "math/big"

const (
	Szabo = 1e6
)

// Szabo returns the value in Szabo units.
func (u Unit) Szabo() *big.Float {
	return u.convertTo(Szabo)
}

// NewSzabo create instance of Unit with convert Szabo to Wei units and returns pointer to it.
// Then you can use Unit for get the value in supported units.
func NewSzabo(value *big.Float) *Unit {
	return convertFrom(value, Szabo)
}
