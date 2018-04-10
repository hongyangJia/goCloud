package main

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
	"compress/gzip"
	"bytes"
	"encoding/json"
)

const ORIGIN = "https://api.hadax.com"

const ORIGIN_URL = "wss://api.hadax.com/ws"

const TAG = "received"

var (
	ping = "req: market.lxtbtc.kline.1min"
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
	if _, err := ws.Write([]byte(requestJson())); err != nil {
		log.Fatal(err)
	}

	for {
		var msg = make([]byte, size)
		_,err= ws.Read(msg);
		convert(msg)
	}
}

func convert(byte []byte) {
	v, err := Unzip(byte)
	if err != nil {
		log.Fatal(err)
	}
	_, b, err := UnDate(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(TAG, string(b))
	fmt.Println("||")
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

func requestJson() string {
	a := make(map[string](string))
	a["sub"] = "market.cdcbtc.trade.detail"
	a["id"] = "id1"
	v1, _ := json.Marshal(a)
	println(string(v1))
	return string(v1);
}


type parpam struct {
	req string
	id  string
}
