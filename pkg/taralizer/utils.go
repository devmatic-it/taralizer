package taralizer

import (
	"encoding/json"
	"log"
)

func GetMapIntValue(data map[string]interface{}, key string) int64 {
	result := int64(-1)
	val, exists := data[key]
	if exists {
		valNum, exists := val.(json.Number)
		if exists {
			result, _ = valNum.Int64()
		}
	}

	return result
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
