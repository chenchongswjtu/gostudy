package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	fig := figure.NewFigure("Hello world", "", true)
	fmt.Println(fig.String())

	// 获得map的key值
	//if _, found := colors[color]; !found {
	//	log.Fatalf("invalid color. must be one of: %s", reflect.ValueOf(colors).MapKeys())
	//}
	fig = figure.NewColorFigure("Hello world", "", "red", true)
	fmt.Println(fig.ColorString())

	fig = figure.NewFigure("Hello world", "slant", true)
	fmt.Println(fig.String())

	myFigure := figure.NewFigure("Hello world", "alphabet", true)
	fmt.Println(myFigure.String())
}
