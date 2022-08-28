package structs

type ExampleStructure struct {
	key string
}

func (es ExampleStructure) Get() (string, error) {
	return es.key, nil
}
