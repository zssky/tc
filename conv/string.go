package conv

import (
	"fmt"
)

// ConvToString - Convert value to string type
func ConvToString(i interface{}) string {
	return fmt.Sprintf("%v", i)
}

// CheckMapKeyToString - Check whether the Map key is string type and return value
func CheckMapKeyToString(m map[string]interface{}, key string) (string, error) {
	if _, ok := m[key]; !ok {
		return "", fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].(string); ok {
		return v, nil
	}

	return "", fmt.Errorf(key + " is not string type")
}
