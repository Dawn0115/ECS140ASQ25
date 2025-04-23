package main

import "hw1/expr"

func parseError() {
    // "1 +" is invalid syntax → Parse fails, must stay the same
    _ = expr.ParseAndEval("1 +", nil)
}