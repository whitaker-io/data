package data

import (
	"fmt"
)

// Uint64 provides uint64 if available or else returns an error
func (d Data) Uint64(key string) (uint64, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x uint64
	var ok bool

	if x, ok = val.(uint64); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustUint64 provides uint64 from key or panics
func (d Data) MustUint64(key string) uint64 {
	val, err := d.Uint64(key)

	if err != nil {
		panic(err)
	}

	return val
}

// Uint64Or provides uint64 from key or defaultVal
func (d Data) Uint64Or(key string, defaultVal uint64) uint64 {
	val, err := d.Uint64(key)

	if err != nil {
		return defaultVal
	}

	return val
}
