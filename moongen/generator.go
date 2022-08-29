package main

import (
	"flag"
	"fmt"

	"github.com/GuilhermeCaruso/mooncake/moongen/builder"
	"github.com/GuilhermeCaruso/mooncake/moongen/config"
	"github.com/GuilhermeCaruso/mooncake/moongen/parser"
)

const (
	helper = `
mooncake is an simple way to generate mocks based on interface implementation.

To generate mocked files is necessary follow two basic steps.

First: create mooncake.yaml file
	mock:
	package: mocks
	path: <interface_location>
	files:
		- <interface_file>.go
	output: <output_location>
	prefix: <optional_prefix>

Second: run
	moongen --file mooncake.yaml
`
)

var (
	file = flag.String("file", "mooncake.yaml", "mooncake configuration file path")
)

func init() {
	flag.Usage = func() {
		fmt.Print(helper)
	}
	flag.Parse()
}

func main() {
	mf := config.NewConfig(*file)
	mf.PrepareFolder()
	parsedFiles := parseFiles(mf.Files)
	builder.NewBuilder(mf.Package).BuildFiles(parsedFiles)
}

func parseFiles(f []config.ConfigFile) []builder.BuilderRef {
	filesToGenerate := f

	queue := make(chan builder.BuilderRef)
	defer close(queue)

	files := make([]builder.BuilderRef, 0)

	p := parser.NewParser()
	for _, ftg := range filesToGenerate {
		go func(fp config.ConfigFile) {
			result := p.Parse(fp.Original)
			queue <- builder.BuilderRef{
				OriginalPath: fp.Original,
				NewPath:      fp.New,
				File:         result,
			}
		}(ftg)
	}
	for range filesToGenerate {
		files = append(files, <-queue)
	}

	return p.PrepareNested(files)
}
