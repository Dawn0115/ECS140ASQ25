package depth
import "fmt"
import (
	"hw1/expr"
)

// Depth should return the maximum number of AST nodes between the root of the
// tree and any leaf (literal or variable) in the tree.
func Depth(e expr.Expr) uint {
	switch v := e.(type) {
	case expr.Literal, expr.Var:
		// A literal or variable is a leaf node so its depth is 1.
		return 1
	case expr.Unary:
		// For a unary expression, depth is 1 plus the depth of its single operand.
		return 1 + Depth(v.X)
	case expr.Binary:
		// For a binary expression, depth is 1 plus the maximum depth of its operands.
		left := Depth(v.X)
		right := Depth(v.Y)
		if left > right {
			return 1 + left
		}
		return 1 + right
	default:
		// If the type is not supported, panic as the test expects.
		panic(fmt.Sprintf("Depth: unexpected expression type %T", e))
	}
}
	
