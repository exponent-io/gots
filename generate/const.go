package generate

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
	"strings"
)

func writeConst(out io.Writer, fileSet *token.FileSet, n ast.Node) {
	// ast.Print(fileSet, n)
	// return
	viz := constVisitor{
		out:     out,
		fileSet: fileSet,
	}
	ast.Walk(&viz, n)
}

type constVisitor struct {
	out     io.Writer
	fileSet *token.FileSet
}

func (v *constVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.GenDecl:
		if n.Tok == token.CONST {
			fmt.Fprintf(v.out, "// %v\n", v.fileSet.Position(n.Pos()))

			var viz constDeclVisitor
			ast.Walk(&viz, n)

			done := make(map[string]struct{})
			for _, t := range viz.consts {
				if _, ok := done[t.constType]; !ok && t.constType != "" {
					fmt.Fprintf(v.out, "export type %v = string;\n", t.constType)
					done[t.constType] = struct{}{}
				}
			}
			fmt.Fprintf(v.out, "\n")
			writeConstSpecs(v.out, viz.consts)
			fmt.Fprintf(v.out, "\n")
		}
		return nil
	}
	return v
}

type ConstSpec struct {
	constName  string
	constType  string
	constValue string
}

type constDeclVisitor struct {
	consts   []ConstSpec
	lastType string
}

func (v *constDeclVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.ValueSpec:
		name := n.Names[0].String()
		if tv, ok := n.Type.(*ast.Ident); ok {
			v.lastType = tv.Name
		}
		if lv, ok := n.Values[0].(*ast.BasicLit); ok {
			if lv.Kind == token.STRING {
				if v.lastType == "" {
					v.lastType = "string"
				}
				v.consts = append(v.consts, ConstSpec{
					constName:  name,
					constType:  v.lastType,
					constValue: lv.Value,
				})
			}
		}

		return nil
	}
	return v
}

func writeConstSpecs(w io.Writer, consts []ConstSpec) {
	var nameWidth int
	for _, c := range consts {
		w := len(c.constName) + len(c.constType) + 1
		if w > nameWidth {
			nameWidth = w
		}
	}
	for _, c := range consts {
		fmt.Fprintf(w, "export const %s:%s%s = %s;\n",
			c.constName, c.constType, strings.Repeat(" ", nameWidth-len(c.constName)-len(c.constType)-1),
			c.constValue,
		)
	}
}
