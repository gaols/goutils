package goutils

import (
	"fmt"
	"strconv"
)

// Str2Int convert string to int.
func Str2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

// ToInt convert string and any other int types to int.
func ToInt(v interface{}) (int, error) {
	switch v.(type) {
	case string:
		return Str2Int(v.(string))
	case int:
		return v.(int), nil
	case int8:
		return int(v.(int8)), nil
	case int32:
		return int(v.(int32)), nil
	case int64:
		return int(v.(int64)), nil
	case uint8:
		return int(v.(uint8)), nil
	case uint32:
		return int(v.(uint32)), nil
	case uint64:
		return int(v.(uint64)), nil
	}

	return -1, fmt.Errorf("cannot parse value: %v", v)
}

// ToIntWithDefault tries to convert string and any other int types to int, 
// return a default value provided if error occurs.
func ToIntWithDefault(v interface{}, defVal int) int {
	ret, err := ToInt(v)
	if err == nil {
		return ret
	}

	return defVal
}
