package generate

import "go/ast"

type FieldSpec struct {
	fieldName string
	optional  string
	typeName  string
	jsonName  string
}

func fieldSpecs(n ast.Node) []FieldSpec {
	var viz fieldVisitor
	ast.Walk(&viz, n)
	return viz.fields
}

type fieldVisitor struct {
	fields []FieldSpec
}

func (v *fieldVisitor) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.Field:
		var fld FieldSpec
		fld.fieldName = n.Names[0].Name
		fld.jsonName = fld.fieldName

		fld.typeName = typeName(n.Type)

		tag := ""
		if n.Tag != nil && len(n.Tag.Value) > 2 {
			var tagOpts tagOptions

			tag = n.Tag.Value[1 : len(n.Tag.Value)-1]
			tag = getStructTag(tag, "json")
			fld.jsonName, tagOpts = parseTag(tag)

			if fld.jsonName == "-" {
				return nil
			} else if tagOpts.Contains("omitempty") {
				fld.optional = "?"
			}
		}

		v.fields = append(v.fields, fld)
		return nil
	}
	return v
}
