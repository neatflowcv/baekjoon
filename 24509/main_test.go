package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name  string
		input *Input
		want  []int
	}{
		{
			name: "case 1",
			input: &Input{
				n: 4,
				persons: []*Person{
					{
						no:  1,
						arr: []int{49, 81, 74, 66},
					},
					{
						no:  2,
						arr: []int{88, 70, 91, 38},
					},
					{
						no:  3,
						arr: []int{94, 80, 85, 79},
					},
					{
						no:  4,
						arr: []int{91, 27, 44, 60},
					},
				},
			},
			want: []int{3, 1, 2, 4},
		},
		{
			name: "case 2",
			input: &Input{
				n: 8,
				persons: []*Person{
					{
						no:  3,
						arr: []int{65, 21, 8, 61},
					},
					{
						no:  8,
						arr: []int{69, 2, 19, 58},
					},
					{
						no:  7,
						arr: []int{85, 40, 42, 29},
					},
					{
						no:  4,
						arr: []int{45, 1, 58, 56},
					},
					{
						no:  2,
						arr: []int{4, 37, 83, 33},
					},
					{
						no:  6,
						arr: []int{75, 60, 12, 42},
					},
					{
						no:  5,
						arr: []int{80, 20, 74, 49},
					},
					{
						no:  1,
						arr: []int{8, 26, 93, 33},
					},
				},
			},
			want: []int{7, 6, 1, 3},
		},
		{
			name: "case 3",
			input: &Input{
				n: 4,
				persons: []*Person{
					{
						no:  1,
						arr: []int{0, 0, 0, 0},
					},
					{
						no:  2,
						arr: []int{0, 0, 0, 0},
					},
					{
						no:  3,
						arr: []int{0, 0, 0, 0},
					},
					{
						no:  4,
						arr: []int{0, 0, 0, 0},
					},
				},
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ret := Solution(tc.input)

			require.Equal(t, tc.want, ret)
		})
	}
}
