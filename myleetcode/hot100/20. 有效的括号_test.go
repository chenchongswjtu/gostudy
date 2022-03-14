package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	inputs := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
	}

	excepted := []bool{
		true,
		true,
		false,
		false,
		true,
	}
	for i, input := range inputs {
		assert.Equal(t, excepted[i], isValid(input))
	}
}
