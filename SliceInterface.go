package data

import (
	"fmt"
)

// SliceInterface provides []interface{} if available or else returns an error
func (d Data) SliceInterface(key string) ([]interface{}, error) {
	val, err := d.navigate(key)

	if err != nil {
		return nil, err
	}

	var x []interface{}
	var ok bool

	if x, ok = val.([]interface{}); !ok {
		return nil, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustSliceInterface provides []interface{} from key or panics
func (d Data) MustSliceInterface(key string) []interface{} {
	val, err := d.SliceInterface(key)

	if err != nil {
		panic(err)
	}

	return val
}

// SliceInterfaceOr provides []interface{} from key or defaultVal
func (d Data) SliceInterfaceOr(key string, defaultVal []interface{}) []interface{} {
	val, err := d.SliceInterface(key)

	if err != nil {
		return defaultVal
	}

	return val
}
