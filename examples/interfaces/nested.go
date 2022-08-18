package interfaces

import (
	"time"

	generated_interface "github.com/GuilhermeCaruso/mooncake/examples/generated"
)

type Other interface {
	Get() string
}

type InternalInterface interface {
	InternalMethod() (string, generated_interface.NewMockMyNested)
	NewMethod(string) (string, int, time.Ticker)
}

type InterfaceGeneric[T, Z any] interface {
	Test(T) T
}

type InterfaceMultiple[T any, Z any] interface {
	Test(T) (T, Z)
}

type Test struct {
	A string
}

type RootInterface interface {
	InternalInterface
}

type RootInterfaceWithTwo interface {
	InternalInterface
	Other
}
