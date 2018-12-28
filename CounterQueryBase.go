package counter

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type CounterQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewCounterQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *CounterQueryRepository) {
    ret = &CounterQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *CounterQueryRepository) FindAll() (ret []*Counter, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Counter, len(result))
		for i, e := range result {
            ret[i] = e.(*Counter)
		}
    }
    return
}

func (o *CounterQueryRepository) FindById(id eventhorizon.UUID) (ret *Counter, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Counter)
    }
    return
}

func (o *CounterQueryRepository) CountAll() (ret int, err error) {
    var result []*Counter
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *CounterQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Counter
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *CounterQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *CounterQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









