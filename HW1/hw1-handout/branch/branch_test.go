package branch

import (
	// "fmt"
	"testing"
)

//!+ComputeBranchFactors
func TestComputeBranchFactors(t *testing.T) {
	var test_code = `
	package main

	import (
		"fmt"
		"eval"
	)

	func sequential_if() {
		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}

		if true {
			fmt.Println("is true")
		} else {
			fmt.Println("isn't true")
		}
	}

	func single_for() {
		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}

		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}

		for i := 0; i < 10; i += 1 {
			fmt.Println("is loop")
		}
	}

	func single_range() {
		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}

		for i := range []int{1, 1, 2, 3, 5, 8} {
			fmt.Println("is fibonacci")
		}
	}

	func no_branches() {
		return 42
	}

	func single_switch() {
		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		default:
			fmt.Println("It isn't five...")
		}

		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		default:
			fmt.Println("It isn't five...")
		}
	}

	func single_typeswitch() {
		var x interface{} = "test"
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		default:
			fmt.Println("It's not a string...")
		}
	}

	func nested_if(x int) float64 {
		var result float64
		if x < 0 {
			if x > -5 {
				result := -0.5
			} else {
				result := -1
			}
		} else if x > 0 {
			result := 1
		} else {
			result := 0
		}
		return result
	}

	func nested_for_if() {
		for i := 0; i < 10; i += 1 {
			if i > 5 {
				fmt.Println("is filter")
			}
		}
	}

	func nested_switch_if(x int) {
		switch x > 5 {
		case true:
			if x > 10 {
				fmt.Println("is really big")
			}
		default:
			fmt.Println("is default")
		}
	}

	func mixed_switch_no_default_for_if() {
		switch 5 {
		case 0:
			// pass
		case 5:
			fmt.Println("It's five!")
		}

		for i := 0; i < 10; i += 1 {
			if i > 5 {
				fmt.Println("is filter")
			}
		}
	}

	func single_typeswitch_no_default() {
		var x interface{} = "test"
		switch x.(type) {
		case uint:
			// pass
		case string:
			fmt.Println("It's a string!")
		}
	}
	


	func nested_if_no_else(x int) float64 {
		var result float64 = 0
		if x < 0 {
			if x > -5 {
				result := -0.5
			} else {
				result := -1
			}
		} else if x > 0 {
			result := 1
		}
		return result
	}
	func with_selects() {
    select {
    case a <- ch:
    case b := <-ch2:
    default:
    }
    select {
    case <-ch3:
    }
}

	func with_closure() {
		if cond1 {
		}
		fn := func() {
			// these three should *not* be counted:
			if cond2 {}
			for i := range []int{1, 2} {}
			select { case <-ch: }
		}
		for range []int{1, 2, 3} {
		}
	}

	func infinite_for() {
		for {
		}
	}

	func nested_closure() {
		outer := func() {
			if innerCond {
			}
			inner := func() {
				if deepCond {
				}
			}
		}
		if topCond {
		}
	}
	type P struct{}

	func (p P) Method() {
		// three IfStmt: if, else-if, nested if
		if condA {
		} else if condB {
		} else {
			if condC {
			}
		}
		// two nested SwitchStmt
		switch v := some(); v {
		case 1:
			switch t := other(); t {
			case "x":
			case "y":
			}
		case 2:
		}
	}

	func (p *P) PtrMethod() {
		var x interface{}
		// one TypeSwitchStmt
		switch x.(type) {
		case int:
		case string:
		default:
		}
	}

	func ClosureTest() {
		// one top-level IfStmt
		if top {
		}
		// these must be ignored (FuncLit)
		go func() {
			if inner {}
			for {}
			select { case <-ch: }
		}()
	}

	func ComplexLoop() {
		// one ForStmt
		for i := 0; i < 3; i++ {
			// one SelectStmt
			select {
			case <-ch1:
			case ch2 <- v:
			}
			// one RangeStmt + one nested IfStmt
			for idx := range arr {
				if filter(idx) {
				}
			}
		}
	}

	func OnlySelect() {
		// bare SelectStmt
		select {}
	}

	func OnlyTypeSwitch() {
		var x interface{}
		// bare TypeSwitchStmt without default
		switch x.(type) {
		case bool:
		case float64:
	}
	`

	tests := []struct {
		name     string
		branches uint
	}{
		{"sequential_if", 4},
		{"single_for", 3},
		{"single_range", 5},
		{"no_branches", 0},
		{"single_switch", 2},
		{"single_typeswitch", 3},
		{"nested_if", 3},
		{"nested_for_if", 2},
		{"nested_switch_if", 2},
		{"mixed_switch_no_default_for_if", 3},
		{"single_typeswitch_no_default", 1},
		{"nested_if_no_else", 3},
		{"with_selects", 2},
		{"with_closure", 2},
		{"infinite_for", 1},
		{"nested_closure", 1},
		{"Method", 5},
		{"PtrMethod", 1},
		{"ClosureTest", 1},
		{"ComplexLoop", 4},
		{"OnlySelect", 1},
		{"OnlyTypeSwitch", 1},
	}

	branch_factors := ComputeBranchFactors(test_code)

	for _, test := range tests {
		if branch_factors[test.name] != test.branches {
			t.Errorf("ComputeBranchFactors(%v) = %d, want %d\n",
				test.name, branch_factors[test.name], test.branches)
		}
	}
}

//!-ComputeBranchFactors

func TestComputeBranchFactors_Fail(t *testing.T) {

	func() {

		defer func() {
			if recover() == nil {
				t.Errorf("did not panic, but should\n")
			}
		}()
		ComputeBranchFactors("not a valid go program")
	}()

}
