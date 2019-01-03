package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input  []int64
	output []int64
}{
	{
		[]int64{1, 2, 3, 4, 5},
		[]int64{120, 60, 40, 30, 24},
	},
	{
		[]int64{3, 2, 1},
		[]int64{2, 3, 6},
	},
}

func TestProduct(t *testing.T) {
	for _, v := range tests {
		r := product(v.input)
		assert.Equal(t, v.output, r)
	}
}
