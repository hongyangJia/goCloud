package main

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
	"compress/gzip"
	"bytes"
)

const ORIGIN = "https://api.huobipro.com"

const ORIGIN_URL = "wss://api.huobipro.com/ws"

const TAG = "received"

var (
	ping = "ping: 18212558000"
	size = 512
)

func main() {
	initSend()
}

func initSend() {
	ws, err := websocket.Dial(ORIGIN_URL, "", ORIGIN)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(ping)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, size)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf(TAG+" zip : %s.\n", msg[:n])
	convert(msg)
}

func convert(byte []byte) {
	v, err := Unzip(byte)
	if err != nil {
		log.Fatal(err)
	}
	n, b, err := UnDate(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(TAG+"upzip: %s.\n", b[:n])
}

func UnDate(reader *gzip.Reader) (n int, b []byte, err error) {
	var msg = make([]byte, 512)
	number, err := reader.Read(msg)
	return number, msg, err;
}

func Unzip(data []byte) (*gzip.Reader, error) {
	b := bytes.NewReader(data)
	return gzip.NewReader(b)
}
