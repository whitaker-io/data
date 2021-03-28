package data

import (
	"fmt"
)

// Int provides int if available or else returns an error
func (d Data) Int(key string) (int, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x int
	var ok bool

	if x, ok = val.(int); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustInt provides int from key or panics
func (d Data) MustInt(key string) int {
	val, err := d.Int(key)

	if err != nil {
		panic(err)
	}

	return val
}

// IntOr provides int from key or defaultVal
func (d Data) IntOr(key string, defaultVal int) int {
	val, err := d.Int(key)

	if err != nil {
		return defaultVal
	}

	return val
}
