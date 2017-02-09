package conv

import (
	"fmt"
)

// CheckMapKeyToMapInterface - Check whether  the Map Key is map[string]interface{} type and return value
func CheckMapKeyToMapInterface(m map[string]interface{}, key string) (map[string]interface{}, error) {
	if _, ok := m[key]; !ok {
		return nil, fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].(interface{}); ok {
		return v.(map[string]interface{}), nil
	}

	return nil, fmt.Errorf(key + " is not map[string]interface{} type")
}

// CheckMapKeyToInterfaceSlice - Check whether the Map Key is []interface{} type and return value
func CheckMapKeyToInterfaceSlice(m map[string]interface{}, key string) ([]interface{}, error) {
	if _, ok := m[key]; !ok {
		return nil, fmt.Errorf(key + " is nil")
	}

	if v, ok := m[key].([]interface{}); ok {
		return v, nil
	}
	return nil, fmt.Errorf(key + " is not []interface{} type")
}

//  CheckInterfaceToMapInterface - Check whether the interface is map[string]interface{} type and return value
func CheckInterfaceToMapInterface(i interface{}) (map[string]interface{}, error) {
	if v, ok := i.(interface{}); ok {
		return v.(map[string]interface{}), nil
	}
	return nil, fmt.Errorf("%v is not interface{} type", i)
}
