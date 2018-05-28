package sub

import (
	"goCloud/src/websocket/topic"
	"encoding/json"
	"time"
)

const KLINE = "KLine"
const SUB = "sub"
const SUSUB = "unsub"
const ID = "id"

var topics = make(map[string](interface{}))

func KLine(symbol string, period string) string {
	topics[SUB] = topic.KLine(symbol, period)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func MarketDepth(symbol string, types string) string {
	topics[SUB] = topic.MarketDepth(symbol, types)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func TradeDetail(symbol string) string {
	topics[SUB] = topic.TradeDetail(symbol)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func MarketDetail(symbol string) string {
	topics[SUB] = topic.MarketDetail(symbol)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func MarketDetail1(symbol string) string {
	topics[SUB] = symbol
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

func UnsubMarketLine(symbol string, types string) string {
	topics := make(map[string](interface{}))
	topics[SUSUB] = topic.KLine(symbol, types)
	topics[ID] = "id1"
	v, _ := json.Marshal(topics)
	return string(v)
}

var klinMap = make(map[string](interface{}))

const FROM = "from"
const TO = "to"

func LastDayKLine(symbol string) string {
	klinMap[SUB] = topic.KLine(symbol, "1day")
	klinMap[ID] = "id1"
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	klinMap[FROM] = yesTime.Unix()
	klinMap[TO] = nTime.Unix()
	v, _ := json.Marshal(klinMap)
	return string(v)
}

func LastDayKLin1e() string {
	v, _ := json.Marshal(klinMap)
	return string(v);
}
