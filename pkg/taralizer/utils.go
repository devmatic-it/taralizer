package taralizer

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

func GetMapIntValue(data map[string]interface{}, key string, location string) (int64, error) {
	result := int64(-1)
	val, exists := data[key]
	if exists {
		// OPA returns numbers as float64
		if floatVal, ok := val.(float64); ok {
			result = int64(floatVal)
		} else if intVal, ok := val.(int); ok {
			result = int64(intVal)
		} else if int64Val, ok := val.(int64); ok {
			result = int64Val
		} else if valNum, ok := val.(json.Number); ok {
			res, err := valNum.Int64()
			if err != nil {
				return 0, fmt.Errorf("parse error in %s: key '%s' is not a valid number", location, key)
			}
			result = res
		} else {
			return 0, fmt.Errorf("parse error in %s: key '%s' is not a number", location, key)
		}
	} else {
		return 0, fmt.Errorf("parse error in %s: cannot find key '%s' in given map", location, key)
	}

	return result, nil
}

func GetMapStringValue(data map[string]interface{}, key string, location string) (string, error) {
	val, exists := data[key]
	if exists {
		str, ok := val.(string)
		if !ok {
			return "", fmt.Errorf("parse error in %s: key '%s' is not a string", location, key)
		}
		return str, nil
	} else {
		return "", fmt.Errorf("parse error in %s: cannot find key '%s' in given map", location, key)
	}
}

type StringWriter struct {
	buf *string
}

func NewStringWriter(buf *string) StringWriter {
	w := StringWriter{}
	w.buf = buf
	return w
}

func (sw StringWriter) Write(p []byte) (n int, err error) {
	str := string(p)
	slog.Info("writing to string buffer", "content", str)
	*sw.buf = *sw.buf + str
	return len(str), nil
}

func (sw StringWriter) String() string {
	return *sw.buf
}
