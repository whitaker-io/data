package data

import (
	"fmt"
)

// MapStringInterface provides map[string]interface{} if available or else returns an error
func (d Data) MapStringInterface(key string) (map[string]interface{}, error) {
	val, err := d.navigate(key)

	if err != nil {
		return nil, err
	}

	var x map[string]interface{}
	var ok bool

	if x, ok = val.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid type for key")
	}

	return x, nil
}

// MustMapStringInterface provides map[string]interface{} from key or panics
func (d Data) MustMapStringInterface(key string) map[string]interface{} {
	val, err := d.MapStringInterface(key)

	if err != nil {
		panic(err)
	}

	return val
}

// MapStringInterfaceOr provides map[string]interface{} from key or defaultVal
func (d Data) MapStringInterfaceOr(key string, defaultVal map[string]interface{}) map[string]interface{} {
	val, err := d.MapStringInterface(key)

	if err != nil {
		return defaultVal
	}

	return val
}
