package gots

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

var nativeTypes map[reflect.Kind]string = map[reflect.Kind]string{
	reflect.Bool:    "boolean",
	reflect.Int:     "number",
	reflect.Int8:    "number",
	reflect.Int16:   "number",
	reflect.Int32:   "number",
	reflect.Int64:   "number",
	reflect.Uint:    "number",
	reflect.Uint8:   "number",
	reflect.Uint16:  "number",
	reflect.Uint32:  "number",
	reflect.Uint64:  "number",
	reflect.Float32: "number",
	reflect.Float64: "number",
	reflect.String:  "string",
}

type byName []reflect.Type

func (n byName) Len() int           { return len(n) }
func (n byName) Less(i, j int) bool { return n[i].Name() < n[j].Name() }
func (n byName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func Convert(w io.Writer, vals ...interface{}) error {
	// loop through and resolve all dependencies first
	deps := map[reflect.Type]bool{}
	for _, i := range vals {
		t := reflect.TypeOf(i)
		resolveTypeDeps(t, deps)
	}

	types := []reflect.Type{}
	for t, _ := range deps {
		types = append(types, t)
	}
	sort.Sort(byName(types))
	for _, t := range types {
		err := writeType(t, w)
		if err != nil {
			return err
		}
	}
	return nil
}

func ConvertToString(vals ...interface{}) (string, error) {
	buf := &bytes.Buffer{}
	err := Convert(buf, vals...)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func resolveTypeDeps(t reflect.Type, deps map[reflect.Type]bool) {
	if _, ok := deps[t]; ok {
		return
	} else {
		deps[t] = true
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := tag(f)
		if !tag.exclude {
			if f.Type.Kind() == reflect.Struct {
				resolveTypeDeps(f.Type, deps)
			} else if f.Type.Kind() == reflect.Slice {
				if f.Type.Elem().Kind() == reflect.Struct {
					resolveTypeDeps(f.Type.Elem(), deps)
				}
			}
		}
	}
}

func writeType(t reflect.Type, w io.Writer) error {

	fmt.Fprintf(w, "\nexport interface %s {\n", t.Name())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		tag := tag(f)
		if !tag.exclude {
			err := writeField(w, tag.name, typeToString(f.Type))
			if err != nil {
				return err
			}
		}
	}

	fmt.Fprintf(w, "}\n")
	return nil
}

func typeToString(t reflect.Type) string {
	k := t.Kind()
	switch k {
	case reflect.Struct:
		return t.Name()
	case reflect.Slice:
		return typeToString(t.Elem()) + "[]"
	default:
		if tm, ok := nativeTypes[k]; ok {
			return tm
		}
	}
	panic(fmt.Sprintf("don't know how to handle type %v", t))
}

func writeField(w io.Writer, name string, typ string) error {
	_, err := fmt.Fprintf(w, "  %s: %s;\n", name, typ)
	return err
}

type jsonTag struct {
	name      string
	exclude   bool
	omitempty bool
}

func tag(f reflect.StructField) (j jsonTag) {
	tag := f.Tag.Get("json")
	if tag == "" {
		j.name = f.Name
		j.exclude = unicode.IsLower(rune(f.Name[0]))
		return
	}

	s := strings.Split(tag, ",")
	if len(s) > 0 {
		j.name = s[0]
		j.exclude = j.name == "-" || j.name == ""

		for _, p := range s[1:] {
			if strings.TrimSpace(p) == "omitempty" {
				j.omitempty = true
			}
		}
	}
	return
}
