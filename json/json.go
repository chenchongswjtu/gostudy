package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name,omitempty"` // omitempty在name为空时反序列化之后的json中不会输出name{"age":27}，没有omitempty时是会输出的{"name":"","age":27}
	Age  int    `json:"age"`
}

func main() {
	p := person{
		Name: "",
		Age:  27,
	}

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json marshal err:", err)
	}

	fmt.Println(string(data))
}
