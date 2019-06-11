package string

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"unicode/utf8"
)

// ReplaceWithRunes returns a new string. It replaces the runes in the range of given string.
func ReplaceWithRunes(str string, replacement rune, start, end int) string {
	if start < 0 || end >= len(str) || start > end {
		return str
	}

	return str[:start] + string(replacement) + str[end:]
}

// ReplaceWithStr returns a new string. It replaces the string in the range of given string.
func ReplaceWithStr(str, replacement string, start, end int) string {
	if start < 0 || end >= len(str) || start > end {
		return str
	}

	return str[:start] + replacement + str[end:]
}

// ReplaceWithRune returns a new string by replacing one of its element with given rune.
func ReplaceWithRune(str string, index int, replacement rune) string {
	if index < 0 || index >= len(str) {
		return str
	}
	result := []rune(str)
	result[index] = replacement

	return string(result)
}

// ReverseStrings accepts a string array and reverses each of its element.
func ReverseStrings(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}

// Reverse reverses the given string.
func Reverse(s string) string {
	return ReverseRange(s, 0, len(s)-1)
}

// ReverseRange reverses the given part of string.
func ReverseRange(s string, start, end int) string {
	if len(s) <= 1 || end <= start || end > len(s) || start < 0 || end < 0 {
		return s
	}

	sArr := []rune(s)
	// if replace i, j = i + 1, j - 1 with i++, j--
	// it will throw error msg: syntax error: missing { after for clause
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		temp := sArr[i]
		sArr[i] = sArr[j]
		sArr[j] = temp
	}

	return string(sArr)
}

// NonEmpty returns a slice holding only the non-empty string.
func NonEmpty(strings []string) []string {
	count := 0
	for _, v := range strings {
		if v != "" {
			strings[count] = v
			count++
		}
	}

	return strings[:count]
}

// NonEmptyBySlice is the same as NonEmpty(). Implemented with slice and not modify original data structure.
func NonEmptyBySlice(strings []string) []string {
	newStrings := strings[:0] // 0 length slice
	for _, v := range strings {
		if v != "" {
			newStrings = append(newStrings, v)
		}
	}

	return newStrings
}

// EqualSlice checks whether two slices of string type is equal.
func EqualSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

// IsAnagram checks whether two strings contain same characters.
// Input: s1 = "test", s2 = "etst"
// Output: true
func IsAnagram(s1, s2 string) bool {
	if len(s1) == 0 && len(s2) == 0 {
		return true
	} else if len(s1) == 0 || len(s2) == 0 || len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)
	for _, v := range s1 {
		count, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = count + 1
		}
	}
	for _, v := range s2 {
		count, ok := m[v]
		if !ok || count == 0 {
			return false
		}
		m[v] = count - 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func GenerateToken() string {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}

	return base64.URLEncoding.EncodeToString(rb)
}

// Compare return 0 if they are the same string. Otherwise, return -1 if a is "smaller" than b.
func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func IndexByte(s string, b byte) int {
	for index := 0; index < len(s); index++ {
		if s[index] == b {
			return index
		}
	}

	return -1
}

/*
   in package "unicode/utf8"
   const (
           RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
           RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
           MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
           UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
   )
*/
func IndexRune(s string, r rune) int {
	if r < utf8.RuneSelf {
		return IndexByte(s, byte(r))
	} else {
		for k, v := range s {
			if v == r {
				return k
			}
		}
	}

	return -1
}

// EqualFoldSimple only deals with English case.
// if input is like Chinese, please call EqualFold.
func EqualFoldSimple(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] < s2[i] {
			if s2[i] != s1[i]+'a'-'A' {
				return false
			}
		} else if s1[i] > s2[i] {
			if s1[i] != s2[i]+'a'-'A' {
				return false
			}
		}
	}

	return true
}

func Fields(s string, sep rune) []string {
	if len(s) == 0 {
		a := make([]string, 1)
		return a
	}

	inField := false
	var n int = 0

	for _, v := range s {
		if v == sep && !inField {
			inField = true
			n++
		} else if v != sep {
			inField = false
		}
	}

	if s[len(s)-1] != byte(sep) {
		n++
	}

	strArr := make([]string, n)
	index := 0
	fieldStart := -1

	for k, v := range s {
		if v == sep {
			if fieldStart != -1 { // important while the spe occurs in the first place
				strArr[index] = s[fieldStart:k]
				index++
				fieldStart = -1
			}
		} else if fieldStart == -1 {
			fieldStart = k
		}
	}

	if fieldStart != -1 { // important because last field might end at EOF.
		strArr[index] = s[fieldStart:]
	}

	return strArr
}

// HasPrefix checks whether pre is the prefix of s.
func HasPrefix(s, pre string) bool {
	return len(s) >= len(pre) && s[0:len(pre)] == pre
}

// HasSuffix checks whether suf is the suffix of s.
func HasSuffix(s, suf string) bool {
	return len(s) >= len(suf) && s[len(s)-len(suf):] == suf
}

// Contain returns true if substr is the sub string of s.
func Contain(s, substr string) bool {
	for i := 0; i < len(s)-len(substr)+1; i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}

	return false
}

// IntToString converts an int to string.
// Input: []int{1, 2, 3}
// Output: [1, 2, 3]
func IntToString(values []int) string {
	var buffer bytes.Buffer
	buffer.WriteByte('[')
	for k, v := range values {
		if k > 0 {
			buffer.WriteByte(',')
		}
		fmt.Fprintf(&buffer, "%d", v)
	}
	buffer.WriteByte(']')

	return buffer.String()
}

// To-Do: need to run performance test between JoinByBuffer and Join
func JoinByBuffer(strArr []string, sep string) string {
	if len(strArr) == 0 {
		return ""
	} else if len(strArr) == 1 {
		return strArr[0]
	}

	var buffer bytes.Buffer
	for k, v := range strArr {
		buffer.WriteString(v)
		if k != len(strArr)-1 {
			buffer.WriteString(sep)
		}
	}

	return buffer.String()
}

func Join(strArr []string, sep string) string {
	if len(strArr) == 0 {
		return ""
	} else if len(strArr) == 1 {
		return strArr[0]
	}

	n := len(sep) * (len(strArr) - 1)
	for i := 0; i < len(strArr); i++ {
		n += len(strArr[i])
	}
	b := make([]byte, n)
	bp := copy(b, strArr[0])
	for _, s := range strArr[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}

	return string(b)
}

func LastIndexByte(s string, b byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == b {
			return i
		}
	}

	return -1
}

// Title will change the first character of each string to upper case.
// Input: "test1 !test2 test3"
// Output: Test1 !test2 Test3
// strings.Title() output: Test1 !Test2 Test3
func Title(s string) string {
	var c rune
	sArr := []rune(s)
	var sep bool
	for k, v := range sArr {
		if v == c {
			sep = true
			continue
		}
		if sep {
			if 'a' <= v && v <= 'z' {
				sArr[k] = v - ('a' - 'A')
			}
			sep = false
			continue
		}
		if k == 0 && 'a' <= v && v <= 'z' {
			sArr[0] = v - ('a' - 'A')
		}
	}

	return string(sArr)
}

// TrimAllSpace removes all spaces in both fron and end field of given string.
func TrimAllSpace(s string) string {
	n := 0
	index := 0
	var space rune
	for _, v := range s {
		if v == space {
			n++
		}
	}
	rArr := make([]rune, len(s)-n)
	for _, v := range s {
		if v != space {
			rArr[index] = v
			index++
		}
	}

	return string(rArr)
}
