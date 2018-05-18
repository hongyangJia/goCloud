package main

import (
	"goCloud/src/websocket/name"
	"goCloud/src/websocket"
	"goCloud/src/websocket/api"
	"fmt"
	"goCloud/src/websocket/sub"
	"goCloud/src/conversion"
	"log"
	"time"
)

func main() {

	  go start()
	time.Sleep(1 * time.Second)
	 start1()
}

func start() {
	topic := sub.TradeDetail(name.AE)
	con := websocket.Conifg{api.HADAX_ORIGIN_URL, api.HADAX_ORIGIN, topic, receive}
	websocket.Start(&con)
}
func start1() {
	topic := sub.MarketDetail(name.AE)
	con := websocket.Conifg{api.HADAX_ORIGIN_URL, api.HADAX_ORIGIN, topic, receive}
	websocket.Start(&con)
}

func receive(v string, state string) {
	//fmt.Println(v)
	if conversion.Ping(v) != nil {
		return
	}
	err := conversion.Status(v);
	if err != nil {
		log.Fatal(err)
	}
	switch state {
	case sub.KLINE:
		fmt.Println("KLINE")
	case sub.MARKET_DEPTH:
		fmt.Println("MARKET_DEPTH")
	case sub.TRADE_DETAIL:
		conversion.TradeDetail(v)
	case sub.MARKET_DETAIL:
		conversion.Launch(v)
	default:
		fmt.Println("default")
	}

}
