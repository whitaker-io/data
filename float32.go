package data

import (
	"fmt"
)

// Float32 provides float32 if available or else returns an error
func (d Data) Float32(key string) (float32, error) {
	val, err := d.navigate(key)

	if err != nil {
		return 0, err
	}

	var x float32
	var ok bool

	if x, ok = val.(float32); !ok {
		return 0, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustFloat32 provides float32 from key or panics
func (d Data) MustFloat32(key string) float32 {
	val, err := d.Float32(key)

	if err != nil {
		panic(err)
	}

	return val
}

// Float32Or provides float32 from key or defaultVal
func (d Data) Float32Or(key string, defaultVal float32) float32 {
	val, err := d.Float32(key)

	if err != nil {
		return defaultVal
	}

	return val
}
