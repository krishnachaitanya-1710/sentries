package utilities

import (
	"../logger"
	"os"
)

// IfThen evaluates a condition, if true returns the parameters otherwise nil
func IfThen(condition bool, a interface{}) interface{} {
	if condition {
		return a
	}
	return nil
}

// IfThenElse evaluates a condition and executes logger, if true logs the first parameter otherwise the second
func IfThenElse(condition bool, a, b string) {
	if condition {
		logger.StreamGreen(a)
	} else {
		logger.StreamRed(b)
	}
}

// DefaultIfNil checks if the value is nil, if true returns the default value otherwise the original
func DefaultIfNil(value interface{}, defaultValue interface{}) interface{} {
	if value != nil {
		return value
	}
	return defaultValue
}

// Check whether the file is exist and creates of file doesn't exists
func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

// FirstNonNil returns the first non nil parameter
func FirstNonNil(values ...interface{}) interface{} {
	for _, value := range values {
		if value != nil {
			return value
		}
	}
	return nil
}

// Looping over elements in slices, arrays, maps, channels or strings is often better done with a range loop.

func ExecuteRule(condition bool, message, violationContext string) {
	if condition {
		logger.StreamGreen("pass	- " + message)
	} else {
		logger.StreamRed("fail	- " + message)
		//logger.StreamBlue(violationContext)
	}
}

func GetKeysofMap(userMap map[string]*string) []string {
	keys := make([]string, 0, len(userMap))
	for k := range userMap {
		keys = append(keys, k)
	}
	return keys
}

func getValues(userMap map[string]string) []string {
	values := make([]string, 0, len(userMap))
	for _, value := range userMap {
		values = append(values, value)
	}
	return values
}
