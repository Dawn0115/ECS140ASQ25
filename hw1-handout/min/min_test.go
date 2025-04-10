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
	// TODO add more tests for 100% test coverage
}

func TestMin(t *testing.T) {
	for i, test := range tests {
		m := Min(test.in)
		if m != test.out {
			t.Errorf("#%d: Min(%v)=%d; want %d", i, test.in, m, test.out)
		}
	}
}
