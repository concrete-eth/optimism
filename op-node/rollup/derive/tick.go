package derive

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

const (
	TickFuncSignature = "tick()"
	TickArguments     = 0
	TickLen           = 4 + 32*TickArguments
)

var (
	TickFuncBytes4       = crypto.Keccak256([]byte(TickFuncSignature))[:4]
	TickDepositerAddress = common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddead0001")
	TickAddress          = predeploys.TickAddr
)

func TickDeposit(rollupCfg *rollup.Config, sysCfg eth.SystemConfig, seqNumber uint64, block eth.BlockInfo, l2BlockTime uint64) (*types.DepositTx, error) {
	source := L1InfoDepositSource{
		L1BlockHash: block.Hash(),
		SeqNumber:   seqNumber,
	}
	out := &types.DepositTx{
		SourceHash:          source.SourceHash(),
		From:                TickDepositerAddress,
		To:                  &TickAddress,
		Mint:                nil,
		Value:               big.NewInt(0),
		Gas:                 rollupCfg.TickGasLimit,
		IsSystemTransaction: true,
		Data:                TickFuncBytes4,
	}
	if rollupCfg.IsRegolith(l2BlockTime) {
		out.IsSystemTransaction = false
	}
	return out, nil
}

func TickDepositBytes(rollupCfg *rollup.Config, sysCfg eth.SystemConfig, seqNumber uint64, l1Info eth.BlockInfo, l2BlockTime uint64) ([]byte, error) {
	dep, err := TickDeposit(rollupCfg, sysCfg, seqNumber, l1Info, l2BlockTime)
	if err != nil {
		return nil, fmt.Errorf("failed to create tick tx: %w", err)
	}
	l1Tx := types.NewTx(dep)
	opaqueL1Tx, err := l1Tx.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to encode tick tx: %w", err)
	}
	return opaqueL1Tx, nil
}
