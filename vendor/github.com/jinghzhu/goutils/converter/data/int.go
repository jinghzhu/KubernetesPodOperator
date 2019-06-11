package data

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/jinghzhu/goutils/valiadator"
)

// StrToInt8 turns a string into int8 boolean.
func StrToInt8(str string) (int8, error) {
	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, err
	}

	return int8(i), nil
}

// StrToInt16 turns a string into a int16.
func StrToInt16(str string) (int16, error) {
	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return 0, err
	}

	return int16(i), nil
}

// StrToInt32 turns a string into a int32.
func StrToInt32(str string) (int32, error) {
	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(i), nil
}

// StrToInt turns a string into a int.
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StrToInt64 turns a string into a int64.
func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// InterfaceToInt64 converts input to an integer type 64, or 0 if input is not an integer.
func InterfaceToInt64(value interface{}) (res int64, err error) {
	val := reflect.ValueOf(value)

	switch value.(type) {
	case int, int8, int16, int32, int64:
		res = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		res = int64(val.Uint())
	case string:
		if valiadator.IsInt(val.String()) {
			res, err = strconv.ParseInt(val.String(), 0, 64)
			if err != nil {
				res = 0
			}
		} else {
			err = fmt.Errorf("math: square root of negative number %g", value)
			res = 0
		}
	default:
		err = fmt.Errorf("math: square root of negative number %g", value)
		res = 0
	}

	return
}

// Int64Pointer returns a pointer to of int64 value passed in.
func Int64Pointer(v int64) *int64 {
	return &v
}

// Int64Val returns the value of int64 pointer passed in or 0 if the pointer is nil.
func Int64Val(v *int64) int64 {
	if v != nil {
		return *v
	}

	return 0
}

// Int64Slice converts a slice of int64 values into a slice of int64 pointers
func Int64Slice(src []int64) []*int64 {
	dst := make([]*int64, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}

	return dst
}

// Int64ValSlice converts a slice of int64 pointers into a slice of int64 values
func Int64ValSlice(src []*int64) []int64 {
	dst := make([]int64, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}

	return dst
}

// Int64Map converts a string map of int64 values into a string map of int64 pointers
func Int64Map(src map[string]int64) map[string]*int64 {
	dst := make(map[string]*int64)
	for k, val := range src {
		v := val
		dst[k] = &v
	}

	return dst
}

// Int64ValMap converts a string map of int64 pointers into a string
// map of int64 values
func Int64ValMap(src map[string]*int64) map[string]int64 {
	dst := make(map[string]int64)
	for k, val := range src {
		if val != nil {
			dst[k] = *val
		}
	}

	return dst
}
