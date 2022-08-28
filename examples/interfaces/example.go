package interfaces

type SimpleInterface interface {
	Get() (string, error)
}

type NestedInterface interface {
	SimpleInterface
}

type GenericInterface[T, Z any] interface {
	Test(T) (T, Z)
}

type GenericNestedInterface[T, Z any] interface {
	GenericInterface[T, Z]
}
