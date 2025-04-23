package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	"strconv"
    "strings"
    "fmt"

)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) (modified bool) {
    ast.Inspect(node, func(n ast.Node) bool {
        // 1) Only look at call expressions
        call, ok := n.(*ast.CallExpr)
        if !ok {
            return true
        }
        // 2) Ensure it's expr.ParseAndEval
        sel, ok := call.Fun.(*ast.SelectorExpr)
        if !ok || sel.Sel.Name != "ParseAndEval" {
            return true
        }
        pkg, ok := sel.X.(*ast.Ident)
        if !ok || pkg.Name != "expr" {
            return true
        }
        // 3) Exactly two args
        if len(call.Args) != 2 {
            return true
        }
        // 4) Unwrap parentheses
        arg := call.Args[0]
        if p, ok := arg.(*ast.ParenExpr); ok {
            arg = p.X
        }
        // 5) Must be a string literal
        lit, ok := arg.(*ast.BasicLit)
        if !ok || lit.Kind != token.STRING {
            return true
        }
        // 6) Unquote 
        raw, err := strconv.Unquote(lit.Value)
        if err != nil {
            return true
        }
        raw = strings.TrimSpace(raw)
        // 7) Parse & simplify
        parsedExpr, err := expr.Parse(raw)
        if err != nil {
            return true
        }
        simplifiedExpr := simplify.Simplify(parsedExpr, expr.Env{})
        // 8) Format to a string
        simplifiedStr := fmt.Sprint(simplifiedExpr)
        // 9) Replace the literal in the AST
        call.Args[0] = &ast.BasicLit{
            Kind:  token.STRING,
            Value: strconv.Quote(simplifiedStr),
        }
        modified = true
        return true
    })
    return
}
func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
