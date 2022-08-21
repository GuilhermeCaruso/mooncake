package parser

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"github.com/GuilhermeCaruso/mooncake/generator/models"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(fp string) models.File {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, fp, nil, 0)

	if err != nil {
		log.Fatalf("Unable to find file: %s", err.Error())
	}

	return p.parse(file)
}

func (p Parser) parse(f *ast.File) models.File {
	newFile := new(models.File)
	for _, decl := range f.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			log.Fatalf("Invalid declaration: %s", decl)
		}
		switch d.Tok {
		case token.IMPORT:
			newFile.Imports = append(newFile.Imports, p.getImports(d.Specs)...)
		case token.TYPE:
			newFile.Implementations = append(newFile.Implementations, p.getImplementations(d.Specs)...)
		}
	}
	return *newFile
}

func (p Parser) getImports(is []ast.Spec) []models.Import {
	mi := make([]models.Import, 0)

	for _, i := range is {
		spec, ok := i.(*ast.ImportSpec)
		if !ok {
			log.Fatalf("Invalid import spec: %s", i)
		}
		mi = append(mi, models.Import{Path: spec.Path.Value, Name: spec.Name.String()})
	}
	return mi
}

func (p Parser) getImplementations(is []ast.Spec) []models.Implementation {
	implementations := make([]models.Implementation, 0)

	for _, s := range is {
		ts, ok := s.(*ast.TypeSpec)
		if !ok {
			log.Fatalf("Invalid type spec: %v", ts)
		}

		if it, ok := ts.Type.(*ast.InterfaceType); ok {
			implementation := new(models.Implementation)
			implementation.Name = ts.Name.String()
			if ts.TypeParams != nil {
				implementation.Params = p.getParams(ts.TypeParams.List)
			}
			implementation.Methods,
				implementation.References = p.getMethods(it.Methods)
			implementations = append(implementations, *implementation)
		}
	}
	return implementations
}

func (p Parser) getParams(fl []*ast.Field) []models.Param {
	params := make([]models.Param, 0)
	for _, f := range fl {
		var name bytes.Buffer
		for idx, n := range f.Names {
			name.WriteString(n.String())
			if idx < len(f.Names)-1 {
				name.WriteString(", ")
			}
		}
		params = append(params, models.Param{
			Name: name.String(), Type: fmt.Sprintf("%s", f.Type)})
	}
	return params
}

func (p Parser) getMethods(fl *ast.FieldList) ([]models.Method, []string) {
	ms := make([]models.Method, 0)
	refs := make([]string, 0)

	for _, m := range fl.List {
		switch m.Type.(type) {
		case *ast.Ident:
			if len(m.Names) == 0 {
				refs = append(refs, m.Type.(*ast.Ident).String())
			}
		case *ast.FuncType:
			method := new(models.Method)
			method.Name = m.Names[0].String()
			f := m.Type.(*ast.FuncType)
			method.Params = p.getArgs("param", f.Params.List)
			method.Results = p.getArgs("result", f.Results.List)
			ms = append(ms, *method)
		}
	}
	return ms, refs
}

func (p Parser) getArgs(key string, f []*ast.Field) []models.Arg {
	args := make([]models.Arg, 0)
	for idx, p := range f {
		switch p.Type.(type) {
		case *ast.Ident:
			args = append(args, models.Arg{
				Name: fmt.Sprintf("%s%v", key, idx),
				Type: p.Type.(*ast.Ident).String(),
			})
		case *ast.SelectorExpr:
			se := p.Type.(*ast.SelectorExpr)
			args = append(args, models.Arg{
				Name: fmt.Sprintf("%s%v", key, idx),
				Type: fmt.Sprintf("%s.%s", se.X, se.Sel),
			})
		}
	}
	return args
}
