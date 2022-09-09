package builder

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/GuilhermeCaruso/mooncake/moongen/models"
	"github.com/GuilhermeCaruso/mooncake/moongen/template"
	"golang.org/x/tools/imports"
)

type Builder struct {
	pkg string
}

type BuilderRef struct {
	b            *bytes.Buffer
	f            *os.File
	c            []byte
	OriginalPath string
	NewPath      string
	File         models.File
}

func (br *BuilderRef) build(pkg string) {
	br.prepareBuffer()
	br.writeHeader()
	br.writePkg(pkg)
	br.writeImports()
	br.writeMockBase()
	br.formatContent()
	br.createFile()
	br.writeFile()
}

func (br *BuilderRef) writeMockStructure(template string, i models.Implementation) {
	pWith, pWithout := i.ParamsToString()
	replacedStructure := strings.ReplaceAll(template, "%s", i.Name)
	replacedStructure = strings.ReplaceAll(replacedStructure, "%i", pWith)
	replacedStructure = strings.ReplaceAll(replacedStructure, "%k", pWithout)

	br.b.WriteString(replacedStructure)
	br.b.WriteString("\n")
}

func (br *BuilderRef) writeMethods(template string, i models.Implementation, m models.Method) {
	_, pWithout := i.ParamsToString()
	replacedMethod := strings.ReplaceAll(template, "%s", i.Name)
	replacedMethod = strings.ReplaceAll(replacedMethod, "%m", m.Name)
	replacedMethod = strings.ReplaceAll(replacedMethod, "%k", pWithout)
	replacedMethod = strings.ReplaceAll(replacedMethod, "%p", models.GetArgListString(m.Params))
	replacedMethod = strings.ReplaceAll(replacedMethod, "%u", models.GetArgGenericListString(m.Params))
	replacedMethod = strings.ReplaceAll(replacedMethod, "%r", models.GetArgListString(m.Results))
	replacedMethod = strings.ReplaceAll(replacedMethod, "%a", models.GetResultListString(m.Results))
	br.b.WriteString(replacedMethod)
	br.b.WriteString("\n")
}

func (br *BuilderRef) writeMockBase() {
	for _, i := range br.File.Implementations {
		br.writeMockStructure(template.MOCK_STRUCTURE, i)
		br.writeMockStructure(template.CONSTRUCTOR_MOCK_FUNCTION, i)
		for _, m := range i.Methods {
			br.writeMethods(template.METHOD, i, m)
		}
		br.writeMockStructure(template.CONSTRUCTOR_INTERNAL_MOCK, i)
		for _, m := range i.Methods {
			br.writeMethods(template.INTERNAL_METHOD, i, m)
		}

	}
}

func (br *BuilderRef) formatContent() {
	formattedContent, err := imports.Process(br.NewPath, br.b.Bytes(), &imports.Options{})
	if err != nil {
		log.Fatalf("something went wrong when trying to formmat file. err=%q", err.Error())
	}
	br.c = formattedContent
}

func (br *BuilderRef) createFile() {
	f, err := os.Create(br.NewPath)
	if err != nil {
		log.Fatalf("something went wrong when trying to write file. err=%q", err.Error())
	}
	br.f = f
}

func (br *BuilderRef) writeFile() {
	br.f.Write(br.c)
	br.f.Close()
}

func (br *BuilderRef) prepareBuffer() {
	var f bytes.Buffer
	br.b = &f
}

func (br *BuilderRef) reset() {
	if _, err := os.Stat(br.NewPath); errors.Is(err, os.ErrNotExist) {
		br.createFile()
		return
	}

	err := os.Remove(br.NewPath)
	if err != nil {
		log.Fatalf("something went wrong when trying to remove file. err=%q", err.Error())
	}
	br.createFile()
}

func (br *BuilderRef) writeHeader() {
	br.b.WriteString(fmt.Sprintf(template.FILE_HEADER,
		time.Now().Format("2006-01-02 15:04:05"), br.OriginalPath))
}

func (br *BuilderRef) addGeneralImports() {
	br.File.Imports = append(br.File.Imports, models.Import{
		Path: "\"reflect\"",
	}, models.Import{
		Path: "\"github.com/GuilhermeCaruso/mooncake\"",
	})
}

func (br *BuilderRef) writeImports() {
	br.addGeneralImports()

	if len(br.File.Imports) > 0 {

		var buffer bytes.Buffer

		buffer.WriteString("\nimport")

		if len(br.File.Imports) > 1 {
			buffer.WriteString(" (\n")
		}

		for _, i := range br.File.Imports {
			buffer.WriteString("\t")
			if i.Name != "" && i.Name != "<nil>" {
				buffer.WriteString(fmt.Sprintf("%s ", i.Name))
			}
			buffer.WriteString(fmt.Sprintf("%s\n", i.Path))
		}

		if len(br.File.Imports) > 1 {
			buffer.WriteString(")")
		}

		br.b.WriteString(buffer.String())
	}
}
func (br BuilderRef) writePkg(pkg string) {
	br.b.WriteString(fmt.Sprintf("package %s\n", pkg))
}

func NewBuilder(pkg string) *Builder {
	return &Builder{
		pkg: pkg,
	}
}

func (b Builder) BuildFiles(refs []BuilderRef) {
	for _, r := range refs {
		r.build(b.pkg)
	}
}
