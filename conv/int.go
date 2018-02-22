package conv

import (
	"fmt"
	"strconv"
)

// StringToInt64 - Convert string to type int64
func StringToInt64(v string) (int64, error) {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, fmt.Errorf(v + " is not int64 type")
	}

	return i, nil
}

// StringToInt - Convert string to type int
func StringToInt(v string) (int, error) {
	i, err := StringToInt64(v)
	if err != nil {
		return 0, err
	}

	return int(i), nil
}

// BoolToInt - Convert bool value to type int
func BoolToInt(v bool) (int64, error) {
	i, err := StringToInt64(fmt.Sprintf("%d", v))
	if err != nil {
		return 0, err
	}

	return i, nil
}

// CheckMapKeyToInt64 - Check whether the Map Key is int64 type and return value
func CheckMapKeyToInt64(m map[string]interface{}, key string) (int64, error) {
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
func CheckMapKeyToInt(m map[string]interface{}, key string) (int, error) {
	if _, ok := m[key]; !ok {
		return 0, fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].(float64); ok {
		return int(v), nil
	}

	return 0, fmt.Errorf(key + " is not int type")
}
