package counter

import (
    "encoding/json"
    "fmt"
    "github.com/go-ee/utils/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
)
const (
     CounterCreatedEvent eventhorizon.EventType = "CounterCreated"
     CounterDeletedEvent eventhorizon.EventType = "CounterDeleted"
     CounterIncrementedEvent eventhorizon.EventType = "CounterIncremented"
     CounterUpdatedEvent eventhorizon.EventType = "CounterUpdated"
)




type CounterCreated struct {
    Id string `json:"id" eh:"optional"`
    Count int `json:"count" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CounterDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CounterIncremented struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CounterUpdated struct {
    Id string `json:"id" eh:"optional"`
    Count int `json:"count" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}




type CounterEventType struct {
	name  string
	ordinal int
}

func (o *CounterEventType) Name() string {
    return o.name
}

func (o *CounterEventType) Ordinal() int {
    return o.ordinal
}

func (o CounterEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CounterEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CounterEventTypes().ParseCounterEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CounterEventType %q", lit.Name)
        }
	}
	return
}

func (o CounterEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CounterEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CounterEventTypes().ParseCounterEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CounterEventType %q", lit)
        }
    }
    return
}

func (o *CounterEventType) IsCounterCreated() bool {
    return o == _counterEventTypes.CounterCreated()
}

func (o *CounterEventType) IsCounterDeleted() bool {
    return o == _counterEventTypes.CounterDeleted()
}

func (o *CounterEventType) IsCounterIncremented() bool {
    return o == _counterEventTypes.CounterIncremented()
}

func (o *CounterEventType) IsCounterUpdated() bool {
    return o == _counterEventTypes.CounterUpdated()
}

type counterEventTypes struct {
	values []*CounterEventType
    literals []enum.Literal
}

var _counterEventTypes = &counterEventTypes{values: []*CounterEventType{
    {name: "CounterCreated", ordinal: 0},
    {name: "CounterDeleted", ordinal: 1},
    {name: "CounterIncremented", ordinal: 2},
    {name: "CounterUpdated", ordinal: 3}},
}

func CounterEventTypes() *counterEventTypes {
	return _counterEventTypes
}

func (o *counterEventTypes) Values() []*CounterEventType {
	return o.values
}

func (o *counterEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *counterEventTypes) CounterCreated() *CounterEventType {
    return _counterEventTypes.values[0]
}

func (o *counterEventTypes) CounterDeleted() *CounterEventType {
    return _counterEventTypes.values[1]
}

func (o *counterEventTypes) CounterIncremented() *CounterEventType {
    return _counterEventTypes.values[2]
}

func (o *counterEventTypes) CounterUpdated() *CounterEventType {
    return _counterEventTypes.values[3]
}

func (o *counterEventTypes) ParseCounterEventType(name string) (ret *CounterEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CounterEventType), ok
	}
	return
}



