package data

import "strconv"

// StrToFloat64 converts string to float, or 0.0 if the input is not a float.
func StrToFloat64(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}

	return res, err
}

// StrToFloat32 turn a string into a float32.
func StrToFloat32(str string) (float32, error) {
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}
