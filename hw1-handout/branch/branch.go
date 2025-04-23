package branch

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// branchCount should count the number of branching statements in the function.
func branchCount(fn *ast.FuncDecl) uint {
    var count uint

    // ast.Inspect traverses fn.Body in depth-first order.
    ast.Inspect(fn.Body, func(n ast.Node) bool {
        if n == nil {
            return false // travel down to the bottom
        }

		if _, isLit := n.(*ast.FuncLit); isLit {
            return false
        }


      	
        switch n.(type) {
        case *ast.IfStmt:         
            count++
        case *ast.ForStmt:        
            count++
        case *ast.RangeStmt:      
            count++
        case *ast.SwitchStmt:    
            count++
        case *ast.TypeSwitchStmt: 
            count++
        case *ast.SelectStmt:     
            count++
        }

        return true // keep traversing  
    })

    return count
}
// ComputeBranchFactors returns a map from the name of the function in the given
// Go code to the number of branching statements it contains.
func ComputeBranchFactors(src string) map[string]uint {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	m := make(map[string]uint)
	for _, decl := range f.Decls {
		switch fn := decl.(type) {
		case *ast.FuncDecl:
			m[fn.Name.Name] = branchCount(fn)
		}
	}

	return m
}
