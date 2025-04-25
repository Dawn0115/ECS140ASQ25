package simplify

import (
	"hw1/expr"
	"testing"
)

//!+Simplify
func TestSimplify(t *testing.T) {
	tests := []struct {
		expr string
		env  expr.Env
		want string
	}{
		{"5 + 2", expr.Env{}, "7"},
		{"X", expr.Env{}, "X"},
		{"X", expr.Env{"X": 2}, "2"},
		{"-X", expr.Env{"X": 2}, "-2"},
		{"-X", expr.Env{}, "(-X)"},
		{"X", expr.Env{}, "X"},
		{"--X", expr.Env{"X": 2}, "2"},
		{"Y", expr.Env{"X": 2}, "Y"},
		{"Y*0", expr.Env{"X": 2}, "0"},
		{"Y+0", expr.Env{"X": 2}, "Y"},
		{"Y*1", expr.Env{"X": 2}, "Y"},
		{"0*Y", expr.Env{"X": 2}, "0"},
		{"0*Y", expr.Env{"X": 2}, "0"},
		{"0+Y", expr.Env{"X": 2}, "Y"},
		{"1*Y", expr.Env{"X": 2}, "Y"},
		{"1*Y*1", expr.Env{"X": 2}, "Y"},
		{"X+2", expr.Env{"X": 2}, "4"},
		{"+5", expr.Env{}, "5"},
		
		{"10 / X", expr.Env{"X": 2}, "5"},
		{"X*Y", expr.Env{"X":2, "Y": 3}, "6"},
		{"X+Y", expr.Env{"X":2, "Y": 3}, "5"},
		{"Y+2", expr.Env{"X": 2}, "(Y + 2)"},
		{"3 + 5 + X", expr.Env{"X": 2}, "10"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"2 + X + 3", expr.Env{}, "((2 + X) + 3)"},
		{"X + 2 + 3", expr.Env{}, "((X + 2) + 3)"},
		{"2 + 3 + X", expr.Env{}, "(5 + X)"},
		{"(X + X) - Y", expr.Env{"X": 2}, "(4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"-(X + X) - Y", expr.Env{"X": 2}, "(-4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 8}, "((X + X) - 8)"},
		{"10 - 1 + X - Y", expr.Env{}, "((9 + X) - Y)"},
		{"X + 3 + 5", expr.Env{}, "((X + 3) + 5)"},
		{"5 + 2", nil, "7"},
		{"+5", nil, "5"},
		{"-5", nil, "-5"},
		{"--5", nil, "5"},
		{"++5", nil, "5"},
		{"+-5", nil, "-5"},
		{"-+5", nil, "-5"},

		// --- variables and env lookup ---
		{"X", nil, "X"},
		{"X", expr.Env{"X": 2}, "2"},
		{"Y", expr.Env{"X": 2}, "Y"},

		// --- unary over var, with & without env ---
		{"-X", nil, "(-X)"},
		{"-X", expr.Env{"X": 3}, "-3"},
		{"--X", expr.Env{"X": 4}, "4"},
		{"+-X", nil, "(+(-X))"},
		{"-+X", nil, "(-(+X))"},

		// --- pure binary literal ops ---
		{"10 - 1 + 3", nil, "12"},
		{"2 * 3 + 4", nil, "10"},
		{"10 / 2", nil, "5"},
		{"10 / 3", nil, "3"},             // integer division semantics
		{"10 / 0", nil, "(10 / 0)"},      // divide-by-zero fall-through
		{"0 / 5", nil, "0"},
		{"0 / X", expr.Env{"X": 5}, "0"},
		{"X / 0", nil, "(X / 0)"},

		// --- identity rules for + and * ---
		{"X + 0", nil, "X"},
		{"0 + X", nil, "X"},
		{"Y + 0", nil, "Y"},
		{"0 + Y", nil, "Y"},
		{"X * 1", nil, "X"},
		{"1 * X", nil, "X"},
		{"Y * 1", nil, "Y"},
		{"1 * Y", nil, "Y"},
		{"X * 0", nil, "0"},
		{"0 * X", nil, "0"},
		{"Y * 0", nil, "0"},
		{"0 * Y", nil, "0"},

		// --- combinations of identity rules ---
		{"(5 + 0) * 1", nil, "5"},
		{"0 * (X + 2)", nil, "0"},
		{"(X + 0) * Y", nil, "(X * Y)"},
		{"(X * 1) + 0", nil, "X"},

		// --- nested binary/unary with env ---
		{"(X + X) - Y", expr.Env{"X": 2}, "(4 - Y)"},
		{"(X + X) - Y", expr.Env{"Y": 1}, "((X + X) - 1)"},
		{"-(X + 2)", expr.Env{"X": 3}, "-5"},
		{"-(X + 2)", nil, "(-(X + 2))"},
		{"-((X * 1) + 0)", expr.Env{"X": 7}, "-7"},

		// --- deeper nesting & mixed operators ---
		{"1 + (2 + (3 + 4))", nil, "10"},
		{"-(1 + (2 * (3 + 0)))", nil, "-7"},
		{"(X + Y) * (0 + Z)", expr.Env{"X": 1, "Y": 2, "Z": 3}, "9"},
		{"(X + Y) * (0 + Z)", nil, "((X + Y) * Z)"},

		// --- non-simplifiable mix of vars & literals ---
		{"X + 1", nil, "(X + 1)"},
		{"1 + X", nil, "(1 + X)"},
		{"X - 1", nil, "(X - 1)"},
		{"1 - X", nil, "(1 - X)"},
		{"X * Y", nil, "(X * Y)"},
		{"X / Y", nil, "(X / Y)"},
	}

	for _, test := range tests {
		e, err := expr.Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		// Run the method
		result := Simplify(e, test.env)

		// Display the result
		got := expr.Format(result)

		// Check the result
		if got != test.want {
			t.Errorf("Simplify(%s) in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

type Foo int

func (Foo) Eval(env expr.Env) float64 {
	return 0.0
}

func (Foo) Check(vars map[expr.Var]bool) error {
	return nil
}
func TestSimplify_Fail(t *testing.T) {

	func() {
		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()

		var f Foo
		Simplify(f, expr.Env{})
	}()

}

//!-Simplify
