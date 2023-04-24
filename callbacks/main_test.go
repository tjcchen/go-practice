package main

import (
	"test"
	"github.com/stretchr/testify/assert"
)

func add(a, b int) int {
	return a + b
}

func TestIncrease(t *testing.T) {
	t.Log("start testing")
	result := add(1, 2)
	assert.Equal(t, result, 3)
}