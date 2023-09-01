package contract

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContractInput struct {
	Version       int32  `json:"version"`
	TxId          string `json:"txId"`
	LastBlockTime int64  `json:"lastBlockTime"`
}

type ContractResult struct {
	Result int32 `json:"result"`
	Value  int32 `json:"value"`
}

type ContractOutput struct {
	Stdin  ContractInput     `json:"stdin"`
	Result ContractResult    `json:"result"`
	Args   []string          `json:"args"`
	Env    map[string]string `json:"env"`
}

func ParseInput() (error, ContractOutput) {
	var ret ContractOutput

	err := json.NewDecoder(os.Stdin).Decode(&ret.Stdin)

	return err, ret
}

func SendOutput(out ContractOutput) error {
	j, err := json.Marshal(&out)

	if err == nil {
		fmt.Println(string(j))
	}

	return err
}
