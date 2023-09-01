package main

import (
	"os"
	"strings"

	"github.com/versatus/versatus-go/contract"
)

func main() {
	_, myContract := contract.ParseInput()

	myContract.Result.Result = 1
	myContract.Result.Value = 3
	myContract.Args = os.Args
	myContract.Env = make(map[string]string)
	for _, item := range os.Environ() {
		fields := strings.Split(item, "=")
		myContract.Env[fields[0]] = fields[1]
		if fields[0] == "RETURN_FAIL" {
			panic("Asked to panic through TEST_RETURN_FAIL environment variable")
		}
	}

	_ = contract.SendOutput(myContract)
}
