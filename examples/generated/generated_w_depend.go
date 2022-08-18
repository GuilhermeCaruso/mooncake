package generated_interface

import (
	"reflect"

	"github.com/GuilhermeCaruso/mooncake"
)

type NewMockMyNested struct {
	agent              *mooncake.MooncakeAgent
	internalController *InternalNewMockMyNested
}

func NewNewMockMyNested(agent *mooncake.MooncakeAgent) *NewMockMyNested {
	internal := new(NewMockMyNested)
	internal.agent = agent
	internal.internalController = &InternalNewMockMyNested{
		mock: internal,
	}
	return internal
}

func (mmy *NewMockMyNested) InternalMethod() (string, NewMockMyNested) {
	method := "InternalMethod"

	result := mmy.agent.GetCall(method)

	return0, _ := result[0].Value.(string)
	return1, _ := result[1].Value.(NewMockMyNested)

	return return0, return1
}

func (mmy *NewMockMyNested) NewMethod(arg0 string) (string, int) {
	method := "NewMethod"

	result := mmy.agent.GetCall(method)

	return0, _ := result[0].Value.(string)
	return1, _ := result[1].Value.(int)

	return return0, return1
}

func (immi *NewMockMyNested) Prepare() *InternalNewMockMyNested {
	return immi.internalController
}

type InternalNewMockMyNested struct {
	mock *NewMockMyNested
}

func (mmy *InternalNewMockMyNested) InternalMethod() *mooncake.AgentController {
	method := "InternalMethod"
	methodType := reflect.TypeOf((*NewMockMyNested)(nil).InternalMethod)

	return mmy.mock.agent.SetCall(method, methodType)
}

func (mmy *InternalNewMockMyNested) NewMethod(arg0 interface{}) *mooncake.AgentController {
	method := "NewMethod"
	methodType := reflect.TypeOf((*NewMockMyNested)(nil).NewMethod)

	return mmy.mock.agent.SetCall(method, methodType)
}
