package counter

import (
    "github.com/looplab/eventhorizon"
)
        
type Counter struct {
    Id string `json:"id" eh:"optional"`
    Count int `json:"count" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewCounter() (ret *Counter) {
    ret = &Counter{}
    return
}
func (o *Counter) EntityID() eventhorizon.UUID { return o.Id }










