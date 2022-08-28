package mooncake

import (
	"fmt"
	"log"
	"reflect"
)

var (
	ErrInvalidNumberOfReturns = func(key string, expected, got int) error {
		return fmt.Errorf("invalid number of returns for %s. expected=%v got=%v", key, expected, got)
	}

	ErrInvalidTypeOfReturn = func(key string, expected, got interface{}) error {
		return fmt.Errorf("invalid type of return for %s. expected=%v got=%v", key, expected, got)
	}
)

type AgentController struct {
	key           string
	returnValues  []ReturnDetail
	lifeTime      MooncakeLifetime
	lifeTimeCount int
}

type ReturnDetail struct {
	Value interface{}
	DType reflect.Type
}

func NewAgentController(key string, imp reflect.Type) AgentController {
	newAgentController := new(AgentController)
	newAgentController.key = key
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
		log.Fatalln(ErrInvalidNumberOfReturns(ag.key,
			len(ag.returnValues), len(args)).Error())
	}
	for idx, arg := range args {
		if reflect.TypeOf(arg) != ag.returnValues[idx].DType {
			log.Fatalln(ErrInvalidTypeOfReturn(ag.key,
				ag.returnValues[idx].DType, reflect.TypeOf(arg)).Error())
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
