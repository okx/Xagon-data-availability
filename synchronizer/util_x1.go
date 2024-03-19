package synchronizer

import (
	"encoding/json"
	"strings"

	"github.com/0xPolygon/cdk-data-availability/etherman/smartcontracts/oldpolygonzkevm"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// UnpackTxDataFork6 unpacks the keys in a SequenceBatches event
func UnpackTxDataFork6(txData []byte) ([]common.Hash, error) {
	a, err := abi.JSON(strings.NewReader(oldpolygonzkevm.PolygonzkevmABI))
	if err != nil {
		return nil, err
	}
	method, err := a.MethodById(txData[:4])
	if err != nil {
		log.Errorf("failed to get method: %v", err)
		return nil, err
	}
	data, err := method.Inputs.Unpack(txData[4:])
	if err != nil {
		log.Errorf("failed to unpack data: %v", err)
		return nil, err
	}
	var batches []oldpolygonzkevm.PolygonZkEVMBatchData
	bytes, err := json.Marshal(data[0])
	if err != nil {
		log.Errorf("failed to marshal data: %v", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &batches)
	if err != nil {
		log.Errorf("failed to unmarshal data: %v", err)
		return nil, err
	}

	var keys []common.Hash
	for _, batch := range batches {
		keys = append(keys, batch.TransactionsHash)
	}
	return keys, nil
}
