package interfaces

import (
	"context"
	"time"
)

type Other interface {
	Get() string
}

type InternalInterface interface {
	ReturnContext(context.Context) context.Context
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
	Internal(string) string
}
