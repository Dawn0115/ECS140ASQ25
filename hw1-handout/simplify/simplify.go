package simplify

import (
	"hw1/expr"
)

// Simplify should return the simplified expresion
func Simplify(e expr.Expr, env expr.Env) expr.Expr {
	//TODO implement the simplify
	switch e := e.(type) {
	case expr.Var:
		if val, ok := env[e]; ok {
			return expr.Literal(val)
		}
		return e
	case expr.Literal:
		return e
	case expr.Unary:
		updagedX := Simplify(e.X, env)

	panic("TODO: implement this!")
}
