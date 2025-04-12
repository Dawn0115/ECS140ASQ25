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
		updatedX := Simplify(e.X, env)
		if val, ok := updatedX.(expr.Literal); ok {
			switch e.Op{
				case '-':
					return expr.Literal(-val)
				case '+':
					return expr.Literal(val)
			}
		
		}
		return expr.Unary{Op: e.Op, X: updatedX}
	
	case expr.Binary:
		updatedX := Simplify(e.X, env)
		updatedY := Simplify(e.Y, env)
		valx, okx := updatedX.(expr.Literal); 
		valy, oky := updatedY.(expr.Literal); 
		if okx && oky {
				switch e.Op {
				case '+':
					return expr.Literal(valx + valy)
				case '-':
					return expr.Literal(valx - valy)
				case '*':
					return expr.Literal(valx * valy)
				case '/':
					if valy != 0 {
						return expr.Literal(valx / valy)
					}
				}
			}
		
		if e.Op == '+'{
			if okx && valx == 0 {
				return updatedY
			}
			if oky && valy == 0 {
				return updatedX
			}
		}
		if e.Op == '*'{
			if okx && valx == 0 {
				return expr.Literal(0)
			}
			if oky && valy == 0 {
				return expr.Literal(0)
			}
			if okx && valx == 1 {
				return updatedY
			}
			if oky && valy == 1 {
				return updatedX
			}
		}
		return expr.Binary{Op: e.Op, X: updatedX, Y: updatedY}

	
}
	panic("simplify: unreachable code reached")

}