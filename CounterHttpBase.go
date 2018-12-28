package counter

import (
    "context"
    "github.com/go-ee/utils/eh"
    "github.com/go-ee/utils/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "net/http"
)
type CounterHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *CounterQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewCounterHttpQueryHandler(queryRepository *CounterQueryRepository) (ret *CounterHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &CounterHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *CounterHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllCounter", w, r)
}

func (o *CounterHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByCounterId", w, r)
}

func (o *CounterHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllCounter", w, r)
}

func (o *CounterHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByCounterId", w, r)
}

func (o *CounterHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllCounter", w, r)
}

func (o *CounterHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByCounterId", w, r)
}


type CounterHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewCounterHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *CounterHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &CounterHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *CounterHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateCounter{Id: id}, w, r)
}

func (o *CounterHttpCommandHandler) Increment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&IncrementCounter{Id: id}, w, r)
}

func (o *CounterHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateCounter{Id: id}, w, r)
}

func (o *CounterHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteCounter{Id: id}, w, r)
}


type CounterRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *CounterHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *CounterHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewCounterRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CounterRouter) {
    pathPrefix = pathPrefix + "/" + "counters"
    entityFactory := func() eventhorizon.Entity { return NewCounter() }
    repo := readRepos(string(CounterAggregateType), entityFactory)
    queryRepository := NewCounterQueryRepository(repo, context)
    queryHandler := NewCounterHttpQueryHandler(queryRepository)
    commandHandler := NewCounterHttpCommandHandler(context, commandBus)
    ret = &CounterRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *CounterRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountCounterById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountCounterAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistCounterById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistCounterAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindCounterById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindCounterAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateCounter").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "increment").
        Name("IncrementCounter").HandlerFunc(o.CommandHandler.Increment)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateCounter").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteCounter").HandlerFunc(o.CommandHandler.Delete)
    return
}


type CounterRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    CounterRouter *CounterRouter `json:"counterRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewCounterRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CounterRouter) {
    pathPrefix = pathPrefix + "/" + "counter"
    counterRouter := NewCounterRouter(pathPrefix, context, commandBus, readRepos)
    ret = &CounterRouter{
        PathPrefix: pathPrefix,
        CounterRouter: counterRouter,
    }
    return
}

func (o *CounterRouter) Setup(router *mux.Router) (err error) {
    if err = o.CounterRouter.Setup(router); err != nil {
        return
    }
    return
}









