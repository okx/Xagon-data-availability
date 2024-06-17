package rpcsyncer

import (
	"encoding/json"
	"fmt"

	"github.com/0xPolygon/cdk-data-availability/rpc"
	"github.com/0xPolygon/cdk-data-availability/types"
)

// BatchData is an abbreviated structure that only contains the number and L2 batch data
type BatchData struct {
	Number      types.ArgUint64 `json:"number"`
	BatchL2Data types.ArgBytes  `json:"batchL2Data,omitempty"`
	Empty       bool            `json:"empty"`
}

// BatchDataResult is a list of BatchData for a BatchFilter
type BatchDataResult struct {
	Data []*BatchData `json:"data"`
}

func BatchesByNumbers(url string, from, to uint64) ([]*BatchData, error) {
	var batchNumbers []string
	for i := from; i <= to; i++ {
		batchNumbers = append(batchNumbers, fmt.Sprintf("%v", i))
	}

	foo := make(map[string][]string, 0)
	foo["numbers"] = batchNumbers // nolint: gosec
	response, err := rpc.JSONRPCCall(url, "zkevm_getBatchDataByNumbers", foo)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {

		return nil, fmt.Errorf("%d - %s", response.Error.Code, response.Error.Message)
	}

	var result *BatchDataResult
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func BatchNumber(url string) (uint64, error) {
	response, err := rpc.JSONRPCCall(url, "zkevm_batchNumber")
	if err != nil {
		return 0, err
	}

	if response.Error != nil {

		return 0, fmt.Errorf("%d - %s", response.Error.Code, response.Error.Message)
	}

	var result string
	err = json.Unmarshal(response.Result, &result)
	if err != nil {
		return 0, err
	}

	bigBatchNumber := DecodeBig(result)

	batchNumber := bigBatchNumber.Uint64()

	return batchNumber, nil
}
