package main

import (
	"github.com/tidwall/gjson"
)
const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
func main() {
	value := gjson.Get(json, "ch")
    println(value)

}