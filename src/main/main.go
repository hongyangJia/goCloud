package main

import (
	"net/http"
	"goCloud/src/websocket/name"
	"goCloud/src/websocket"
	"goCloud/src/websocket/api"
	"log"
	"fmt"
	"goCloud/src/websocket/sub"
)

func main() {
	http.HandleFunc("/topics", topics)
	err := http.ListenAndServe("192.168.21.238:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func topics(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	start(w);
}

func start(w http.ResponseWriter) {
	topic := sub.MarketDetail(name.AAC)
	con := websocket.Conifg{api.HADAX_ORIGIN_URL, api.HADAX_ORIGIN, topic, w, receive}
	websocket.Start(&con)
}

func receive(s string, w http.ResponseWriter) {
	fmt.Print(s)
}
