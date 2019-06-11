package data

import (
	"math"
	"strconv"
	"testing"

	"github.com/jinghzhu/goutils/converter"
)

func TestStrToUint8(t *testing.T) {
	validInts := []uint8{0, 1, math.MaxUint8}
	invalidInts := []string{"1.233", "a", "false", strconv.FormatUint(math.MaxUint64, 10)}

	for _, f := range validInts {
		_, err := StrToUint8(converter.Uint8ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but encounter error %s\n", f, err.Error())
		}
	}
	for _, f := range invalidInts {
		_, err := StrToUint8(f)
		if err == nil {
			t.Errorf("Should encounter error but pass for %+v\n", f)
		}
	}
}

func TestStrToUint16(t *testing.T) {
	validUints := []uint16{0, 1, math.MaxUint8, math.MaxUint16}
	invalidUints := []string{"1.233", "a", "false", strconv.FormatUint(math.MaxUint64, 10)}

	for _, f := range validUints {
		_, err := StrToUint16(converter.Uint16ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but encounter error %s\n", f, err.Error())
		}
	}
	for _, f := range invalidUints {
		_, err := StrToUint16(f)
		if err == nil {
			t.Errorf("Should encounter error but pass for %+v\n", f)
		}
	}
}

func TestStrToUint32(t *testing.T) {
	validUints := []uint32{0, 1, math.MaxUint8, math.MaxUint16, math.MaxUint32}
	invalidUints := []string{"1.233", "a", "false", strconv.FormatUint(math.MaxUint64, 10)}

	for _, f := range validUints {
		_, err := StrToUint32(converter.Uint32ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but encounter error %s\n", f, err.Error())
		}
	}
	for _, f := range invalidUints {
		_, err := StrToUint32(f)
		if err == nil {
			t.Errorf("Should encounter error but pass for %+v\n", f)
		}
	}
}

func TestStrToUint64(t *testing.T) {
	validUints := []uint64{0, 1, math.MaxUint8, math.MaxUint16, math.MaxUint32, math.MaxUint64}
	invalidUints := []string{"1.233", "a", "false"}

	for _, f := range validUints {
		_, err := StrToUint64(converter.Uint64ToStr(f))
		if err != nil {
			t.Errorf("Should pass for %+v but encounter error %s\n", f, err.Error())
		}
	}
	for _, f := range invalidUints {
		_, err := StrToUint64(f)
		if err == nil {
			t.Errorf("Should encounter error but pass for %+v\n", f)
		}
	}
}
