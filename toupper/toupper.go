package main

import (
	"fmt"
	"strings"
)

func toUpper(str string) string {
	str = strings.ReplaceAll(str, ".", "_")
	return strings.ToUpper(str)
}

func main() {
	str := "core.peer.deliveryService.delayPeersRatio"
	fmt.Println(toUpper(str))
}
