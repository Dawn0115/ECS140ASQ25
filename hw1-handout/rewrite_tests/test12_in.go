package main

import "hw1/expr"

func nested() {
    _ = expr.ParseAndEval("1 + 2", nil)
    fn := func() {
        _ = expr.ParseAndEval("2 + 3", nil)
    }
}