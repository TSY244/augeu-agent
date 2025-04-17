package agent

import "augeu-agent/internal/pkg/weibu"

func (a *Agent) GetWeiBuConf() *weibu.Config {
	return &weibu.Config{
		Jwt:      a.Jwt,
		Header:   a.Header,
		Conf:     a.Conf,
		ClientId: a.ClientId,
	}
}
