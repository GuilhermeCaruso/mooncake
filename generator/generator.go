package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type File struct {
	Path  string   `yaml:"path"`
	Files []string `yaml:"files"`
}

func main() {
	fs := token.NewFileSet()
	file, _ := parser.ParseFile(fs, "./examples/interfaces/nested.go", nil, 0)
	// ast.Print(fs, file)
	parse(file)
}

type General struct {
	Imports    []Import
	Interfaces []Interface
}
type Import struct {
	Path string
	Name string
}

func getImport(importSpec []ast.Spec) []Import {
	mapImport := make([]Import, 0)
	for _, is := range importSpec {
		spec, ok := is.(*ast.ImportSpec)
		if !ok {
			fmt.Printf("invalid spec. err=%s", is)
			os.Exit(1)
		}
		mapImport = append(mapImport, Import{
			Path: spec.Path.Value,
			Name: spec.Name.String(),
		})
	}
	return mapImport
}

type Interface struct {
	Name       string
	Generic    []Generic
	Methods    []Method
	References []string
}

type Generic struct {
	Name string
	Type string
}

type Method struct {
	Name    string
	Args    []Arg
	Returns []Arg
}

type Arg struct {
	Name string
	Type string
}

func getReference(idt *ast.Ident) string {
	return idt.String()
}
func getMethods(methods *ast.FieldList) ([]Method, []string) {
	methodsToReturn := make([]Method, 0)
	references := make([]string, 0)

	for _, m := range methods.List {
		//maybe switch
		if ident, ok := m.Type.(*ast.Ident); ok {
			if len(m.Names) == 0 {
				references = append(references, ident.String())
				continue
			}
		}
		method := new(Method)
		method.Name = m.Names[0].String()

		if field, ok := m.Type.(*ast.FuncType); ok {
			for idx, p := range field.Params.List {
				if t, ok := p.Type.(*ast.Ident); ok {
					method.Args = append(method.Args, Arg{
						Name: fmt.Sprintf("arg%v", idx),
						Type: t.String(),
					})
				}
			}
			for idx, p := range field.Results.List {
				if t, ok := p.Type.(*ast.Ident); ok {
					method.Returns = append(method.Args, Arg{
						Name: fmt.Sprintf("result%v", idx),
						Type: t.String(),
					})
				}
			}
		}

		methodsToReturn = append(methodsToReturn, *method)
	}
	return methodsToReturn, references
}

func getInterface(typeSpec []ast.Spec) []Interface {
	interfaces := make([]Interface, 0)

	for _, ts := range typeSpec {

		typeSpec, ok := ts.(*ast.TypeSpec)
		if !ok {
			fmt.Printf("invalid spec. err=%s", ts)
			os.Exit(1)
		}

		if iface, ok := typeSpec.Type.(*ast.InterfaceType); ok {
			ifaceDetails := new(Interface)
			ifaceDetails.Name = typeSpec.Name.String()

			if typeSpec.TypeParams != nil {
				for _, tp := range typeSpec.TypeParams.List {
					var name bytes.Buffer
					for idx, n := range tp.Names {
						name.WriteString(n.String())
						if idx < len(tp.Names)-1 {
							name.WriteString(", ")
						}
					}
					ifaceDetails.Generic = append(ifaceDetails.Generic, Generic{
						Name: name.String(),
						Type: fmt.Sprintf("%s", tp.Type),
					})
				}
			}
			methods, references := getMethods(iface.Methods)

			ifaceDetails.Methods = methods
			ifaceDetails.References = references

			interfaces = append(interfaces, *ifaceDetails)
		}
	}

	return interfaces

}

func parse(f *ast.File) {
	g := new(General)
	for _, dcs := range f.Decls {
		decl, ok := dcs.(*ast.GenDecl)

		if !ok {
			fmt.Printf("invalid decl. err=%s", dcs)
			os.Exit(1)
		}

		switch decl.Tok {
		case token.IMPORT:
			g.Imports = append(g.Imports, getImport(decl.Specs)...)
		case token.TYPE:
			g.Interfaces = append(g.Interfaces, getInterface(decl.Specs)...)
		default:
			fmt.Println("not implemented")
		}

	}

	for idx, k := range g.Interfaces {
		for _, r := range k.References {
			for _, ik := range g.Interfaces {
				if r == ik.Name {
					g.Interfaces[idx].Methods = append(g.Interfaces[idx].Methods, ik.Methods...)
				}
			}
		}

	}

	b, _ := json.Marshal(g)
	fmt.Println(string(b))
}
