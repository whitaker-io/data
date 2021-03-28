package data

import (
	"fmt"
)

// Uint provides uint if available or else returns an error
func (d Data) Uint(key string) (uint, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x uint
	var ok bool

	if x, ok = val.(uint); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustUint provides uint from key or panics
func (d Data) MustUint(key string) uint {
	val, err := d.Uint(key)

	if err != nil {
		panic(err)
	}

	return val
}

// UintOr provides uint from key or defaultVal
func (d Data) UintOr(key string, defaultVal uint) uint {
	val, err := d.Uint(key)

	if err != nil {
		return defaultVal
	}

	return val
}
