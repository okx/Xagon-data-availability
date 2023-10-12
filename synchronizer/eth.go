package synchronizer

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/0xPolygon/cdk-data-availability/config"
	"github.com/0xPolygonHermez/zkevm-node/etherman"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/datacommittee"
	"github.com/0xPolygonHermez/zkevm-node/etherman/smartcontracts/polygonzkevm"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func newRPCEtherman(cfg config.L1Config) (*etherman.Client, error) {
	return newEtherman(cfg, cfg.RpcURL)
}

func newWSEtherman(cfg config.L1Config) (*etherman.Client, error) {
	return newEtherman(cfg, cfg.WsURL)
}

// newEtherman constructs an etherman client that only needs the free contract calls
func newEtherman(cfg config.L1Config, url string) (*etherman.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout.Duration)
	defer cancel()
	ethClient, err := ethclient.DialContext(ctx, url)
	if err != nil {
		log.Errorf("error connecting to %s: %+v", url, err)
		return nil, err
	}
	zkevm, err := polygonzkevm.NewPolygonzkevm(common.HexToAddress(cfg.CDKValidiumAddress), ethClient)
	if err != nil {
		return nil, err
	}
	dataCommittee, err :=
		datacommittee.NewDatacommittee(common.HexToAddress(cfg.DataCommitteeAddress), ethClient)
	if err != nil {
		return nil, err
	}
	return &etherman.Client{
		EthClient:     ethClient,
		ZkEVM:         zkevm,
		DataCommittee: dataCommittee,
	}, nil
}

// ParseEvent unpacks the keys in a SequenceBatches event
func ParseEvent(event *polygonzkevm.PolygonzkevmSequenceBatches, txData []byte) (uint64, []common.Hash, error) {
	a, err := abi.JSON(strings.NewReader(polygonzkevm.PolygonzkevmMetaData.ABI))
	if err != nil {
		return 0, nil, err
	}
	method, err := a.MethodById(txData[:4])
	if err != nil {
		return 0, nil, err
	}
	data, err := method.Inputs.Unpack(txData[4:])
	if err != nil {
		return 0, nil, err
	}
	var batches []polygonzkevm.PolygonZkEVMBatchData
	bytes, err := json.Marshal(data[0])
	if err != nil {
		return 0, nil, err
	}
	err = json.Unmarshal(bytes, &batches)
	if err != nil {
		return 0, nil, err
	}

	var keys []common.Hash
	for _, batch := range batches {
		if len(batch.Transactions) > 0 {
			hash := crypto.Keccak256Hash(batch.Transactions)
			keys = append(keys, hash)
			log.Infof("parse no dac, batch num:%v:batch timestamp:%v, calc hash:%s", event.NumBatch, batch.Timestamp, hash.String())
		} else {
			keys = append(keys, batch.TransactionsHash)
			log.Infof("parse use dac, batch num:%v, batch timestamp:%v, hash:%s", event.NumBatch, batch.Timestamp, hex.EncodeToString(batch.TransactionsHash[:]))
		}
	}
	return event.Raw.BlockNumber, keys, nil
}
