package taralizer

import (
	"log"
)

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
