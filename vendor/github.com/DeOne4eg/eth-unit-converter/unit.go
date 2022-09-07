package unit

import (
	"errors"
	"math/big"
	"strings"
)

var (
	// precision required for big.Float.
	// Use SetPrecision to change the value.
	precision uint = 1024
)

// Unit store value in Wei unit.
// Wei should be always big.Int.
type Unit struct {
	Value *big.Int
}

// String return value of Unit in string format.
func (u Unit) String() string {
	return u.Value.String()
}

// convertTo transform Unit.Value (always wei) to another unit and return pointer to big.Float.
func (u Unit) convertTo(unit float64) *big.Float {
	val := newBigFloat().SetInt(u.Value)
	diff := newBigFloat().Quo(big.NewFloat(Wei), big.NewFloat(unit))

	return newBigFloat().Quo(val, diff)
}

// convertFrom transform value in passed unit.
// Return pointer to Unit.
func convertFrom(value *big.Float, unit float64) *Unit {
	result := new(big.Int)
	newBigFloat().Mul(value, big.NewFloat(Wei/unit)).Int(result)

	return &Unit{Value: result}
}

// SetPrecision set the precision value for big.Float.
// Decrease value for faster conversion.
// Increase value if you need to work with very small or very large numbers.
// It is very easy to understand that you need to increase this value.
// For example: if you convert 1e1000 number and receive not zeros in the end of number then increase value.
func SetPrecision(prec uint) {
	precision = prec
}

// ParseUnit create instance of Unit from passed value and unit name.
// Support units: wei, kwei, mwei, gwei, szabo, finney, ether.
// You may use it for convert units between themselves.
func ParseUnit(value *big.Float, unit string) (*Unit, error) {
	unit = strings.ToLower(unit)

	switch unit {
	case "ether":
		return NewEther(value), nil
	case "wei":
		intVal := new(big.Int)
		value.Int(intVal)
		return NewWei(intVal), nil
	case "kwei":
		return NewKWei(value), nil
	case "mwei":
		return NewMWei(value), nil
	case "gwei":
		return NewGWei(value), nil
	case "finney":
		return NewFinney(value), nil
	case "szabo":
		return NewSzabo(value), nil
	default:
		return nil, errors.New("unknown unit name")
	}
}

// newBigFloat create new instance of big.Float with default properties.
func newBigFloat() *big.Float {
	return new(big.Float).SetPrec(precision)
}
