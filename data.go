package data

import (
	"fmt"
	"strings"
)

// Data wrapper on map[string]interface{} to provide values for keys
type Data map[string]interface{}

func (d Data) navigate(key string) (interface{}, error) {
	parts := strings.SplitN(key, ".", 2)

	if len(parts) == 1 {
		return d[parts[0]], nil
	}

	if val, ok := d[parts[0]]; !ok {
		return nil, fmt.Errorf("missing key")
	} else if m, ok := val.(map[string]interface{}); !ok {
		return nil, fmt.Errorf("invalid key path")
	} else {
		return Data(m).navigate(parts[1])
	}
}

// Set is a method used to set a value at a given path, but will not
// create the path if missing
func (d Data) Set(key string, val interface{}) error {
	parts := strings.SplitN(key, ".", 2)

	if len(parts) == 1 {
		d[parts[0]] = val
		return nil
	}

	var m map[string]interface{}

	if x, ok := d[parts[0]]; !ok {
		return fmt.Errorf("missing key path")
	} else if m, ok = x.(map[string]interface{}); !ok {
		return fmt.Errorf("invalid key path")
	}

	return Data(m).Set(parts[1], val)
}

// SetCreate is a method used to set a value at a given path, and will
// create the path if missing, but will not overwrite
func (d Data) SetCreate(key string, val interface{}) error {
	parts := strings.SplitN(key, ".", 2)

	if len(parts) == 1 {
		d[parts[0]] = val
		return nil
	}

	var m map[string]interface{}

	if x, ok := d[parts[0]]; !ok {
		m = map[string]interface{}{}
		d[parts[0]] = m
	} else if m, ok = x.(map[string]interface{}); !ok {
		return fmt.Errorf("invalid key path")
	}

	return Data(m).SetCreate(parts[1], val)
}

// SetOverwrite is a method used to set a value at a given path, and will
// create the path if missing and will overwrite values if they are not correct
func (d Data) SetOverwrite(key string, val interface{}) {
	parts := strings.SplitN(key, ".", 2)

	if len(parts) == 1 {
		d[parts[0]] = val
		return
	}

	var m map[string]interface{}

	if x, ok := d[parts[0]]; !ok {
		m = map[string]interface{}{}
		d[parts[0]] = m
	} else if m, ok = x.(map[string]interface{}); !ok {
		m = map[string]interface{}{}
		d[parts[0]] = m
	}

	Data(m).SetOverwrite(parts[1], val)
}
