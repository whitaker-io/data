package data

import (
	"fmt"
)

// Int64 provides int64 if available or else returns an error
func (d Data) Int64(key string) (int64, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x int64
	var ok bool

	if x, ok = val.(int64); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustInt64 provides int64 from key or panics
func (d Data) MustInt64(key string) int64 {
	val, err := d.Int64(key)

	if err != nil {
		panic(err)
	}

	return val
}

// Int64Or provides int64 from key or defaultVal
func (d Data) Int64Or(key string, defaultVal int64) int64 {
	val, err := d.Int64(key)

	if err != nil {
		return defaultVal
	}

	return val
}
