package mooncake

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"sync"
)

type MooncakeAgent struct {
	mutex sync.Mutex
	queue map[string]*AgentController
}

func NewAgent() *MooncakeAgent {
	ma := new(MooncakeAgent)
	ma.CleanQueue()
	return ma
}

func (ma *MooncakeAgent) CleanQueue() *MooncakeAgent {
	ma.queue = map[string]*AgentController{}
	return ma
}

func (ma *MooncakeAgent) SetCall(key string, typeImpl reflect.Type) *AgentController {
	ma.mutex.Lock()
	defer ma.mutex.Unlock()

	var agentTo AgentController

	if v, has := ma.queue[key]; !has {
		ma.queue[key] = new(AgentController)
		agentTo = NewAgentController(typeImpl)
	} else {
		v.lifeTimeCount++
		v.lifeTime = LT_REPEAT
		agentTo = *v
	}
	ma.queue[key] = &agentTo
	return &agentTo
}

func (ma *MooncakeAgent) GetCall(key string) []ReturnDetail {
	ma.mutex.Lock()
	defer ma.mutex.Unlock()
	if v, has := ma.queue[key]; has {
		valueToReturn := v.returnValues
		v.lifeTimeCount--
		if v.lifeTimeCount <= 0 && v.lifeTime != LT_ANY_TIME {
			delete(ma.queue, key)
		}
		return valueToReturn
	} else {
		fmt.Println(errors.New("No call registered"))
		os.Exit(1)
	}
	return []ReturnDetail{}
}
