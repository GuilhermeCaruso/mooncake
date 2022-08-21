package models

type File struct {
	Imports         []Import
	Implementations []Implementation
}
type Import struct {
	Path string
	Name string
}

type Implementation struct {
	Name       string
	Params     []Param
	Methods    []Method
	References []string
}

type Param struct {
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
