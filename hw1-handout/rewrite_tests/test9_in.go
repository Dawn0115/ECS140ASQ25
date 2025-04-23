// test2_in.go
package main

import (
	e "hw1/expr"
	"fmt"
)

func main() {
	fmt.Println(e.ParseAndEval("1+2", e.Env{}))
}