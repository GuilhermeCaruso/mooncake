package main

import (
	"github.com/GuilhermeCaruso/mooncake/generator/builder"
	"github.com/GuilhermeCaruso/mooncake/generator/config"
	"github.com/GuilhermeCaruso/mooncake/generator/parser"
)

func main() {
	mf := config.NewConfig("./examples/mooncake.yaml")
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
