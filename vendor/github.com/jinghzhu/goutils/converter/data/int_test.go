package data

import (
	"math"
	"strconv"
	"testing"

	"github.com/jinghzhu/goutils/converter"
)

func TestStrToInt8(t *testing.T) {
	validInts := []int8{0, 1, -1, math.MaxInt8, math.MinInt8}
	invalidInts := []string{"1.233", "a", "false", strconv.Itoa(int(math.MaxInt64))}

	for _, f := range validInts {
		_, err := StrToInt8(converter.Int8ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but got error %+v\n", f, err)
		}
	}
	for _, f := range invalidInts {
		_, err := StrToInt8(f)
		if err == nil {
			t.Errorf("Should get error but pass for %+v\n", f)
		}
	}
}

func TestStrToInt16(t *testing.T) {
	validInts := []int16{0, 1, -1, math.MaxInt8, math.MinInt8, math.MaxInt16, math.MinInt16}
	invalidInts := []string{"1.233", "a", "false", strconv.Itoa(int(math.MaxInt64))}

	for _, f := range validInts {
		_, err := StrToInt16(converter.Int16ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but got error %+v\n", f, err)
		}
	}
	for _, f := range invalidInts {
		_, err := StrToInt16(f)
		if err == nil {
			t.Errorf("Should get error but pass for %+v\n", f)
		}
	}
}

func TestStrToInt32(t *testing.T) {
	validInts := []int32{0, 1, -1, math.MaxInt8, math.MinInt8, math.MaxInt16, math.MinInt16, math.MinInt32, math.MaxInt32}
	invalidInts := []string{"1.233", "a", "false", strconv.Itoa(int(math.MaxInt64))}

	for _, f := range validInts {
		_, err := StrToInt32(converter.Int32ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but got error %+v\n", f, err)
		}
	}
	for _, f := range invalidInts {
		_, err := StrToInt32(f)
		if err == nil {
			t.Errorf("Should get error but pass for %+v\n", f)
		}
	}
}

func TestStrToInt64(t *testing.T) {
	validInts := []int64{0, 1, -1, math.MaxInt8, math.MinInt8, math.MaxInt16, math.MinInt16, math.MinInt32, math.MaxInt32, math.MaxInt64, math.MinInt64}
	invalidInts := []string{"1.233", "a", "false"}

	for _, f := range validInts {
		_, err := StrToInt64(converter.Int64ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but got error %+v\n", f, err)
		}
	}
	for _, f := range invalidInts {
		_, err := StrToInt64(f)
		if err == nil {
			t.Errorf("Should get error but pass for %+v\n", err)
		}
	}
}

func TestInterfaceToInt64(t *testing.T) {
	testCases := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
	expected := []int64{1000, -123, 0, 0}
	for k, v := range testCases {
		result, _ := InterfaceToInt64(v)
		if result != expected[k] {
			t.Log("Case ", k, ": expected ", expected[k], " when result is ", result)
			t.FailNow()
		}
	}
}
