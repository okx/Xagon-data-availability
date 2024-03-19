package synchronizer

import (
	"encoding/json"
	"strings"

	"github.com/0xPolygon/cdk-data-availability/etherman/smartcontracts/polygonvalidium"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// UnpackTxData unpacks the keys in a SequenceBatches event
func UnpackTxData(txData []byte) ([]common.Hash, error) {
	a, err := abi.JSON(strings.NewReader(polygonvalidium.PolygonvalidiumMetaData.ABI))
	if err != nil {
		log.Errorf("failed to parse abi: %v", err)
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

	var batches []polygonvalidium.PolygonValidiumEtrogValidiumBatchData
	bytes, err := json.Marshal(data[0])
	if err != nil {
		log.Errorf("failed to marshal data: %v", err)
		return nil, err
	}

	if err = json.Unmarshal(bytes, &batches); err != nil {
		log.Errorf("failed to unmarshal data: %v", err)
		return nil, err
	}

	keys := make([]common.Hash, len(batches))
	for i, batch := range batches {
		keys[i] = batch.TransactionsHash
	}
	return keys, nil
}
