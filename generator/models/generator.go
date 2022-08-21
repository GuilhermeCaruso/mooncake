package models

import "log"

type Generator struct {
	Files []File
}

func (g *Generator) RegisterFile(f File) {
	g.Files = append(g.Files, f)
}

func (g *Generator) PrepareNested() {
	mim := make(map[string][]Method)
	for _, f := range g.Files {
		for _, i := range f.Implementations {
			mim[i.Name] = i.Methods
		}
	}

	for idxf, f := range g.Files {
		for idxi, i := range f.Implementations {
			for _, r := range i.References {
				imp, has := mim[r]
				if !has {
					log.Fatalf("Unknown reference: %s", r)
				}
				g.Files[idxf].Implementations[idxi].Methods = append(
					g.Files[idxf].Implementations[idxi].Methods,
					imp...)
			}
		}
	}
}
