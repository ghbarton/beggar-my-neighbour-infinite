package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		name string
		data bool
		exp  bool
	}{
		{
			name: "returns true",
			data: false,
			exp:  true,
		},
		{
			name: "returns false",
			data: true,
			exp:  false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			audit := solution(tc.data)
			assert.Equal(t, tc.exp, audit)
		})
	}
}
