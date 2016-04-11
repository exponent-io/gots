package generate

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

func writeReflectLine(out io.Writer, fileSet *token.FileSet, n ast.Node) {
	viz := reflectVisitor{
		out:     out,
		fileSet: fileSet,
	}
	ast.Walk(&viz, n)
}

type reflectVisitor struct {
	out     io.Writer
	fileSet *token.FileSet
}

func (v *reflectVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.TypeSpec:
		if _, ok := n.Type.(*ast.StructType); !ok {
			// only process the type if it is a struct
			return nil
		}
		fmt.Fprintf(v.out, "\t%q: func() interface{} { return new(%s) }, ", n.Name, n.Name)
		fmt.Fprintf(v.out, "// %v\n", v.fileSet.Position(n.Pos()))
		return nil
	}
	return v
}
