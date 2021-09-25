package taralizer

import (
	"encoding/json"
	"log"
)

func GetMapIntValue(data map[string]interface{}, key string, location string) int64 {
	result := int64(-1)
	val, exists := data[key]
	if exists {
		valNum, exists := val.(json.Number)
		if exists {
			result, _ = valNum.Int64()
		} else {
			log.Fatalf("Parse Error in %s: Given key '%s' is not a number", location, key)
		}
	} else {
		log.Fatalf("Parse Error in %s: Cannot find key '%s' in given map", location, key)
	}

	return result
}

func GetMapStringValue(data map[string]interface{}, key string, location string) string {
	val, exists := data[key]
	if exists {
		return val.(string)
	} else {
		log.Fatalf("Parse Error in %s: Cannot find key '%s' in given map", location, key)
	}

	return ""
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
	log.Println(str)
	*sw.buf = *sw.buf + str
	return len(str), nil
}

func (sw StringWriter) String() string {
	return *sw.buf
}
