package generate

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strings"
)

func writeStruct(out io.Writer, fileSet *token.FileSet, n ast.Node) {
	viz := structVisitor{
		out:     out,
		fileSet: fileSet,
	}
	ast.Walk(&viz, n)
}

type structVisitor struct {
	out     io.Writer
	fileSet *token.FileSet
}

func (v *structVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.TypeSpec:
		fmt.Fprintf(v.out, "// %v\n", v.fileSet.Position(n.Pos()))
		if n.Doc != nil {
			for _, c := range n.Doc.List {
				fmt.Fprintf(v.out, "// %v\n", c.Text)
			}
		}
		fmt.Fprintf(v.out, "export interface %v {\n", n.Name)

		fields := fieldSpecs(n.Type)
		writeFieldSpecs(v.out, fields)

		fmt.Fprintf(v.out, "}\n\n")
		return nil
	}
	return v
}

func writeFieldSpecs(w io.Writer, fields []FieldSpec) {
	var nameWidth, typeWidth int
	for _, f := range fields {
		jsonWidth := len(f.jsonName) + len(f.optional)
		if jsonWidth > nameWidth {
			nameWidth = jsonWidth
		}
		if len(f.typeName) > typeWidth {
			typeWidth = len(f.typeName)
		}
	}
	for _, f := range fields {
		fmt.Fprintf(w, "  %v%v:%v %v;%v // %v\n",
			f.jsonName, f.optional, strings.Repeat(" ", nameWidth-len(f.jsonName)-len(f.optional)),
			f.typeName, strings.Repeat(" ", typeWidth-len(f.typeName)),
			f.fieldName)
	}
}
