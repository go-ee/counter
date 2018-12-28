package counter

import (
    "errors"
    "fmt"
    "github.com/go-ee/utils/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type CounterCommandHandler struct {
    CreateHandler func (*CreateCounter, *Counter, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteCounter, *Counter, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    IncrementHandler func (*IncrementCounter, *Counter, eh.AggregateStoreEvent) (err error)  `json:"incrementHandler" eh:"optional"`
    UpdateHandler func (*UpdateCounter, *Counter, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *CounterCommandHandler) AddCreatePreparer(preparer func (*CreateCounter, *Counter) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CounterCommandHandler) AddDeletePreparer(preparer func (*DeleteCounter, *Counter) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CounterCommandHandler) AddIncrementPreparer(preparer func (*IncrementCounter, *Counter) (err error) ) {
    prevHandler := o.IncrementHandler
	o.IncrementHandler = func(command *IncrementCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CounterCommandHandler) AddUpdatePreparer(preparer func (*UpdateCounter, *Counter) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *CounterCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateCounterCommand:
        err = o.CreateHandler(cmd.(*CreateCounter), entity.(*Counter), store)
    case DeleteCounterCommand:
        err = o.DeleteHandler(cmd.(*DeleteCounter), entity.(*Counter), store)
    case IncrementCounterCommand:
        err = o.IncrementHandler(cmd.(*IncrementCounter), entity.(*Counter), store)
    case UpdateCounterCommand:
        err = o.UpdateHandler(cmd.(*UpdateCounter), entity.(*Counter), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *CounterCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, CounterAggregateType); err == nil {
            store.StoreEvent(CounterCreatedEvent, &CounterCreated{
                Id: command.Id,
                Count: command.Count,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CounterAggregateType); err == nil {
            store.StoreEvent(CounterDeletedEvent, &CounterDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.IncrementHandler = func(command *IncrementCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CounterAggregateType); err == nil {
            store.StoreEvent(CounterIncrementedEvent, &CounterIncremented{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateCounter, entity *Counter, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, CounterAggregateType); err == nil {
            store.StoreEvent(CounterUpdatedEvent, &CounterUpdated{
                Id: command.Id,
                Count: command.Count,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type CounterEventHandler struct {
    CreatedHandler func (*CounterCreated, *Counter) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*CounterDeleted, *Counter) (err error)  `json:"deletedHandler" eh:"optional"`
    IncrementedHandler func (*CounterIncremented, *Counter) (err error)  `json:"incrementedHandler" eh:"optional"`
    UpdatedHandler func (*CounterUpdated, *Counter) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *CounterEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case CounterCreatedEvent:
        err = o.CreatedHandler(event.Data().(*CounterCreated), entity.(*Counter))
    case CounterDeletedEvent:
        err = o.DeletedHandler(event.Data().(*CounterDeleted), entity.(*Counter))
    case CounterIncrementedEvent:
        err = o.IncrementedHandler(event.Data().(*CounterIncremented), entity.(*Counter))
    case CounterUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*CounterUpdated), entity.(*Counter))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *CounterEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(CounterCreatedEvent, func() eventhorizon.EventData {
		return &CounterCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *CounterCreated, entity *Counter) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, CounterAggregateType); err == nil {
            entity.Id = event.Id
            entity.Count = event.Count
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CounterDeletedEvent, func() eventhorizon.EventData {
		return &CounterDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *CounterDeleted, entity *Counter) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CounterAggregateType); err == nil {
            *entity = *NewCounter()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CounterIncrementedEvent, func() eventhorizon.EventData {
		return &CounterIncremented{}
	})

    //default handler implementation
    o.IncrementedHandler = func(event *CounterIncremented, entity *Counter) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CounterAggregateType); err == nil {
            entity.Count = ee.lang.Attribute@327b636c
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(CounterUpdatedEvent, func() eventhorizon.EventData {
		return &CounterUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *CounterUpdated, entity *Counter) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, CounterAggregateType); err == nil {
            entity.Id = event.Id
            entity.Count = event.Count
        }
        return
    }
    return
}


const CounterAggregateType eventhorizon.AggregateType = "Counter"

type CounterAggregateInitializer struct {
    *eh.AggregateInitializer
    *CounterCommandHandler
    *CounterEventHandler
    ProjectorHandler *CounterEventHandler `json:"projectorHandler" eh:"optional"`
}



func NewCounterAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CounterAggregateInitializer) {
    
    commandHandler := &CounterCommandHandler{}
    eventHandler := &CounterEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewCounter() }
    ret = &CounterAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(CounterAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(CounterAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        CounterCommandTypes().Literals(), CounterEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, commandBus, readRepos), CounterCommandHandler: commandHandler, CounterEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type CounterEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    CounterAggregateInitializer *CounterAggregateInitializer `json:"counterAggregateInitializer" eh:"optional"`
}

func NewCounterEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CounterEventhorizonInitializer) {
    counterAggregateInitializer := NewCounterAggregateInitializer(eventStore, eventBus, commandBus, readRepos)
    ret = &CounterEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        commandBus: commandBus,
        CounterAggregateInitializer: counterAggregateInitializer,
    }
    return
}

func (o *CounterEventhorizonInitializer) Setup() (err error) {
    
    if err = o.CounterAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









