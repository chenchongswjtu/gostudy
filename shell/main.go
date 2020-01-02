package main

import (
	"fmt"

	sh "github.com/codeskyblue/go-sh"
)

func main() {
	if out, err := sh.Command("ls").Output(); err == nil {
		fmt.Println(string(out))
	}

	if out, err := sh.Command("docker", "images").Command("grep", "ccenv").Output(); err == nil {
		fmt.Println(string(out))
	}
}
