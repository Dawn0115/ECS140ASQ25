package min

import "testing"

type Test struct {
	in  []int
	out int
}

var tests = []Test{
	{
		in:  []int{-1, 0, 1, 2, 4},
		out: -1,
	},
	{
		in:  []int{1},
		out: 1,
	},
	{
		in:  []int{},
		out: 0,
	},
	{
		in:  nil,
		out: 0,
	},
	{
		in:  []int{1,2,3,4},
		out: 1, // min stays
	},
	{
		in:  []int{1,1,1,1},
		out: 1, // all same value
	},
	{
		in:  []int{8,-3,2,-10},
		out: -10, // min updated multiple times
	},
	{in: []int{0},               out: 0},

	// single-element negative
	{in: []int{-5},              out: -5},

	// descending-order: multiple positive updates
	{in: []int{4, 3, 2, 1},      out: 1},

	// min in the middle (only one update)
	{in: []int{5, 1, 5},          out: 1},
	{in: []int{3, 4, 2},          out: 2},

	// all-equal negatives
	{in: []int{-2, -2, -2},      out: -2},

	// mixture of large pos and neg
	{in: []int{1000000, -1000000, 0}, out: -1000000},
}

func TestMin(t *testing.T) {
	for i, test := range tests {
		m := Min(test.in)
		if m != test.out {
			t.Errorf("#%d: Min(%v)=%d; want %d", i, test.in, m, test.out)
		}
	}
}
