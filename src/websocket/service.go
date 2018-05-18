package websocket

import (
	"log"
	"golang.org/x/net/websocket"
	"goCloud/src/zips"
	"goCloud/src/websocket/api"
	"goCloud/src/common"
	"goCloud/src/websocket/sub"
)
const (
	TAG = "received"
	SIZE  = 512
)

func Start(conifg *Conifg) {
	ws, err := websocket.Dial(conifg.Url, api.PROTOCOL, conifg.Origin)
	defer ws.Close()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte(conifg.Topics.Parameter)); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, SIZE)
	for {
		_, err = ws.Read(msg);
		if err != nil {
			log.Fatal(err)
		}
		convert(msg, ws,conifg)
	}
}

func convert(byt []byte, ws *websocket.Conn,conifg *Conifg) {
	v, err := zips.Unzip(byt)
	if err != nil {
		log.Fatal(err)
	}
	_, b, err := zips.UnDate(v)
	if err != nil {
	}
	conifg.Call(string(b),conifg.Topics.State)
	pong,err:=common.ReplacePong(string(b))
	if err!=nil {
		return
	}
	if _, err := ws.Write([]byte(pong)); err != nil {
		log.Fatal(err)
	}
}

type Conifg struct {
	 Url string
	 Origin string
	 Topics  *sub.Detail
	 Call func(v string,state string)
}

