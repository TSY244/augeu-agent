package engUtils

import "sync"

type Env struct {
	Data *sync.Map
}

var (
	CoreEnv = NewEnv()
)

func init() {

	CoreEnv = NewEnv()

}

func NewEnv() *Env {
	return &Env{
		Data: new(sync.Map),
	}
}

func (e *Env) SetEnv(key string, value string) {
	e.Data.Store(key, value)
}

func (e *Env) GetEnv(key string) string {
	if value, ok := e.Data.Load(key); ok {
		return value.(string)
	}
	return ""
}
