package main_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AnyTest(t *testing.T) {
	bookId := "-12"

	bookIdNum, err := strconv.ParseUint(bookId, 10, 32)

	fmt.Println(bookIdNum)
	assert.Error(t, err, "")
}
