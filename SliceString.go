package data

import (
	"fmt"
)

// SliceString provides []string if available or else returns an error
func (d Data) SliceString(key string) ([]string, error) {
	val, err := d.navigate(key)

	if err != nil {
		return nil, err
	}

	var x []string
	var ok bool

	if x, ok = val.([]string); !ok {
		return nil, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustSliceString provides []string from key or panics
func (d Data) MustSliceString(key string) []string {
	val, err := d.SliceString(key)

	if err != nil {
		panic(err)
	}

	return val
}

// SliceStringOr provides []string from key or defaultVal
func (d Data) SliceStringOr(key string, defaultVal []string) []string {
	val, err := d.SliceString(key)

	if err != nil {
		return defaultVal
	}

	return val
}
