package main

import (
"golang.org/x/net/websocket"
"log"
"fmt"
"bytes"
"compress/gzip"
)

func main()  {
	ExampleDial()
}

func ExampleDial() {
	origin := "https://api.huobipro.com"
	url := "wss://api.huobipro.com/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Received: %s.\n", msg[:n])

	b := bytes.NewReader(msg)

	v1,_:=gzip.NewReader(b)
	var v int
	var msg1 = make([]byte, 512)
	if v, err = v1.Read(msg1); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg1[:v])
}