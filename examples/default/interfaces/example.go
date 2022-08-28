package interfaces

type SimpleInterface interface {
	Get() (string, error)
}

type NestedInterface interface {
	SimpleInterface
}
