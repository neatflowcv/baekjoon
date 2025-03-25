package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	t.Parallel()
	tt := []struct {
		input *Input
		want  bool
	}{
		{
			input: &Input{
				n:     7,
				k:     3,
				left:  "1110110",
				right: "1011001",
			},
			want: true,
		},
		{
			input: &Input{
				n:     6,
				k:     2,
				left:  "110101",
				right: "011001",
			},
			want: false,
		},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			t.Parallel()

			ret := Solution(tc.input)

			require.Equal(t, tc.want, ret)
		})
	}
}
