package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"github.com/GuilhermeCaruso/mooncake/generator/models"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(fp string) models.File {

	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, fp, nil, 0)

	fmt.Println(file, err)
	return models.File{}
}
