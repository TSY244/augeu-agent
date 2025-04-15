package engine

import (
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

type Engine struct {
	DataContext *context.DataContext
	eng         *engine.Gengine
}

func NewEngine() *Engine {
	ctx := context.NewDataContext()
	return &Engine{
		DataContext: ctx,
		eng:         engine.NewGengine(),
	}
}
