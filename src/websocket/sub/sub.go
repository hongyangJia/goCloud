package sub

import (
	"goCloud/src/websocket/topic"
	"encoding/json"
)
var topics= make(map[string](string))

func KLine(symbol string, period string) string {
	topics["sub"] = topic.KLine(symbol, period)
	topics["id"] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func MarketDepth(symbol string, types string) string {
	topics["sub"] = topic.MarketDepth(symbol, types)
	topics["id"] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func TradeDetail(symbol string) string {
	topics["sub"] = topic.TradeDetail(symbol)
	topics["id"] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func MarketDetail(symbol string) string {
	topics["sub"] = topic.MarketDetail(symbol)
	topics["id"] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}