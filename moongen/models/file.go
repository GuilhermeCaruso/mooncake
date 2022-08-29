package models

import (
	"bytes"
	"fmt"
)

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

func (i Implementation) ParamsToString() (string, string) {
	var pWithType bytes.Buffer
	var pWithoutType bytes.Buffer

	for idx, p := range i.Params {
		if idx == 0 {
			pWithType.WriteString("[")
			pWithoutType.WriteString("[")
		}

		pWithType.WriteString(p.stringWithType())
		pWithoutType.WriteString(p.Name)

		if idx < len(i.Params)-1 {
			pWithType.WriteString(",")
			pWithoutType.WriteString(", ")
		}

		if idx == len(i.Params)-1 {
			pWithType.WriteString("]")
			pWithoutType.WriteString("]")
		}
	}

	return pWithType.String(), pWithoutType.String()
}

type Param struct {
	Name string
	Type string
}

func (p Param) stringWithType() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type)
}

type Method struct {
	Name    string
	Params  []Arg
	Results []Arg
}

type Arg struct {
	Name string
	Type string
}

func GetArgListString(pList []Arg) string {
	var buffer bytes.Buffer

	for idx, p := range pList {
		buffer.WriteString(fmt.Sprintf("%s %s", p.Name, p.Type))
		if idx < len(pList)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

func GetResultListString(pList []Arg) string {
	var buffer bytes.Buffer

	for idx, p := range pList {
		if idx > 0 {
			buffer.WriteString("\t")
		}
		buffer.WriteString(fmt.Sprintf("%s,_ = result[%v].Value.(%s)", p.Name, idx, p.Type))
		if idx != len(pList)-1 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

func GetArgGenericListString(pList []Arg) string {
	var buffer bytes.Buffer

	for idx, p := range pList {
		buffer.WriteString(fmt.Sprintf("%s interface{}", p.Name))
		if idx < len(pList)-1 {
			buffer.WriteString(", ")
		}
	}

	return buffer.String()
}

func (p Arg) stringWithType() string {
	return fmt.Sprintf("%s %s", p.Name, p.Type)
}
