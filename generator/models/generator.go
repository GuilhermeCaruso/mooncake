package models

type Generator struct {
	Files []File
}

func (g *Generator) RegisterFile(f File) {
	g.Files = append(g.Files, f)
}
