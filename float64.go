package data

import (
	"fmt"
)

// Float64 provides float64 if available or else returns an error
func (d Data) Float64(key string) (float64, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x float64
	var ok bool

	if x, ok = val.(float64); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustFloat64 provides float64 from key or panics
func (d Data) MustFloat64(key string) float64 {
	val, err := d.Float64(key)

	if err != nil {
		panic(err)
	}

	return val
}

// Float64Or provides float64 from key or defaultVal
func (d Data) Float64Or(key string, defaultVal float64) float64 {
	val, err := d.Float64(key)

	if err != nil {
		return defaultVal
	}

	return val
}
