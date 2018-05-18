package sub

import (
	"goCloud/src/websocket/topic"
	"encoding/json"
)

const KLINE = "KLine"
const MARKET_DEPTH = "MarketDepth"
const TRADE_DETAIL = "TradeDetail"
const MARKET_DETAIL = "MarketDetail"
const SUB = "sub"
const ID = "id"

var topics = make(map[string](string))

func KLine(symbol string, period string) *Detail{
	topics[SUB] = topic.KLine(symbol, period)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return &Detail{string(v), KLINE}
}

func MarketDepth(symbol string, types string)*Detail {
	topics[SUB] = topic.MarketDepth(symbol, types)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return &Detail{string(v), MARKET_DEPTH}
}

func TradeDetail(symbol string) *Detail {
	topics[SUB] = topic.TradeDetail(symbol)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return &Detail{string(v), TRADE_DETAIL}
}

func MarketDetail(symbol string) *Detail {
	topics[SUB] = topic.MarketDetail(symbol)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return &Detail{string(v), MARKET_DETAIL}
}

type Detail struct {
	Parameter string
	State     string
}
