package data

// Interface provides interface{} if available or else returns an error
func (d Data) Interface(key string) (interface{}, error) {
	return d.navigate(key)
}

// MustInterface provides interface{} from key or panics
func (d Data) MustInterface(key string) interface{} {
	val, err := d.Interface(key)

	if err != nil {
		panic(err)
	}

	return val
}

// InterfaceOr provides interface{} from key or defaultVal
func (d Data) InterfaceOr(key string, defaultVal interface{}) interface{} {
	val, err := d.Interface(key)

	if err != nil {
		return defaultVal
	}

	return val
}
