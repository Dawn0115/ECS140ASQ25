package main

import "hw1/expr"

func demo() {
    a := expr.ParseAndEval("10 - 3", nil)
    b := expr.ParseAndEval("2*3 + 4", nil)
}