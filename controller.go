package mooncake

import (
	"fmt"
	"os"
	"reflect"
)

type AgentController struct {
	returnValues  []ReturnDetail
	lifeTime      MooncakeLifetime
	lifeTimeCount int
}

type ReturnDetail struct {
	Value interface{}
	DType reflect.Type
}

func NewAgentController(imp reflect.Type) AgentController {
	newAgentController := new(AgentController)

	newAgentController.lifeTime = LT_ONE_CALL
	newAgentController.lifeTimeCount = 1

	for x := 0; x < imp.NumOut(); x++ {
		newAgentController.returnValues = append(newAgentController.returnValues, ReturnDetail{
			Value: nil,
			DType: imp.Out(x),
		})
	}
	return *newAgentController
}

func (ag *AgentController) SetReturn(args ...interface{}) *AgentController {
	if len(ag.returnValues) != len(args) {
		fmt.Println("invalid number of returns")
		os.Exit(1)
	}
	for idx, arg := range args {
		if reflect.TypeOf(arg) != ag.returnValues[idx].DType {
			fmt.Println("invalid type of return")
			os.Exit(1)
		}
		ag.returnValues[idx].Value = arg
	}
	return ag
}

func (ag *AgentController) AnyTime() *AgentController {
	ag.lifeTime = LT_ANY_TIME
	return ag
}

func (ag *AgentController) Repeat(count int) *AgentController {
	ag.lifeTime = LT_REPEAT
	ag.lifeTimeCount = count
	return ag
}
