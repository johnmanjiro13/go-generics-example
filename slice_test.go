package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	incrementFunc := func(i int) int {
		return i + 1
	}
	wantInt := []int{2, 3, 4, 5, 6}
	gotInt := Map(ints, incrementFunc)
	assert.Equal(t, wantInt, gotInt)

	convertFunc := func(i int) string {
		return strconv.Itoa(i)
	}
	wantStr := []string{"1", "2", "3", "4", "5"}
	gotStr := Map(ints, convertFunc)
	assert.Equal(t, wantStr, gotStr)
}
