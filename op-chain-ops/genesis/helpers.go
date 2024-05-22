package genesis

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// ImplementationSlot represents the EIP 1967 implementation storage slot
	ImplementationSlot = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
	// AdminSlot represents the EIP 1967 admin storage slot
	AdminSlot = common.HexToHash("0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103")
	// The devBalance is the amount of wei that a dev account is funded with.
	devBalance = hexutil.MustDecodeBig("0x200000000000000000000000000000000000000000000000000000000000000")
)

// DevAccounts represent the standard hardhat development accounts.
// These are funded if the deploy config has funding development
// accounts enabled.
var DevAccounts = []common.Address{
	common.HexToAddress("0xa2FC986eFdcc89A451fAe091ddDD324f68C7F6C6"), // 0
	common.HexToAddress("0x7651669f5A2253375EA511Af8badf1Eb34242843"), // 1
	common.HexToAddress("0x88cb73d010499941b66dF26F3845DC3a95659d17"), // 2
	common.HexToAddress("0xF46a50040E7BC6206abEd2b58b191562401986F2"), // 3
	common.HexToAddress("0x132cF18c9970B747691bd9bCD8295F7A01215da0"), // 4
	common.HexToAddress("0x9C29a907feF47b30eaDb30B5D88adAcb13bAa9f1"), // 5
	common.HexToAddress("0x89F577286a1759f5bA0dF3b8B1e6851BdBc75C67"), // 6
	common.HexToAddress("0x7a31dC7bBB7046b9D90Af35789eAfd288BD3242F"), // 7
	common.HexToAddress("0x05468b4889d16Eb4783090ae23505949cd87B462"), // 8
	common.HexToAddress("0xE9d13c6f0c63bA23fD5902feCA5e9AC9584320E5"), // 9
	common.HexToAddress("0xB71d115a92433a4dc36F5EAF4cb078d3759A0390"), // 10
	common.HexToAddress("0xF1b0CE750801F237793040427ECDD8366504C61b"), // 11
	common.HexToAddress("0xBE6298da164e59641496CEaf09a96e7e61460074"), // 12
	common.HexToAddress("0xB0eeE1743693dE695B94F806E98dAcd8EEb10ec5"), // 13
	common.HexToAddress("0xe0489cCd21089DF1b94BA3773c6332831c7e3C7a"), // 14
	common.HexToAddress("0x8b5AD1159137Fb5d934045E8ac65DE30ed55Db32"), // 15
	common.HexToAddress("0x3DE68e54d2897d4832d27e1f4FCc3f4e15d0B79B"), // 16
	common.HexToAddress("0xe1b8a1269E84c8d4a63Fd32976a04903495c93ED"), // 17
	common.HexToAddress("0xE513567725a74A92fb1a1F675a38BdD343605eEb"), // 18
	common.HexToAddress("0xE59014f168adA2f3ea0170434227549B8C72Fb84"), // 19
	common.HexToAddress("0x8D18E6F89B811695Ed333d248A68C532d1B261ed"), // 20
	// Test account used by geth tests
	// common.HexToAddress("0x71562b71999873DB5b286dF957af199Ec94617F7"),
	// Deployer of create2 deterministic proxy https://github.com/Arachnid/deterministic-deployment-proxy
	// common.HexToAddress("0x3fab184622dc19b6109349b94811493bf2a45362"),
}

// GetBlockFromTag will resolve a Block given an rpc block tag
func GetBlockFromTag(chain ethereum.ChainReader, tag *rpc.BlockNumberOrHash) (*types.Block, error) {
	if hash, ok := tag.Hash(); ok {
		block, err := chain.BlockByHash(context.Background(), hash)
		if err != nil {
			return nil, err
		}
		return block, nil
	} else if num, ok := tag.Number(); ok {
		blockNumber := new(big.Int).SetInt64(num.Int64())
		block, err := chain.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			return nil, err
		}
		return block, nil
	} else {
		return nil, fmt.Errorf("invalid block tag: %v", tag)
	}
}

// uint642Big creates a new *big.Int from a uint64.
func uint642Big(in uint64) *big.Int {
	return new(big.Int).SetUint64(in)
}

func newHexBig(in uint64) *hexutil.Big {
	b := new(big.Int).SetUint64(in)
	hb := hexutil.Big(*b)
	return &hb
}
