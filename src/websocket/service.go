package websocket

import (
	"log"
	"golang.org/x/net/websocket"
	"goCloud/src/common/zips"
	"goCloud/src/websocket/api"
	"goCloud/src/common"
)

const (
	TAG  = "received"
	SIZE = 512
)

func Start(conifg *Conifg) {
	ws, err := websocket.Dial(conifg.Url, api.PROTOCOL, conifg.Origin)
	defer ws.Close()
	if err != nil {
		log.Fatal(err)
	}
	for i, _ := range conifg.Topics {
		if _, err := ws.Write([]byte(i)); err != nil {
			log.Fatal(err)
		}
	}
	var msg = make([]byte, SIZE)
	for {
		_, err = ws.Read(msg);
		if err != nil {
			log.Fatal(err)
		}
		convert(msg, ws, conifg)
	}
}

func convert(byt []byte, ws *websocket.Conn, conifg *Conifg) {
	v, err := zips.Unzip(byt)
	if err != nil {
		log.Fatal(err)
	}
	_, b, err := zips.UnDate(v)

	c, err := conifg.Call(string(b))
	if err != nil {
		if _, err := ws.Write([]byte(c)); err != nil {
			log.Fatal(err)
		}
	}
	pong, err := common.ReplacePong(string(b))
	if err != nil {
		return
	}
	if _, err := ws.Write([]byte(pong)); err != nil {
		log.Fatal(err)
	}
}

type Conifg struct {
	Url    string
	Origin string
	Topics map[string]string
	Call   func(v string) (c string, err error)
}
