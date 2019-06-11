package utils

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Struct2String accepts any interface{} and return to JSON based string.
func Struct2String(v interface{}) string {
	result, err := json.Marshal(v)
	if err != nil {
		errMsg := "Fail to translate to json"
		fmt.Println(errMsg)
		return fmt.Sprintf("%v", v)
	}
	return string(result)
}

// PanicHandler catches a panic and logs an error. Suppose to be called via defer.
func PanicHandler() (caller string, fileName string, lineNum int, stackTrace string, rec interface{}) {
	buf := make([]byte, stackBuffer)
	runtime.Stack(buf, false)
	name, file, line := GetCallerInfo(2)
	if r := recover(); r != nil {
		caller, fileName, stackTrace = name, file, string(buf)
		lineNum = line
		rec = r
		fmt.Printf("%s %s ln%d: PANIC Defered : %v\n", name, file, line, r)
		fmt.Printf("%s %s ln%d: Stack Trace : %s", name, file, line, string(buf))
	}

	return caller, fileName, lineNum, stackTrace, rec
}

// GetCallerInfo returns the name of method caller and file name. It also returns the line number.
func GetCallerInfo(level int) (caller, fileName string, lineNum int) {
	if level < 1 || level > maxCallerLevel {
		level = defaultCallerLevel
	}

	pc, file, line, ok := runtime.Caller(level)
	fileDefault := ""
	lineDefault := -1
	nameDefault := ""
	if ok {
		fileDefault = file
		lineDefault = line
	}
	details := runtime.FuncForPC(pc)
	if details != nil {
		nameDefault = details.Name()
	}

	return nameDefault, fileDefault, lineDefault
}

// GetMountPoints returns all moutpoins in a string array
func GetMountPoints(server string) ([]string, error) {
	b, err := exec.Command("showmount", "-e", server).Output()
	if err != nil {
		fmt.Println("error in showmount: " + err.Error())
		return nil, err
	}
	s := strings.TrimSpace(string(b))
	// The fist line of showmount -e <server> is Exports list on <server>
	firstLine := strings.Index(s, "\n")
	sArr := strings.Split(s[firstLine+1:], "\n")
	for i := 0; i < len(sArr); i++ {
		index := strings.Index(sArr[i], " ")
		temp := sArr[i]
		sArr[i] = temp[:index]
	}
	return sArr, nil
}

// Locate returns the line number and file name in the current goroutine statck trace. The argument skip is the number of stack frames to ascend, with 0 identifying the caller of Caller.
func Locate(skip int) (filename string, line int) {
	if skip < 0 {
		skip = 2
	}
	_, path, line, ok := runtime.Caller(skip)
	file := ""
	if ok {
		_, file = filepath.Split(path)
	} else {
		fmt.Println("Fail to get method caller")
		line = -1
	}
	return file, line
}

// Swap is to swap the value for given position in an array.
func Swap(data []int, i, j int) {
	if len(data) < 2 || i < 0 || i > j || j >= len(data) {
		return
	}
	data[i], data[j] = data[j], data[i]
}

// Retry will retry the given condition function with specific time interval and retry round.
// It will return true if the condition is met. If it is timeout, it will return false. Otherwise,
// it will return the error encountered in the retry round.
func Retry(interval time.Duration, round int, retry func() (bool, error)) (bool, error) {
	if round < 1 {
		round = 1
	}
	if interval > time.Hour {
		interval = time.Hour
	}
	var err error
	done := false
	for i := 0; i < round; i++ {
		done, err = retry()
		if done {
			break
		}
		time.Sleep(interval)
	}

	if done {
		return true, nil
	}
	if err != nil {
		return false, err
	}

	return false, nil
}
