package main

import (
	"goCloud/src/websocket"
	"goCloud/src/websocket/api"
	"goCloud/src/websocket/sub"
	"goCloud/src/conversion"
	"log"
	"goCloud/src/websocket/topic"
	"goCloud/src/websocket/time"
	"goCloud/src/websocket/name"
)

func main() {
	 start()
}

var lis = make(map[string]string)

func start() {

	//aacTradeDetail := sub.TradeDetail(name.OMG)
	//aacMarketDetail := sub.MarketDetail(name.OMG)
	//aacMarketKline := sub.LastDayKLine(name.OMG)

	btcTradeDetail := sub.TradeDetail(name.BTC)
	btcMarketDetail := sub.MarketDetail(name.BTC)
	btcMarketKline := sub.LastDayKLine(name.BTC)

	//lis[aacTradeDetail] = name.OMG
	//lis[aacMarketDetail] = name.OMG
	//lis[aacMarketKline] = name.OMG

	lis[btcTradeDetail] = name.BTC
	lis[btcMarketDetail] = name.BTC
	lis[btcMarketKline] = name.BTC

	 con := websocket.Conifg{api.PRO_ORIGIN_URL, api.PRO_ORIGIN, lis, receive}
	 websocket.Start(&con)


}

func receive(v string) (c string, e error) {
	 //fmt.Println(v)
	if conversion.Ping(v) != nil {
		return
	}
	err := conversion.Status(v)
	if err != nil {
		log.Fatal(err)
	}
	ty := conversion.State(v)
	var state string

	ty1 := sub.LastDayKLin1e()
	if v, ok := lis[ty1]; ok {
		state = v
	}

	ty1 = sub.MarketDetail1(ty)
	if v, ok := lis[ty1]; ok {
		state = v
	}

	switch ty {
	case topic.KLine(state, time.MIN1):
		return
	case topic.MarketDepth(state, time.MIN1):
		return
	case topic.MarketDetail(state):
		conversion.MarketDetail(v, state)
	case topic.TradeDetail(state):
		conversion.TradeDetail(v, state)
	case topic.KLine(state, time.DAY1):
		conversion.MarketClose(v, state)
		//return sub.UnsubMarketLine(state, time.DAY1), errors.New("line")
		return "",nil
	}
	return "", nil

}
