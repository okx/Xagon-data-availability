package synchronizer

import (
	"context"
	"math/big"
	"time"

	"github.com/0xPolygon/cdk-data-availability/config"
	"github.com/0xPolygon/cdk-data-availability/db"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/0xPolygon/cdk-data-availability/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	initBlockTimeout = 15 * time.Second
	minCodeLen       = 2
)

// InitStartBlock initializes the L1 sync task by finding the inception block for the CDKValidium contract
func InitStartBlock(db db.DB, ethClientFactory types.EthClientFactory, l1 config.L1Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), initBlockTimeout)
	defer cancel()

	current, err := getStartBlock(db)
	if err != nil {
		return err
	}
	if current > 0 {
		// no need to resolve start block, it's already been set
		return nil
	}
	log.Info("starting search for start block of contract ", l1.PolygonValidiumAddress)

	ethClient, err := ethClientFactory.CreateEthClient(ctx, l1.RpcURL)
	if err != nil {
		log.Errorf("failed to create eth client: %v", err)
		return err
	}
	log.Infof("find code at start to ")
	startBlock, err := findContractDeploymentBlock(ctx, ethClient, common.HexToAddress(l1.PolygonValidiumAddress))
	if err != nil {
		log.Errorf("failed to find contract deployment block: %v", err)
		return err
	}
	log.Infof("get block %d", startBlock)

	return setStartBlock(db, startBlock.Uint64())
}

func findContractDeploymentBlock(ctx context.Context, eth types.EthClient, contract common.Address) (*big.Int, error) {
	latestBlock, err := eth.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Infof("latest block %d", latestBlock.Number().Uint64())
	firstBlock := findCode(ctx, eth, contract, 0, latestBlock.Number().Int64())
	log.Infof("first block %d", firstBlock)
	return big.NewInt(firstBlock), nil
}

// findCode is an O(log(n)) search for the inception block of a contract at the given address
func findCode(ctx context.Context, eth types.EthClient, address common.Address, startBlock, endBlock int64) int64 {
	if startBlock == endBlock {
		return startBlock
	}
	midBlock := (startBlock + endBlock) / 2 //nolint:gomnd
	if codeLen := codeLen(ctx, eth, address, midBlock); codeLen > minCodeLen {
		time.Sleep(200 * time.Millisecond)
		return findCode(ctx, eth, address, startBlock, midBlock)
	} else {
		time.Sleep(200 * time.Millisecond)
		return findCode(ctx, eth, address, midBlock+1, endBlock)
	}
}

func codeLen(ctx context.Context, eth types.EthClient, address common.Address, blockNumber int64) int64 {
	data, err := eth.CodeAt(ctx, address, big.NewInt(blockNumber))
	if err != nil {
		return 0
	}
	return int64(len(data))
}
