package main

import (
	"encoding/json"
	"fmt"

	"github.com/GuilhermeCaruso/mooncake/generator/models"
	"github.com/GuilhermeCaruso/mooncake/generator/parser"
)

func main() {
	generator := new(models.Generator)
	parseFiles(generator)
	generator.PrepareNested()

	b, _ := json.Marshal(generator.Files)
	fmt.Println(string(b))
}

func parseFiles(generator *models.Generator) {
	filesToGenerate := []string{
		"./examples/interfaces/nested.go",
		"./examples/interfaces/simple.go",
		"./examples/interfaces/reference.go",
	}

	queue := make(chan models.File)
	defer close(queue)

	p := parser.NewParser()
	for _, ftg := range filesToGenerate {
		go func(fp string) {
			queue <- p.Parse(fp)
		}(ftg)
	}
	for range filesToGenerate {
		generator.RegisterFile(<-queue)
	}
}
