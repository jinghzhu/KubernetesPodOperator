package converter

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToString converts input to string.
func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)

	return string(res)
}

// ToJSON converts input to JSON string.
func ToJSON(obj interface{}) (string, error) {
	res, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// StToBool convert the input string to a boolean.
func StToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}
