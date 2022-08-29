package example

type SimpleInterface interface {
	Get() (string, error)
}

type NestedInterface interface {
	SimpleInterface
}
