package param

import (
	"fmt"
	"strings"
)

type StrSplice struct {
	Values map[string]string
}

func NewStrSplice() StrSplice {
	return StrSplice{
		Values: make(map[string]string),
	}
}

func (s *StrSplice) String() string {
	return fmt.Sprintf("%v", s.Values)
}

func (s *StrSplice) Set(value string) error {
	if !strings.Contains(value, ":") {
		return fmt.Errorf("env format error, usage: -env key1:value1")
	}
	key := strings.Split(value, ":")[0]
	value = strings.Split(value, ":")[1]
	s.Values[key] = value
	return nil
}

func (s *StrSplice) GetEnvMap() map[string]string {
	return s.Values
}
