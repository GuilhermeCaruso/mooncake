package structs

type GenericExample[T, Z any] struct {
	value Z
}

func (ge GenericExample[T, Z]) Other(param T) (T, Z) {
	return param, ge.value
}
