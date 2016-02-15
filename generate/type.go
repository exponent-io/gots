package generate

import "go/ast"

func typeName(n ast.Node) string {
	var viz typeVisitor
	ast.Walk(&viz, n)
	return viz.typeName
}

type typeVisitor struct {
	typeName string
}

func (v *typeVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.ArrayType:
		v.typeName = tsType(typeName(n.Elt)) + "[]"
		return nil
	case *ast.Ident:
		v.typeName = tsType(n.Name)
		return nil
	}
	return v
}

// translate native goTypes into TypeScript/Javascript types
func tsType(goType string) string {

	switch goType {
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "interface{}":
		return "any"
	case "int", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64", "float32", "float64":
		return "number"
	}
	return goType
}
