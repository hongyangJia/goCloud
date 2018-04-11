package main

import (
	"log"
	"fmt"
	"golang.org/x/net/websocket"
	"goCloud/src/zips"
	"goCloud/src/websocket/api"
	"goCloud/src/websocket/name"
	"goCloud/src/websocket/time"
	"goCloud/src/websocket/sub"
	"goCloud/src/common"
)

const (
	TAG = "received"
	SIZE  = 512
)
func main() {
	Start()
}

func Start() {

	ws, err := websocket.Dial(api.HADAX_ORIGIN_URL, api.PROTOCOL, api.HADAX_ORIGIN)
	defer ws.Close()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(requestJson())); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, SIZE)
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
	_, b, err := zips.UnDate(v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(TAG, string(b))
	pong,err:=common.ReplacePong(string(b))
	if err!=nil {
		return
	}
	if _, err := ws.Write([]byte(pong)); err != nil {
		log.Fatal(err)
	}
}

func requestJson() string {
	return sub.KLine(name.AAC,time.MIN1);
}

