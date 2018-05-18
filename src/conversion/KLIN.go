package conversion

import (
	"github.com/tidwall/gjson"
	"fmt"
	"errors"
	"goCloud/src/websocket/api"
)

func Launch(v string){
	ch := gjson.Get(v, "ch")
	fmt.Println(ch.String())
	open := gjson.Get(v, "tick.open")
	fmt.Println(open.String())
}

func Status(v string) error{
	ch := gjson.Get(v, "status")
	if ch.String()!="error" {
		return nil
	}
	return errors.New("Status is an error")

}

func Ping(v string) error{
	ch := gjson.Get(v, "ping")
	if ch.String()==api.PROTOCOL {
		return nil
	}
	return errors.New("ping is not null")
}

func TradeDetail(v string)  {

	data := gjson.Get(v, "tick.data")
	re:= data.Array()
	for _, v := range re {
		price := v.Get("price").String()
		direction:=v.Get("direction").String()
		fmt.Println(price)
		fmt.Println(direction)
	}
}