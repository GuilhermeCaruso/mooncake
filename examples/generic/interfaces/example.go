package interfaces

type GenericInterface[T, Z any] interface {
	Other(T) (T, Z)
}

type GenericNestedInterface[T, Z any] interface {
	GenericInterface[T, Z]
}
