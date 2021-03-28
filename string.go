package data

import (
	"fmt"
)

// String provides string if available or else returns an error
func (d Data) String(key string) (string, error) {
	val, err := d.navigate(key)

	if err != nil {
		return "", err
	}

	var x string
	var ok bool

	if x, ok = val.(string); !ok {
		return "", fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustString provides string from key or panics
func (d Data) MustString(key string) string {
	val, err := d.String(key)

	if err != nil {
		panic(err)
	}

	return val
}

// StringOr provides string from key or defaultVal
func (d Data) StringOr(key, defaultVal string) string {
	val, err := d.String(key)

	if err != nil {
		return defaultVal
	}

	return val
}
