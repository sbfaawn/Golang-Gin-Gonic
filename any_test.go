package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hello() string {
	return "Hello"
}

func AnyTest(t *testing.T) {
	assert.Equal(t, "Hello", hello())
}
