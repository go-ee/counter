package counter

import (
    "github.com/looplab/eventhorizon"
)
type CounterInitialHandler struct {
}

func NewCounterInitialHandler() (ret *CounterInitialHandler) {
    ret = &CounterInitialHandler{}
    return
}

func (o *CounterInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *CounterInitialHandler) SetupEventHandler() (err error) {
    return
}


type CounterInitialExecutor struct {
}

func NewCounterInitialExecutor() (ret *CounterInitialExecutor) {
    ret = &CounterInitialExecutor{}
    return
}


type CounterHandlers struct {
    Initial *CounterInitialHandler `json:"initial" eh:"optional"`
}

func NewCounterHandlers() (ret *CounterHandlers) {
    initial := NewCounterInitialHandler()
    ret = &CounterHandlers{
        Initial: initial,
    }
    return
}


type CounterExecutors struct {
    Initial *CounterInitialExecutor `json:"initial" eh:"optional"`
}

func NewCounterExecutors() (ret *CounterExecutors) {
    initial := NewCounterInitialExecutor()
    ret = &CounterExecutors{
        Initial: initial,
    }
    return
}









