package rewrite

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	"strconv"

)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
    ast.Inspect(node, func(n ast.Node) bool {
        // 1) Only look at call expressions
        call, ok := n.(*ast.CallExpr)
        if !ok {
            return true
        }

        // 2) Ensure the call is expr.ParseAndEval
        sel, ok := call.Fun.(*ast.SelectorExpr)
        if !ok || sel.Sel.Name != "ParseAndEval" {
            return true
        }
        pkg, ok := sel.X.(*ast.Ident)
        if !ok || pkg.Name != "expr" {
            return true
        }

        //two arguments
        if len(call.Args) != 2 {
            return true
        }

        // 4) First arg must be a STRING literal
        lit, ok := call.Args[0].(*ast.BasicLit)
        if !ok || lit.Kind != token.STRING {
            return true
        }

        // 5) Unquote the literal: 
        raw, err := strconv.Unquote(lit.Value)
        if err != nil {
            return true
        }
        spacedRaw := ""
        for i := 0; i < len(raw); i++ {
            c := raw[i]
            if c == '+' || c == '-' || c == '*' || c == '/' {
                // Add space before operator (if not at start or after another operator/paren)
                if i > 0 && raw[i-1] != ' ' && raw[i-1] != '+' && raw[i-1] != '-' && raw[i-1] != '*' && raw[i-1] != '/' && raw[i-1] != '(' {
                    spacedRaw += " "
                }
                spacedRaw += string(c)
                // Add space after operator (if not at end or before another operator/paren)
                if i < len(raw)-1 && raw[i+1] != ' ' && raw[i+1] != '+' && raw[i+1] != '-' && raw[i+1] != '*' && raw[i+1] != '/' && raw[i+1] != ')' {
                    spacedRaw += " "
                }
            } else {
                spacedRaw += string(c)
            }
        }
        // 6) Parse it into an expr.Expr
        parsedExpr, err := expr.Parse(raw)
        if err != nil {
            return true
        }

        // 
        simplifiedExpr := simplify.Simplify(parsedExpr, expr.Env{})

        // 8) 
        simplifiedStr := fmt.Sprint(simplifiedExpr)

        // 9) 
        call.Args[0] = &ast.BasicLit{
            Kind:  token.STRING,
            Value: strconv.Quote(simplifiedStr),
        }

        return true
    })
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
