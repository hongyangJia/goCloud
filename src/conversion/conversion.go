package conversion

import (
	"github.com/tidwall/gjson"
	"fmt"
	"errors"
	"goCloud/src/websocket/api"
	"goCloud/src/common"
)

func KLine(v string) {
	ch := gjson.Get(v, "ch")
	fmt.Println(ch.String())
	open := gjson.Get(v, "tick.open")
	fmt.Println(open.String())
}

func Status(v string) error {
	ch := gjson.Get(v, "status")
	if ch.String() != "error" {
		return nil
	}
	return errors.New("Status is an error")

}

func Ping(v string) error {
	ch := gjson.Get(v, "ping")
	if ch.String() == api.PROTOCOL {
		return nil
	}
	return errors.New("ping is not null")
}

func TradeDetail(v string, k string) {
	if open == 0 {
		return
	}
	data := gjson.Get(v, "tick.data")
	re := data.Array()
	for _, v := range re {
		price := v.Get("price").String()
		direction := v.Get("direction").String()
		amount := v.Get("amount").String()
		fmt.Print(k + "  : ")
		fmt.Print(direction)
		fmt.Print("  ")
		fmt.Print(price)
		fmt.Print("  ")
		fmt.Print(amount)
		fmt.Println(" ")
	}
}

func MarketDetail(v string, k string) {
	if open != 0 {
		high := gjson.Get(v, "tick.high");
		low := gjson.Get(v, "tick.low");
		close := gjson.Get(v, "tick.close").Float()
		fmt.Print(k+" 实时价钱 : ", close)
		fmt.Print("  ")
		fmt.Print("涨幅 ： ")
		fmt.Print(common.Percent(close, open), "%")
		fmt.Print("  ")
		fmt.Print("最高价 : ")
		fmt.Print(high)
		fmt.Print("  ")
		fmt.Print("最低价")
		fmt.Println(low)
	}
}

var open float64;

func MarketClose(v string, k string) {
	open = gjson.Get(v, "tick.open").Float()
	fmt.Println(k+" 开盘价 : ", open)
}

func State(v string) string {
	subbed := gjson.Get(v, "ch")
	return subbed.String()
}
