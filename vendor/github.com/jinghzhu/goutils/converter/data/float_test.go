package data

import (
	"math"
	"strconv"
	"testing"

	"github.com/jinghzhu/goutils/converter"
)

func TestStrToFloat64(t *testing.T) {
	testCases := []string{"", "1", "-.01342", "10.", "empty", "1.23e3", ".23e10"}
	expected := []float64{0, 1, -0.01342, 10.0, 0, 1230, 0.23e10}
	for k, v := range testCases {
		res, _ := StrToFloat64(v)
		if res != expected[k] {
			t.Log("Case ", k, ": expected ", expected[k], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestStrToFloat32(t *testing.T) {
	validFloats := []float32{1.0, -1, math.MaxFloat32, math.SmallestNonzeroFloat32, 0, 5.494430303}
	invalidFloats := []string{"a", strconv.FormatFloat(math.MaxFloat64, 'f', -1, 64), "true"}

	for _, f := range validFloats {
		_, err := StrToFloat32(converter.Float32ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but got error %+v\n", f, err)
		}
	}
	for _, f := range invalidFloats {
		_, err := StrToFloat32(f)
		if err == nil {
			t.Errorf("Should get error but pass for %+v\n", f)
		}
	}
}
