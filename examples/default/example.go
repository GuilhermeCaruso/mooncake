package example

type ExampleStructure struct {
}

func (es ExampleStructure) Get() (string, error) {
	return "fixed", nil
}
