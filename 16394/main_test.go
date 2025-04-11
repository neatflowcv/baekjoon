package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name     string
		input    *Input
		expected int
	}{
		{
			name: "case 1",
			input: &Input{
				n: 2018,
			},
			expected: 72,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ret := Solution(tc.input)

			require.Equal(t, tc.expected, ret)
		})
	}
}
