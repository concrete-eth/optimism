package main

import (
	"github.com/ethereum-optimism/optimism/op-program/client"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
	"github.com/ethereum/go-ethereum/concrete"
)

func main() {
	// Default to a machine parsable but relatively human friendly log format.
	// Don't do anything fancy to detect if color output is supported.
	logger := oplog.NewLogger(oplog.CLIConfig{
		Level:  "info",
		Format: "logfmt",
		Color:  false,
	})
	concreteRegistry := concrete.NewRegistry()
	client.Main(logger, concreteRegistry)
}
