package builder

import (
	"fmt"
	"io/ioutil"

	"github.com/GuilhermeCaruso/mooncake/generator/models"
	"github.com/GuilhermeCaruso/mooncake/generator/template"
)

type Builder struct {
	pkg string
}

type BuilderRef struct {
	OriginalPath string
	NewPath      string
	File         models.File
}

func NewBuilder(pkg string) *Builder {
	return &Builder{
		pkg: pkg,
	}
}

func (b Builder) BuildFiles(refs []BuilderRef) {
	for _, r := range refs {
		fmt.Println(ioutil.WriteFile(r.NewPath, []byte(template.FILE_TMP), 0755))
	}
}
