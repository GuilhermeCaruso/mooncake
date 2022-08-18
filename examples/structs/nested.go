package structs

type Nested struct {
}

func (n *Nested) InternalMethod() string {
	return "oi"
}

func (n *Nested) NewMethod(key string) (string, int) {
	return "oi", 20
}
