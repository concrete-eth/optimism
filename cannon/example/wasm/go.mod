module wasm

go 1.20

require (
	github.com/ethereum-optimism/optimism v0.0.0
	github.com/tetratelabs/wazero v1.5.0
)

require (
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
)

replace github.com/ethereum-optimism/optimism v0.0.0 => ../../..
