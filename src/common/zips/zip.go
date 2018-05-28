package zips

import (
	"compress/gzip"
	"bytes"
)

func UnDate(reader *gzip.Reader) (n int, b []byte, err error) {
	var msg = make([]byte, 512)
	number, err := reader.Read(msg)
	return number, msg, err;
}

func Unzip(data []byte) (*gzip.Reader, error) {
	b := bytes.NewReader(data)
	return gzip.NewReader(b)
}