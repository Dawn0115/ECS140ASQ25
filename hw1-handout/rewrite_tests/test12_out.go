package main

import "hw1/expr"

func nested() {
    _ = expr.ParseAndEval("3", nil)
    fn := func() {
        _ = expr.ParseAndEval("5", nil)
    }
}