package conv

import (
	"fmt"
)

// IntToBool - Convert Int to type bool
func IntToBool(v int) bool {
	if v > 0 {
		return true
	}

	return false
}

// Int64ToBool - Convert Int64 to type bool
func Int64ToBool(v int64) bool {
	return IntToBool(int(v))
}

// CheckMapKeyToBool - Check Map key whether is bool type and return value
func CheckMapKeyToBool(m map[string]interface{}, key string) (bool, error) {
	if _, ok := m[key]; !ok {
		return false, fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].(bool); ok {
		return bool(v), nil
	}
	return false, fmt.Errorf(key + " is not bool type")
}
