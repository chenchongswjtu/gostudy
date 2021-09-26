package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {
	value := gjson.Get(json, "name.last") // 直接获取name中的last属性值
	fmt.Println(value.String())
}
