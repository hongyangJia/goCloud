package main

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
	"encoding/json"
	"strings"
	"goCloud/src/zips"
	"goCloud/src/websocket/api"
)

const TAG = "received"

var (
	ping = "req: market.lxtbtc.kline.1min"
	size = 512
)

func main() {
	initSend()
}

func initSend() {
	ws, err := websocket.Dial(api.ORIGIN_URL, "", api.ORIGIN)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(requestJson())); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, size)
	for {
		_, err = ws.Read(msg);
		if err != nil {
			log.Fatal(err)
		}
		convert(msg, ws)
	}

}

func convert(byt []byte, ws *websocket.Conn) {
	v, err := zips.Unzip(byt)
	if err != nil {
		log.Fatal(err)
	}
	_, b, err := zips. UnDate(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(TAG, string(b))
	fmt.Println("||")

	if strings.Contains(string(b), "ping") {
		s := strings.Replace(string(b), "ping", "pong", -1)
		fmt.Println(TAG, s)
		if _, err := ws.Write([]byte(s)); err != nil {
			log.Fatal(err)
		}
	}

}

func requestJson() string {
	a := make(map[string](string))
	a["sub"] = "market.lxtbtc.trade.detail"
	a["id"] = "id1"
	v1, _ := json.Marshal(a)
	println(string(v1))
	return string(v1);
}

