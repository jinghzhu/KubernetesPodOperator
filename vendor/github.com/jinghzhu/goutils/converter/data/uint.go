package data

import "strconv"

// UintPointer returns a pointer to of the uint value passed in.
func UintPointer(v uint) *uint {
	return &v
}

// UintVal returns the value of the uint pointer passed in or
// 0 if the pointer is nil.
func UintVal(v *uint) uint {
	if v != nil {
		return *v
	}

	return 0
}

// UintSlice converts a slice of uint values uinto a slice of
// uint pointers
func UintSlice(src []uint) []*uint {
	dst := make([]*uint, len(src))
	for i := 0; i < len(src); i++ {
		dst[i] = &(src[i])
	}

	return dst
}

// UintValSlice converts a slice of uint pointers uinto a slice of
// uint values
func UintValSlice(src []*uint) []uint {
	dst := make([]uint, len(src))
	for i := 0; i < len(src); i++ {
		if src[i] != nil {
			dst[i] = *(src[i])
		}
	}

	return dst
}

// StrToUint8 turn a string into a uint8
func StrToUint8(str string) (uint8, error) {
	i, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		return 0, err
	}

	return uint8(i), nil
}

// StrToUint16 turn a string into a uint16
func StrToUint16(str string) (uint16, error) {
	i, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}

	return uint16(i), nil
}

// StrToUint32 turn a string into a uint32
func StrToUint32(str string) (uint32, error) {
	i, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(i), nil
}

// StrToUint64 turn a string into a uint64.
func StrToUint64(str string) (uint64, error) {
	return strconv.ParseUint(str, 10, 64)
}
