package agent

import (
	"encoding/base64"
	"encoding/json"
	"os"
)

type LocalRule struct {
	Secrete string `json:"secrete"`
	Rule    string `json:"rule"`
}

func (a *Agent) GetLocalRule(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var localRule LocalRule
	err = json.Unmarshal(data, &localRule)
	if err != nil {
		return err
	}
	// base64 解码
	data, err = base64.StdEncoding.DecodeString(localRule.Rule)
	if err != nil {
		return err
	}

	a.SetRule(string(data))
	return nil
}
