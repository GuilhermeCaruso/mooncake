package example

type GenericExample[T, Z any] struct {
	Value Z
}

func (ge GenericExample[T, Z]) Other(param T) (T, Z) {
	return param, ge.Value
}
