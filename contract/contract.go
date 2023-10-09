package contract

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApplicationInputs struct {
	ContractFn string   `json:"contractFn"`
	Amount     uint64   `json:"amount"`
	Recipients []string `json:"recipients"`
}

type ProtocolInputs struct {
	Version     int32  `json:"version"`
	BlockHeight uint64 `json:"blockHeight"`
	BlockTime   uint64 `json:"blockTime"`
}

type AccountInfo struct {
	AccountAddress string `json:"accountAddress"`
	AccountBalance string `json:"accountBalance"`
}

type ComputeInputs struct {
	Version           int32             `json:"version"`
	AccountInfo       AccountInfo       `json:"accountInfo"`
	ProtocolInputs    ProtocolInputs    `json:"protocolInput"`
	ApplicationInputs ApplicationInputs `json:"applicationInput"`
}

type ComputeOutputs struct {
	Transactions []ComputeTransaction `json:"transactions"`
}

type ComputeTransaction struct {
	Recipient string `json:"recipient"`
	Amount    uint64 `json:"amount"`
}

func ParseInput() (error, ComputeInputs) {
	var ret ComputeInputs

	err := json.NewDecoder(os.Stdin).Decode(&ret)

	return err, ret
}

func SendOutput(out ComputeOutputs) error {
	j, err := json.Marshal(&out)

	if err == nil {
		fmt.Println(string(j))
	}

	return err
}
