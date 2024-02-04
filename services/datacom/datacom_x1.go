package datacom

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

func (d *DataComEndpoints) isEmptyAddress(a common.Address) bool {
	emptyAddress := common.Address{}
	return bytes.Compare(a[:], emptyAddress[:]) == 0
}
