package main

import (
	"context"
	_ "embed"
	"encoding/binary"
	"fmt"
	"os"

	preimage "github.com/ethereum-optimism/optimism/op-preimage"
	"github.com/tetratelabs/wazero"
	wz_api "github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type rawHint string

func (rh rawHint) Hint() string {
	return string(rh)
}

func main() {
	_, _ = os.Stderr.Write([]byte("started!"))

	po := preimage.NewOracleClient(preimage.ClientPreimageChannel())
	hinter := preimage.NewHintWriter(preimage.ClientHinterChannel())

	codeHash := *(*[32]byte)(po.Get(preimage.LocalIndexKey(0)))
	diffHash := *(*[32]byte)(po.Get(preimage.LocalIndexKey(1)))
	claimData := *(*[8]byte)(po.Get(preimage.LocalIndexKey(2)))

	// Hints are used to indicate which things the program will access,
	// so the server can be prepared to serve the corresponding pre-images.
	hinter.Hint(rawHint(fmt.Sprintf("fetch-state %x", codeHash)))
	code := po.Get(preimage.Keccak256Key(codeHash))

	// Multiple pre-images may be fetched based on a hint.
	// E.g. when we need all values of a merkle-tree.
	hinter.Hint(rawHint(fmt.Sprintf("fetch-diff %x", diffHash)))
	diff := po.Get(preimage.Keccak256Key(diffHash))
	diffPartA := po.Get(preimage.Keccak256Key(*(*[32]byte)(diff[:32])))
	diffPartB := po.Get(preimage.Keccak256Key(*(*[32]byte)(diff[32:])))

	a := binary.BigEndian.Uint64(diffPartA)
	b := binary.BigEndian.Uint64(diffPartB)

	mod := newWazeroModule(code, wazero.NewRuntimeConfigInterpreter())
	add := mod.ExportedFunction("add")

	fmt.Printf("computing %d + %d in wasm\n", a, b)

	_ret, err := add.Call(context.Background(), a, b)
	if err != nil {
		fmt.Printf("failed to call add: %v\n", err)
		os.Exit(1)
	}
	sOut := _ret[0]

	sClaim := binary.BigEndian.Uint64(claimData[:])
	if sOut != sClaim {
		fmt.Printf("claim %d is bad! Correct result is %d\n", sOut, sClaim)
		os.Exit(1)
	} else {
		fmt.Printf("claim %d is good!\n", sOut)
		os.Exit(0)
	}
}

func newWazeroModule(code []byte, config wazero.RuntimeConfig) wz_api.Module {
	ctx := context.Background()
	r := wazero.NewRuntimeWithConfig(ctx, config)
	_, err := r.NewHostModuleBuilder("env").Instantiate(ctx)
	if err != nil {
		fmt.Printf("failed to instantiate host module: %v\n", err)
		os.Exit(1)
	}
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	mod, err := r.Instantiate(ctx, code)
	if err != nil {
		fmt.Printf("failed to instantiate module: %v\n", err)
		os.Exit(1)
	}
	return mod
}
