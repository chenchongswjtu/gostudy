package main

import (
	"fmt"

	sh "github.com/codeskyblue/go-sh"
)

func main() {
	if out, err := sh.Command("ls").Output(); err == nil {
		fmt.Println(string(out))
	}

	if out, err := sh.Command("ls").Command("grep", "ccenv").Output(); err == nil {
		fmt.Println(string(out))
	}
}
