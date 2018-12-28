package counter

import (
    "encoding/json"
    "fmt"
    "github.com/go-ee/utils/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
)
const (
     CreateCounterCommand eventhorizon.CommandType = "CreateCounter"
     DeleteCounterCommand eventhorizon.CommandType = "DeleteCounter"
     IncrementCounterCommand eventhorizon.CommandType = "IncrementCounter"
     UpdateCounterCommand eventhorizon.CommandType = "UpdateCounter"
)




        
type CreateCounter struct {
    Id string `json:"id" eh:"optional"`
    Count int `json:"count" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateCounter) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateCounter) AggregateType() eventhorizon.AggregateType  { return CounterAggregateType }
func (o *CreateCounter) CommandType() eventhorizon.CommandType      { return CreateCounterCommand }



        
type DeleteCounter struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteCounter) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteCounter) AggregateType() eventhorizon.AggregateType  { return CounterAggregateType }
func (o *DeleteCounter) CommandType() eventhorizon.CommandType      { return DeleteCounterCommand }



        
type IncrementCounter struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *IncrementCounter) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *IncrementCounter) AggregateType() eventhorizon.AggregateType  { return CounterAggregateType }
func (o *IncrementCounter) CommandType() eventhorizon.CommandType      { return IncrementCounterCommand }



        
type UpdateCounter struct {
    Id string `json:"id" eh:"optional"`
    Count int `json:"count" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateCounter) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateCounter) AggregateType() eventhorizon.AggregateType  { return CounterAggregateType }
func (o *UpdateCounter) CommandType() eventhorizon.CommandType      { return UpdateCounterCommand }





type CounterCommandType struct {
	name  string
	ordinal int
}

func (o *CounterCommandType) Name() string {
    return o.name
}

func (o *CounterCommandType) Ordinal() int {
    return o.ordinal
}

func (o CounterCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CounterCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CounterCommandTypes().ParseCounterCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CounterCommandType %q", lit.Name)
        }
	}
	return
}

func (o CounterCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CounterCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CounterCommandTypes().ParseCounterCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CounterCommandType %q", lit)
        }
    }
    return
}

func (o *CounterCommandType) IsCreateCounter() bool {
    return o == _counterCommandTypes.CreateCounter()
}

func (o *CounterCommandType) IsDeleteCounter() bool {
    return o == _counterCommandTypes.DeleteCounter()
}

func (o *CounterCommandType) IsIncrementCounter() bool {
    return o == _counterCommandTypes.IncrementCounter()
}

func (o *CounterCommandType) IsUpdateCounter() bool {
    return o == _counterCommandTypes.UpdateCounter()
}

type counterCommandTypes struct {
	values []*CounterCommandType
    literals []enum.Literal
}

var _counterCommandTypes = &counterCommandTypes{values: []*CounterCommandType{
    {name: "CreateCounter", ordinal: 0},
    {name: "DeleteCounter", ordinal: 1},
    {name: "IncrementCounter", ordinal: 2},
    {name: "UpdateCounter", ordinal: 3}},
}

func CounterCommandTypes() *counterCommandTypes {
	return _counterCommandTypes
}

func (o *counterCommandTypes) Values() []*CounterCommandType {
	return o.values
}

func (o *counterCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *counterCommandTypes) CreateCounter() *CounterCommandType {
    return _counterCommandTypes.values[0]
}

func (o *counterCommandTypes) DeleteCounter() *CounterCommandType {
    return _counterCommandTypes.values[1]
}

func (o *counterCommandTypes) IncrementCounter() *CounterCommandType {
    return _counterCommandTypes.values[2]
}

func (o *counterCommandTypes) UpdateCounter() *CounterCommandType {
    return _counterCommandTypes.values[3]
}

func (o *counterCommandTypes) ParseCounterCommandType(name string) (ret *CounterCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CounterCommandType), ok
	}
	return
}



