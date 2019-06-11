package converter

import (
	"fmt"
	"testing"
)

func TestStToBool(t *testing.T) {
	testCases := []string{"true", "1", "True", "false", "0", "abcdef"}
	expected := []bool{true, true, true, false, false, false}
	for k, v := range testCases {
		res, _ := StToBool(v)
		if res != expected[k] {
			t.Log("Case ", k, ": expected ", expected[k], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestToJSON(t *testing.T) {
	testCases := []interface{}{"test", map[string]string{"a": "b", "b": "c"}, func() error { return fmt.Errorf("Error") }}
	expected := [][]string{
		{"\"test\"", "<nil>"},
		{"{\"a\":\"b\",\"b\":\"c\"}", "<nil>"},
		{"", "json: unsupported type: func() error"},
	}
	for i, test := range testCases {
		actual, err := ToJSON(test)
		if actual != expected[i][0] {
			t.Errorf("Expected toJSON(%v) to return '%v', got '%v'", test, expected[i][0], actual)
		}
		if fmt.Sprintf("%v", err) != expected[i][1] {
			t.Errorf("Expected error returned from toJSON(%v) to return '%v', got '%v'", test, expected[i][1], fmt.Sprintf("%v", err))
		}
	}
}
