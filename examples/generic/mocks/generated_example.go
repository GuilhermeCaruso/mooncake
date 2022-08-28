// ############################
// Generated by Mooncake
// Date: 2022-08-28 17:30:16
// Source: examples/interfaces/example.go
// ############################
package mocks

import (
	"reflect"

	"github.com/GuilhermeCaruso/mooncake"
)

type MockGenericInterface[T, Z any] struct {
	agent              *mooncake.MooncakeAgent
	internalController *InternalMockGenericInterface[T, Z]
}

func NewMockGenericInterface[T, Z any](agent *mooncake.MooncakeAgent) *MockGenericInterface[T, Z] {
	internal := new(MockGenericInterface[T, Z])
	internal.agent = agent
	internal.internalController = &InternalMockGenericInterface[T, Z]{
		mock: internal,
	}
	return internal
}

func (m *MockGenericInterface[T, Z]) Other(param0 T) (result0 T, result1 Z) {
	method := "Other"

	result := m.agent.GetCall(method)

	result0, _ = result[0].Value.(T)
	result1, _ = result[1].Value.(Z)

	return
}

func (immi *MockGenericInterface[T, Z]) Prepare() *InternalMockGenericInterface[T, Z] {
	return immi.internalController
}

type InternalMockGenericInterface[T, Z any] struct {
	mock *MockGenericInterface[T, Z]
}

func (im *InternalMockGenericInterface[T, Z]) Other(param0 T) *mooncake.AgentController {
	method := "Other"
	methodType := reflect.TypeOf((*MockGenericInterface[T, Z])(nil).Other)

	return im.mock.agent.SetCall(method, methodType)
}

type MockGenericNestedInterface[T, Z any] struct {
	agent              *mooncake.MooncakeAgent
	internalController *InternalMockGenericNestedInterface[T, Z]
}

func NewMockGenericNestedInterface[T, Z any](agent *mooncake.MooncakeAgent) *MockGenericNestedInterface[T, Z] {
	internal := new(MockGenericNestedInterface[T, Z])
	internal.agent = agent
	internal.internalController = &InternalMockGenericNestedInterface[T, Z]{
		mock: internal,
	}
	return internal
}

func (immi *MockGenericNestedInterface[T, Z]) Prepare() *InternalMockGenericNestedInterface[T, Z] {
	return immi.internalController
}

type InternalMockGenericNestedInterface[T, Z any] struct {
	mock *MockGenericNestedInterface[T, Z]
}