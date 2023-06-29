package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.NewString()
	fmt.Println(id)
}
