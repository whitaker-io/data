package data

import (
	"fmt"
)

// MapStringString provides map[string]string if available or else returns an error
func (d Data) MapStringString(key string) (map[string]string, error) {
	val, err := d.navigate(key)

	if err != nil {
		return nil, err
	}

	var x map[string]string
	var ok bool

	if x, ok = val.(map[string]string); !ok {
		return nil, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustMapStringString provides map[string]string from key or panics
func (d Data) MustMapStringString(key string) map[string]string {
	val, err := d.MapStringString(key)

	if err != nil {
		panic(err)
	}

	return val
}

// MapStringStringOr provides map[string]string from key or defaultVal
func (d Data) MapStringStringOr(key string, defaultVal map[string]string) map[string]string {
	val, err := d.MapStringString(key)

	if err != nil {
		return defaultVal
	}

	return val
}
