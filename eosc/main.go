package main

import (
	// Load all contracts here, so we can always read and decode
	// transactions with those contracts.
	_ "github.com/vadim-di/eos-go/msig"
	_ "github.com/vadim-di/eos-go/system"
	_ "github.com/vadim-di/eos-go/token"

	"github.com/vadim-di/eosc/eosc/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
