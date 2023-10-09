package main

import (
	"github.com/versatus/versatus-go/contract"
)

func main() {
	var (
		outputs contract.ComputeOutputs
		txn     contract.ComputeTransaction
	)

	_, inputs := contract.ParseInput()

	amountEach := inputs.ApplicationInputs.Amount / uint64(len(inputs.ApplicationInputs.Recipients))

	for _, r := range inputs.ApplicationInputs.Recipients {
		txn.Recipient = r
		txn.Amount = amountEach
		outputs.Transactions = append(outputs.Transactions, txn)
	}

	_ = contract.SendOutput(outputs)
}
