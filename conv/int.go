package conv

import (
	"fmt"
	"strconv"
)

// ConvStringToInt64 - Convert string to type int64
func ConvStringToInt64(v string) (int64, error) {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(v + " is not int64 type")
	}

	return i, nil
}

// ConvStringToInt - Convert string to type int
func ConvStringToInt(v string) (int, error) {
	i, err := ConvStringToInt64(v)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// CheckMapKeyToInt64 - Check whether the Map Key is int64 type and return value
func ConvMapKeyToInt64(m map[string]interface{}, key string) (int64, error) {
	if _, ok := m[key]; !ok {
		return 0, fmt.Errorf(key + " is nil")
	}

	v, ok := m[key].(string)
	if !ok {
		return 0, fmt.Errorf("Cannot read int64 from map")
	}

	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(key + " is not int64 type")
	}

	return i, nil
}

// CheckMapKeyToInt - Check whether the Map Key it int64 type and return value
func ConvMapKeyToInt(m map[string]interface{}, key string) (int, error) {
	if _, ok := m[key]; !ok {
		return 0, fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].(float64); ok {
		return int(v), nil
	}

	return 0, fmt.Errorf(key + " is not int type")
}
