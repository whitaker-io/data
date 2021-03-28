package data

import (
	"fmt"
)

// Bool provides bool if available or else returns an error
func (d Data) Bool(key string) (bool, error) {
	val, err := d.navigate(key)

	if err != nil {
		return false, err
	}

	var x bool
	var ok bool

	if x, ok = val.(bool); !ok {
		return false, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustBool provides bool from key or panics
func (d Data) MustBool(key string) bool {
	val, err := d.Bool(key)

	if err != nil {
		panic(err)
	}

	return val
}

// BoolOr provides bool from key or defaultVal
func (d Data) BoolOr(key string, defaultVal bool) bool {
	val, err := d.Bool(key)

	if err != nil {
		return defaultVal
	}

	return val
}
