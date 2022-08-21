package interfaces

import "github.com/GuilhermeCaruso/mooncake/generator/models"

type Simple interface {
	Get() (string, error)
}

type SimpleWithRef interface {
	Get() models.File
}
