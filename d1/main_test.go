package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

var tests = []struct {
	l      []int64
	k      int64
	output [][]int64
}{
	{
		[]int64{10, 15, 7, 11}, 17, [][]int64{{10, 7}},
	},
	{
		[]int64{10, 15, 7, 11}, 1, [][]int64{},
	},
	{
		[]int64{10, 15, 7, 10}, 20, [][]int64{{10, 10}},
	},
}

func TestHelloWorld(t *testing.T) {
	for _, v := range tests {
		results := pairOfSums(v.l, v.k)
		assert.Equal(t, v.output, results)
	}
}
