package engine

import (
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

type Engine struct {
	DataContext *context.DataContext
	eng         *engine.Gengine
	ruleBuilder *builder.RuleBuilder
}

func NewEngine() *Engine {
	ctx := context.NewDataContext()
	return &Engine{
		DataContext: ctx,
		eng:         engine.NewGengine(),
		ruleBuilder: builder.NewRuleBuilder(ctx),
	}
}

func (e *Engine) InitObject(name string, obj interface{}) {
	e.DataContext.Add(name, obj)
}

func (e *Engine) LoadRule(rule string) error {
	err := e.ruleBuilder.BuildRuleFromString(rule)
	if err != nil {
		return err
	}
	return e.eng.Execute(e.ruleBuilder, true)
}
