package template

// import "github.com/GuilhermeCaruso/mooncake/generator/models"

// type Template struct {
// 	f models.File
// }

// func NewTemplate(f models.File) *Template {
// 	return &Template{
// 		f: f,
// 	}
// }

const FILE_TMP = `// ############################
// Generated by Mooncake
// Date: 2022-01-03
// Source: xpto
// ############################
package mocks
`

const METHOD_SIGN = `func (%s *%s) %s %s (%s) %s {
	%s
}`
