// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polygonvalidium

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PolygonRollupBaseEtrogBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonRollupBaseEtrogBatchData struct {
	Transactions         []byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonValidiumEtrogValidiumBatchData is an auto generated low-level Go binding around an user-defined struct.
type PolygonValidiumEtrogValidiumBatchData struct {
	TransactionsHash     [32]byte
	ForcedGlobalExitRoot [32]byte
	ForcedTimestamp      uint64
	ForcedBlockHashL1    [32]byte
}

// PolygonvalidiumMetaData contains all meta data concerning the Polygonvalidium contract.
var PolygonvalidiumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"_globalExitRootManager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"_pol\",\"type\":\"address\"},{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"_bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"_rollupManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BatchAlreadyVerified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchNotSequencedOrNotSequenceEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExceedMaxVerifyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchBelowLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalPendingStateNumInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesAlreadyActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesDecentralized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesNotAllowedOnEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForceBatchesOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ForcedDataDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasTokenNetworkMustBeZeroOnEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalExitRootNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HaltTimeoutNotExpiredAfterEmergencyState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HugeTokenMetadataNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchAboveLastVerifiedBatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitNumBatchDoesNotMatchPendingState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitSequencedBatchDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializeTransaction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeBatchTimeTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeForceBatchTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRangeMultiplierBatchFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTimestampSequenceInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewPendingStateTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewStateRootNotInsidePrime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewTrustedAggregatorTimeoutMustBeLower\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughMaticAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughPOLAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldAccInputHashDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OldStateRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPendingAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyRollupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyTrustedSequencer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateNotConsolidable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingStateTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceWithDataAvailabilityNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequenceZeroBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampBelowForcedTimestamp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SequencedTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredRootMustBeDifferentThanNewRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwitchToSameValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionsLengthAboveMax\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutExceedHaltAggregationTimeout\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TrustedAggregatorTimeoutNotExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AcceptAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"forceBatchNum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"}],\"name\":\"ForceBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"InitialSequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"l1InfoRoot\",\"type\":\"bytes32\"}],\"name\":\"SequenceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"}],\"name\":\"SequenceForceBatches\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"SetDataAvailabilityProtocol\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"SetForceBatchAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"SetForceBatchTimeout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"SetTrustedSequencer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"SetTrustedSequencerURL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SwitchSequenceWithDataAvailability\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdminRole\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"lastGlobalExitRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"}],\"name\":\"UpdateEtrogSequence\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"numBatch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"VerifyBatches\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"GLOBAL_EXIT_ROOT_MANAGER_L2\",\"outputs\":[{\"internalType\":\"contractIBasePolygonZkEVMGlobalExitRoot\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_LIST_LEN_LEN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_DATA_LEN_EMPTY_METADATA\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIALIZE_TX_EFFECTIVE_PERCENTAGE\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_UP_ETROG_TX\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_R\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_S\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_INITIALIZE_TX_V\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TIMESTAMP_RANGE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMBridgeV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePolPerForceBatch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataAvailabilityProtocol\",\"outputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"polAmount\",\"type\":\"uint256\"}],\"name\":\"forceBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceBatchTimeout\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"forcedBatches\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasTokenNetwork\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_gasTokenNetwork\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_gasTokenMetadata\",\"type\":\"bytes\"}],\"name\":\"generateInitializeTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalExitRootManager\",\"outputs\":[{\"internalType\":\"contractIPolygonZkEVMGlobalExitRootV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequencer\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"networkID\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_gasTokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"sequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedSequencer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_trustedSequencerURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_networkName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_lastAccInputHash\",\"type\":\"bytes32\"}],\"name\":\"initializeUpgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSequenceWithDataAvailabilityAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastAccInputHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatch\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastForceBatchSequenced\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"lastVerifiedBatch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"newStateRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"aggregator\",\"type\":\"address\"}],\"name\":\"onVerifyBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pol\",\"outputs\":[{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupManager\",\"outputs\":[{\"internalType\":\"contractPolygonRollupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"}],\"name\":\"sequenceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"transactionsHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonValidiumEtrog.ValidiumBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"maxSequenceTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"initSequencedBatch\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"l2Coinbase\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataAvailabilityMessage\",\"type\":\"bytes\"}],\"name\":\"sequenceBatchesValidium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"transactions\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"forcedGlobalExitRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"forcedTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"forcedBlockHashL1\",\"type\":\"bytes32\"}],\"internalType\":\"structPolygonRollupBaseEtrog.BatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"sequenceForceBatches\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIDataAvailabilityProtocol\",\"name\":\"newDataAvailabilityProtocol\",\"type\":\"address\"}],\"name\":\"setDataAvailabilityProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newForceBatchAddress\",\"type\":\"address\"}],\"name\":\"setForceBatchAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newforceBatchTimeout\",\"type\":\"uint64\"}],\"name\":\"setForceBatchTimeout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTrustedSequencer\",\"type\":\"address\"}],\"name\":\"setTrustedSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newTrustedSequencerURL\",\"type\":\"string\"}],\"name\":\"setTrustedSequencerURL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"newIsSequenceWithDataAvailabilityAllowed\",\"type\":\"bool\"}],\"name\":\"switchSequenceWithDataAvailability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"transferAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSequencerURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61010060405234801562000011575f80fd5b506040516200470b3803806200470b83398101604081905262000034916200006f565b6001600160a01b0393841660a052918316608052821660c0521660e052620000d4565b6001600160a01b03811681146200006c575f80fd5b50565b5f805f806080858703121562000083575f80fd5b8451620000908162000057565b6020860151909450620000a38162000057565b6040860151909350620000b68162000057565b6060860151909250620000c98162000057565b939692955090935050565b60805160a05160c05160e051614515620001f65f395f81816104c0015281816109ff01528181610af701528181610c3401528181610e85015281816110c2015281816114d8015281816118d201528181611c3e01528181611ceb01528181612640015281816126a9015281816126cb01528181612826015281816129c401528181612a7c015281816132e90152818161334c0152818161336e01526133ef01525f818161062d0152818161120d015281816112b701528181611e8001528181611f58015281816122600152612f3501525f81816106d601528181610d69015281816113be015281816122cf01528181612b910152612fa401525f818161072e015281816107e301528181611c7a0152818161275d0152612b6601526145155ff3fe608060405234801561000f575f80fd5b5060043610610324575f3560e01c80637a5460c5116101a8578063c7fffd4b116100f3578063def57e541161009e578063e7a7ed0211610079578063e7a7ed0214610763578063eaeb077b14610776578063f35dda4714610789578063f851a44014610791575f80fd5b8063def57e5414610716578063e46761c414610729578063e57a0b4c14610750575f80fd5b8063d02103ca116100ce578063d02103ca146106d1578063d7bc90ff146106f8578063db5b0ed714610703575f80fd5b8063c7fffd4b146106a3578063c89e42df146106ab578063cfa8ed47146106be575f80fd5b8063a3c573eb11610153578063af7f3e021161012e578063af7f3e0214610675578063b0afe1541461067d578063c754c7ed14610689575f80fd5b8063a3c573eb14610628578063a652f26c1461064f578063ada8f91914610662575f80fd5b806391cafe321161018357806391cafe32146105e75780639e001877146105fa5780639f26f84014610615575f80fd5b80637a5460c5146105ab5780637cd76b8b146105cc5780638c3d7301146105df575f80fd5b806342308fab11610273578063542028d51161021e5780636b8616ce116101f95780636b8616ce1461055d5780636e05d2cd1461057c5780636ff512cc146105855780637125702214610598575f80fd5b8063542028d51461053a5780635d6717a514610542578063676870d214610555575f80fd5b80634c21fef31161024e5780634c21fef3146104e25780634e4877061461050657806352bdeb6d14610519575f80fd5b806342308fab14610481578063456052671461048957806349b7b802146104bb575f80fd5b80632acdc2b6116102d35780633c351e10116102ae5780633c351e101461041a5780633cbc795b1461042d57806340b5de6c14610459575f80fd5b80632acdc2b6146103df5780632c111c06146103f457806332c2d15314610407575f80fd5b8063107bf28c11610303578063107bf28c1461039257806311e892d41461039a57806326782247146103b4575f80fd5b8062d0295d14610328578063035089631461034357806305835f371461035e575b5f80fd5b6103306107a9565b6040519081526020015b60405180910390f35b61034b602081565b60405161ffff909116815260200161033a565b61038560405180604001604052806008815260200167202021007270e02560c21b81525081565b60405161033a919061373d565b61038561089b565b6103a260f981565b60405160ff909116815260200161033a565b6001546103c7906001600160a01b031681565b6040516001600160a01b03909116815260200161033a565b6103f26103ed366004613766565b610927565b005b6008546103c7906001600160a01b031681565b6103f26104153660046137c4565b6109fd565b6009546103c7906001600160a01b031681565b60095461044490600160a01b900463ffffffff1681565b60405163ffffffff909116815260200161033a565b6104686001600160f81b031981565b6040516001600160f81b0319909116815260200161033a565b610330602481565b6007546104a390600160401b90046001600160401b031681565b6040516001600160401b03909116815260200161033a565b6103c77f000000000000000000000000000000000000000000000000000000000000000081565b603c546104f690600160a01b900460ff1681565b604051901515815260200161033a565b6103f2610514366004613803565b610a98565b61038560405180604001604052806002815260200161101760f31b81525081565b610385610c25565b6103f2610550366004613902565b610c32565b61034b601f81565b61033061056b366004613803565b60066020525f908152604090205481565b61033060055481565b6103f261059336600461398c565b611042565b6103f26105a63660046139b8565b6110c0565b6103856040518060400160405280600281526020016180b960f01b81525081565b6103f26105da36600461398c565b611693565b6103f2611711565b6103f26105f536600461398c565b6117c9565b6103c773a40d5f56745a118d0906a34e69aec8c0db1cb8fa81565b6103f2610623366004613aa5565b611889565b6103c77f000000000000000000000000000000000000000000000000000000000000000081565b61038561065d366004613ae3565b611db5565b6103f261067036600461398c565b6120b4565b610385612132565b6103306405ca1ab1e081565b6007546104a390600160801b90046001600160401b031681565b6103a260e481565b6103f26106b9366004613b53565b61214e565b6002546103c7906001600160a01b031681565b6103c77f000000000000000000000000000000000000000000000000000000000000000081565b610330635ca1ab1e81565b6103f2610711366004613bc1565b6121ba565b6103f2610724366004613c89565b612929565b6103c77f000000000000000000000000000000000000000000000000000000000000000081565b603c546103c7906001600160a01b031681565b6007546104a3906001600160401b031681565b6103f2610784366004613d00565b612980565b6103a2601b81565b5f546103c7906201000090046001600160a01b031681565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f9081906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610828573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061084c9190613d47565b6007549091505f90610870906001600160401b03600160401b820481169116613d72565b6001600160401b03169050805f0361088a575f9250505090565b6108948183613d99565b9250505090565b600480546108a890613db8565b80601f01602080910402602001604051908101604052809291908181526020018280546108d490613db8565b801561091f5780601f106108f65761010080835404028352916020019161091f565b820191905f5260205f20905b81548152906001019060200180831161090257829003601f168201915b505050505081565b5f546201000090046001600160a01b0316331461095757604051634755657960e01b815260040160405180910390fd5b603c54600160a01b900460ff161515811515036109a0576040517f5f0e7abe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b603c80547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16600160a01b831515021790556040517ff32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41905f90a150565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610a4657604051631736745960e31b815260040160405180910390fd5b806001600160a01b0316836001600160401b03167f9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f596684604051610a8b91815260200190565b60405180910390a3505050565b5f546201000090046001600160a01b03163314610ac857604051634755657960e01b815260040160405180910390fd5b62093a806001600160401b0382161115610af55760405163f5e37f2f60e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b51573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b759190613df0565b610baf576007546001600160401b03600160801b909104811690821610610baf5760405163f5e37f2f60e01b815260040160405180910390fd5b600780547fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff16600160801b6001600160401b038416908102919091179091556040519081527fa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b906020015b60405180910390a150565b600380546108a890613db8565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03163314610c7b57604051631736745960e31b815260040160405180910390fd5b5f54610100900460ff1615808015610c9957505f54600160ff909116105b80610cb25750303b158015610cb257505f5460ff166001145b610d1a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b5f805460ff191660011790558015610d3b575f805461ff0019166101001790555b5f6040518060a001604052806062815260200161447e6062913990505f818051906020012090505f4290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610dc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610de79190613d47565b90505f868483858d610dfa600143613e0b565b60408051602081019790975286019490945260608086019390935260c09190911b6001600160c01b0319166080850152901b6bffffffffffffffffffffffff1916608883015240609c82015260bc0160408051808303601f190181529082905280516020909101206005819055639a908e7360e01b8252600160048301526024820181905291505f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015610ed3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ef79190613e24565b90508b5f60026101000a8154816001600160a01b0302191690836001600160a01b031602179055508a60025f6101000a8154816001600160a01b0302191690836001600160a01b031602179055508960039081610f549190613e84565b506004610f618a82613e84565b508b60085f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555062069780600760106101000a8154816001600160401b0302191690836001600160401b031602179055507fd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c8187858e604051610fe79493929190613f3f565b60405180910390a1505050505050801561103a575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b5f546201000090046001600160a01b0316331461107257604051634755657960e01b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527ff54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc090602001610c1a565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316331461110957604051631736745960e31b815260040160405180910390fd5b5f54610100900460ff161580801561112757505f54600160ff909116105b806111405750303b15801561114057505f5460ff166001145b6111a35760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610d11565b5f805460ff1916600117905580156111c4575f805461ff0019166101001790555b60606001600160a01b03851615611381576040517fc00f14ab0000000000000000000000000000000000000000000000000000000081526001600160a01b0386811660048301527f0000000000000000000000000000000000000000000000000000000000000000169063c00f14ab906024015f60405180830381865afa158015611251573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526112789190810190613f80565b6040517f318aee3d0000000000000000000000000000000000000000000000000000000081526001600160a01b0387811660048301529192505f9182917f00000000000000000000000000000000000000000000000000000000000000009091169063318aee3d906024016040805180830381865afa1580156112fd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113219190613ff1565b915091508163ffffffff165f14611362576009805463ffffffff8416600160a01b026001600160c01b03199091166001600160a01b0384161717905561137e565b600980546001600160a01b0319166001600160a01b0389161790555b50505b6009545f906113aa9088906001600160a01b03811690600160a01b900463ffffffff1685611db5565b90505f818051906020012090505f4290505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611418573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061143c9190613d47565b90505f808483858f61144f600143613e0b565b60408051602081019790975286019490945260608086019390935260c09190911b6001600160c01b0319166080850152901b6bffffffffffffffffffffffff1916608883015240609c82015260bc0160408051808303601f190181529082905280516020909101206005819055639a908e7360e01b8252600160048301526024820181905291507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015611526573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061154a9190613e24565b508c5f60026101000a8154816001600160a01b0302191690836001600160a01b031602179055508b60025f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555088600390816115a69190613e84565b5060046115b38982613e84565b508c60085f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555062069780600760106101000a8154816001600160401b0302191690836001600160401b031602179055507f060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f85838e60405161163793929190614029565b60405180910390a1505050505050801561168a575f805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b5f546201000090046001600160a01b031633146116c357604051634755657960e01b815260040160405180910390fd5b603c80546001600160a01b0319166001600160a01b0383169081179091556040519081527fd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a90602001610c1a565b6001546001600160a01b03163314611755576040517fd1ec4b2300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001545f80547fffffffffffffffffffff0000000000000000000000000000000000000000ffff166001600160a01b039092166201000081029290921790556040519081527f056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e9060200160405180910390a1565b5f546201000090046001600160a01b031633146117f957604051634755657960e01b815260040160405180910390fd5b6008546001600160a01b031661183b576040517fc89374d800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527f5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb90602001610c1a565b6008546001600160a01b031680158015906118ad57506001600160a01b0381163314155b156118cb576040516324eff8c360e01b815260040160405180910390fd5b4262093a807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166330c27dde6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561192c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119509190613e24565b61195a919061405a565b6001600160401b0316111561199b576040517f3d49ed4c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815f8190036119bd5760405163cb591a5f60e01b815260040160405180910390fd5b6103e88111156119e057604051635acfba9d60e11b815260040160405180910390fd5b6007546001600160401b0380821691611a02918491600160401b90041661407a565b1115611a215760405163c630a00d60e01b815260040160405180910390fd5b600754600554600160401b9091046001600160401b0316905f5b83811015611c38575f878783818110611a5657611a5661408d565b9050602002810190611a6891906140a1565b611a71906140bf565b905083611a7d81614129565b825180516020918201208185015160408087015160608801519151959a509295505f94611ad1948794929101938452602084019290925260c01b6001600160c01b0319166040830152604882015260680190565b60408051601f1981840301815291815281516020928301206001600160401b0389165f90815260069093529120549091508114611b215760405163671ebaaf60e11b815260040160405180910390fd5b6001600160401b0386165f90815260066020526040812055611b44600188613e0b565b8403611bb15742600760109054906101000a90046001600160401b03168460400151611b70919061405a565b6001600160401b03161115611bb1576040517fc44a082100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208381015160408086015160608088015183519586018b90529285018790528481019390935260c01b6001600160c01b03191660808401523390911b6bffffffffffffffffffffffff19166088830152609c82015260bc016040516020818303038152906040528051906020012094505050508080611c309061414e565b915050611a3b565b50611ca17f000000000000000000000000000000000000000000000000000000000000000084611c666107a9565b611c709190614166565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169190612d8a565b6005819055600780546001600160401b038416600160401b026fffffffffffffffff000000000000000019909116179055604051639a908e7360e01b81525f906001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690639a908e7390611d3690879086906004016001600160401b03929092168252602082015260400190565b6020604051808303815f875af1158015611d52573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d769190613e24565b6040519091506001600160401b038216907f648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4905f90a250505050505050565b60605f85858573a40d5f56745a118d0906a34e69aec8c0db1cb8fa5f87604051602401611de79695949392919061417d565b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff811bff70000000000000000000000000000000000000000000000000000000017905283519091506060905f03611ee95760f9601f8351611e5d91906141d2565b60405180604001604052806008815260200167202021007270e02560c21b8152507f000000000000000000000000000000000000000000000000000000000000000060405180604001604052806002815260200161101760f31b81525060e487604051602001611ed397969594939291906141ed565b6040516020818303038152906040529050611fbd565b815161ffff1015611f26576040517f248b8f8200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b815160f9611f356020836141d2565b60405180604001604052806008815260200167202021007270e02560c21b8152507f00000000000000000000000000000000000000000000000000000000000000006040518060400160405280600281526020016180b960f01b8152508588604051602001611faa97969594939291906142a4565b6040516020818303038152906040529150505b8051602080830191909120604080515f80825293810180835292909252601b908201526405ca1ab1e06060820152635ca1ab1e608082015260019060a0016020604051602081039080840390855afa15801561201b573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116612068576040517fcd16196600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040515f906120959084906405ca1ab1e090635ca1ab1e90601b906001600160f81b03199060200161435b565b60408051601f198184030181529190529450505050505b949350505050565b5f546201000090046001600160a01b031633146120e457604051634755657960e01b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527fa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce690602001610c1a565b6040518060a001604052806062815260200161447e6062913981565b5f546201000090046001600160a01b0316331461217e57604051634755657960e01b815260040160405180910390fd5b600361218a8282613e84565b507f6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b2081604051610c1a919061373d565b6002546001600160a01b031633146121e5576040516311e7be1560e01b815260040160405180910390fd5b855f8190036122075760405163cb591a5f60e01b815260040160405180910390fd5b6103e881111561222a57604051635acfba9d60e11b815260040160405180910390fd5b61223560244261407a565b866001600160401b0316111561225e57604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b1580156122b6575f80fd5b505af11580156122c8573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015612329573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061234d9190613d47565b600754600554919250600160401b90046001600160401b031690815f805b868110156125cf575f8e8e838181106123865761238661408d565b90506080020180360381019061239c919061439e565b60408101519091506001600160401b03161561250b57856123bc81614129565b9650505f815f01518260200151836040015184606001516040516020016124099493929190938452602084019290925260c01b6001600160c01b0319166040830152604882015260680190565b60408051601f1981840301815291815281516020928301206001600160401b038a165f908152600690935291205490915081146124595760405163671ebaaf60e11b815260040160405180910390fd5b85825f0151836020015184604001518f86606001516040516020016124c7969594939291909586526020860194909452604085019290925260c01b6001600160c01b0319166060808501919091521b6bffffffffffffffffffffffff19166068830152607c820152609c0190565b60405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f9055506125bc565b8051604051612527918591602001918252602082015260400190565b60405160208183030381529060405280519060200120925084815f0151888f8e5f801b6040516020016125a3969594939291909586526020860194909452604085019290925260c01b6001600160c01b0319166060808501919091521b6bffffffffffffffffffffffff19166068830152607c820152609c0190565b6040516020818303038152906040528051906020012094505b50806125c78161414e565b91505061236b565b506007546001600160401b0390811690851611156126005760405163c630a00d60e01b815260040160405180910390fd5b6005839055856001600160401b038581169084161461269d575f6126248487613d72565b90506126396001600160401b03821683613e0b565b91506126717f0000000000000000000000000000000000000000000000000000000000000000826001600160401b0316611c666107a9565b50600780546fffffffffffffffff00000000000000001916600160401b6001600160401b038816021790555b80156127ff57612785337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612725573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127499190613d47565b6127539190614166565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016929190612e38565b603c546040517f3b51be4b0000000000000000000000000000000000000000000000000000000081526001600160a01b0390911690633b51be4b906127d29085908d908d90600401614412565b5f6040518083038186803b1580156127e8575f80fd5b505afa1580156127fa573d5f803e3d5ffd5b505050505b604051639a908e7360e01b81526001600160401b0388166004820152602481018590525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af1158015612874573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128989190613e24565b90506128a48882613d72565b6001600160401b03168c6001600160401b0316146128d557604051630d0386cd60e11b815260040160405180910390fd5b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e7668860405161291091815260200190565b60405180910390a2505050505050505050505050505050565b603c54600160a01b900460ff1661296c576040517f821935b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6129798585858585612e8f565b5050505050565b6008546001600160a01b031680158015906129a457506001600160a01b0381163314155b156129c2576040516324eff8c360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166315064c966040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a1e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a429190613df0565b15612a79576040517f39258d1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663604691696040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ad6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612afa9190613d47565b905082811115612b36576040517f2354600f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611388841115612b59576040516328a69b1f60e21b815260040160405180910390fd5b612b8e6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016333084612e38565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316633ed691ef6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612beb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c0f9190613d47565b600780549192506001600160401b03909116905f612c2c83614129565b91906101000a8154816001600160401b0302191690836001600160401b03160217905550508585604051612c6192919061442b565b6040519081900390208142612c77600143613e0b565b60408051602081019590955284019290925260c01b6001600160c01b031916606083015240606882015260880160408051601f1981840301815291815281516020928301206007546001600160401b03165f9081526006909352912055323303612d35576007546040805183815233602082015260608183018190525f9082015290516001600160401b03909216917ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc9319181900360800190a261103a565b6007546040516001600160401b03909116907ff94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc93190612d7a90849033908b908b9061443a565b60405180910390a2505050505050565b6040516001600160a01b038316602482015260448101829052612e339084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b60408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526134ef565b505050565b6040516001600160a01b0380851660248301528316604482015260648101829052612e899085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401612dcf565b50505050565b6002546001600160a01b03163314612eba576040516311e7be1560e01b815260040160405180910390fd5b835f819003612edc5760405163cb591a5f60e01b815260040160405180910390fd5b6103e8811115612eff57604051635acfba9d60e11b815260040160405180910390fd5b612f0a60244261407a565b846001600160401b03161115612f3357604051630a00feb360e01b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166379e2cf976040518163ffffffff1660e01b81526004015f604051808303815f87803b158015612f8b575f80fd5b505af1158015612f9d573d5f803e3d5ffd5b505050505f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316635ca1e1656040518163ffffffff1660e01b8152600401602060405180830381865afa158015612ffe573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130229190613d47565b600754600554919250600160401b90046001600160401b031690815f5b85811015613278575f8b8b8381811061305a5761305a61408d565b905060200281019061306c91906140a1565b613075906140bf565b805180516020909101206040820151919250906001600160401b0316156131d757856130a081614129565b9650505f818360200151846040015185606001516040516020016130ea9493929190938452602084019290925260c01b6001600160c01b0319166040830152604882015260680190565b60408051601f1981840301815291815281516020928301206001600160401b038a165f9081526006909352912054909150811461313a5760405163671ebaaf60e11b815260040160405180910390fd5b60208381015160408086015160608088015183519586018c90529285018790528481019390935260c01b6001600160c01b0319166080840152908c901b6bffffffffffffffffffffffff19166088830152609c82015260bc0160405160208183030381529060405280519060200120955060065f886001600160401b03166001600160401b031681526020019081526020015f205f905550613263565b8151516201d4c010156131fd576040516328a69b1f60e21b815260040160405180910390fd5b6040805160208101879052908101829052606080820189905260c08d901b6001600160c01b03191660808301528a901b6bffffffffffffffffffffffff191660888201525f609c82015260bc016040516020818303038152906040528051906020012094505b505080806132709061414e565b91505061303f565b506007546001600160401b0390811690841611156132a95760405163c630a00d60e01b815260040160405180910390fd5b6005829055846001600160401b0384811690831614613346575f6132cd8386613d72565b90506132e26001600160401b03821683613e0b565b915061331a7f0000000000000000000000000000000000000000000000000000000000000000826001600160401b0316611c666107a9565b50600780546fffffffffffffffff00000000000000001916600160401b6001600160401b038716021790555b6133c8337f0000000000000000000000000000000000000000000000000000000000000000837f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663477fa2706040518163ffffffff1660e01b8152600401602060405180830381865afa158015612725573d5f803e3d5ffd5b604051639a908e7360e01b81526001600160401b0387166004820152602481018490525f907f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031690639a908e73906044016020604051808303815f875af115801561343d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906134619190613e24565b905061346d8782613d72565b6001600160401b0316896001600160401b03161461349e57604051630d0386cd60e11b815260040160405180910390fd5b806001600160401b03167f3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766876040516134d991815260200190565b60405180910390a2505050505050505050505050565b5f613543826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166135d39092919063ffffffff16565b805190915015612e3357808060200190518101906135619190613df0565b612e335760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610d11565b60606120ac84845f85855f80866001600160a01b031685876040516135f8919061446c565b5f6040518083038185875af1925050503d805f8114613632576040519150601f19603f3d011682016040523d82523d5f602084013e613637565b606091505b509150915061364887838387613653565b979650505050505050565b606083156136c15782515f036136ba576001600160a01b0385163b6136ba5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610d11565b50816120ac565b6120ac83838151156136d65781518083602001fd5b8060405162461bcd60e51b8152600401610d11919061373d565b5f5b8381101561370a5781810151838201526020016136f2565b50505f910152565b5f81518084526137298160208601602086016136f0565b601f01601f19169290920160200192915050565b602081525f61374f6020830184613712565b9392505050565b8015158114613763575f80fd5b50565b5f60208284031215613776575f80fd5b813561374f81613756565b6001600160401b0381168114613763575f80fd5b80356137a081613781565b919050565b6001600160a01b0381168114613763575f80fd5b80356137a0816137a5565b5f805f606084860312156137d6575f80fd5b83356137e181613781565b92506020840135915060408401356137f8816137a5565b809150509250925092565b5f60208284031215613813575f80fd5b813561374f81613781565b634e487b7160e01b5f52604160045260245ffd5b604051608081016001600160401b03811182821017156138545761385461381e565b60405290565b604051601f8201601f191681016001600160401b03811182821017156138825761388261381e565b604052919050565b5f6001600160401b038211156138a2576138a261381e565b50601f01601f191660200190565b5f82601f8301126138bf575f80fd5b81356138d26138cd8261388a565b61385a565b8181528460208386010111156138e6575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f805f60a08688031215613916575f80fd5b8535613921816137a5565b94506020860135613931816137a5565b935060408601356001600160401b038082111561394c575f80fd5b61395889838a016138b0565b9450606088013591508082111561396d575f80fd5b5061397a888289016138b0565b95989497509295608001359392505050565b5f6020828403121561399c575f80fd5b813561374f816137a5565b63ffffffff81168114613763575f80fd5b5f805f805f8060c087890312156139cd575f80fd5b86356139d8816137a5565b955060208701356139e8816137a5565b945060408701356139f8816139a7565b93506060870135613a08816137a5565b925060808701356001600160401b0380821115613a23575f80fd5b613a2f8a838b016138b0565b935060a0890135915080821115613a44575f80fd5b50613a5189828a016138b0565b9150509295509295509295565b5f8083601f840112613a6e575f80fd5b5081356001600160401b03811115613a84575f80fd5b6020830191508360208260051b8501011115613a9e575f80fd5b9250929050565b5f8060208385031215613ab6575f80fd5b82356001600160401b03811115613acb575f80fd5b613ad785828601613a5e565b90969095509350505050565b5f805f8060808587031215613af6575f80fd5b8435613b01816139a7565b93506020850135613b11816137a5565b92506040850135613b21816139a7565b915060608501356001600160401b03811115613b3b575f80fd5b613b47878288016138b0565b91505092959194509250565b5f60208284031215613b63575f80fd5b81356001600160401b03811115613b78575f80fd5b6120ac848285016138b0565b5f8083601f840112613b94575f80fd5b5081356001600160401b03811115613baa575f80fd5b602083019150836020828501011115613a9e575f80fd5b5f805f805f805f60a0888a031215613bd7575f80fd5b87356001600160401b0380821115613bed575f80fd5b818a0191508a601f830112613c00575f80fd5b813581811115613c0e575f80fd5b8b60208260071b8501011115613c22575f80fd5b60208301995080985050613c3860208b01613795565b9650613c4660408b01613795565b9550613c5460608b016137b9565b945060808a0135915080821115613c69575f80fd5b50613c768a828b01613b84565b989b979a50959850939692959293505050565b5f805f805f60808688031215613c9d575f80fd5b85356001600160401b03811115613cb2575f80fd5b613cbe88828901613a5e565b9096509450506020860135613cd281613781565b92506040860135613ce281613781565b91506060860135613cf2816137a5565b809150509295509295909350565b5f805f60408486031215613d12575f80fd5b83356001600160401b03811115613d27575f80fd5b613d3386828701613b84565b909790965060209590950135949350505050565b5f60208284031215613d57575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b03828116828216039080821115613d9257613d92613d5e565b5092915050565b5f82613db357634e487b7160e01b5f52601260045260245ffd5b500490565b600181811c90821680613dcc57607f821691505b602082108103613dea57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215613e00575f80fd5b815161374f81613756565b81810381811115613e1e57613e1e613d5e565b92915050565b5f60208284031215613e34575f80fd5b815161374f81613781565b601f821115612e33575f81815260208120601f850160051c81016020861015613e655750805b601f850160051c820191505b8181101561103a57828155600101613e71565b81516001600160401b03811115613e9d57613e9d61381e565b613eb181613eab8454613db8565b84613e3f565b602080601f831160018114613ee4575f8415613ecd5750858301515b5f19600386901b1c1916600185901b17855561103a565b5f85815260208120601f198616915b82811015613f1257888601518255948401946001909101908401613ef3565b5085821015613f2f57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b6001600160401b0385168152608060208201525f613f606080830186613712565b90508360408301526001600160a01b038316606083015295945050505050565b5f60208284031215613f90575f80fd5b81516001600160401b03811115613fa5575f80fd5b8201601f81018413613fb5575f80fd5b8051613fc36138cd8261388a565b818152856020838501011115613fd7575f80fd5b613fe88260208301602086016136f0565b95945050505050565b5f8060408385031215614002575f80fd5b825161400d816139a7565b602084015190925061401e816137a5565b809150509250929050565b606081525f61403b6060830186613712565b90508360208301526001600160a01b0383166040830152949350505050565b6001600160401b03818116838216019080821115613d9257613d92613d5e565b80820180821115613e1e57613e1e613d5e565b634e487b7160e01b5f52603260045260245ffd5b5f8235607e198336030181126140b5575f80fd5b9190910192915050565b5f608082360312156140cf575f80fd5b6140d7613832565b82356001600160401b038111156140ec575f80fd5b6140f8368286016138b0565b82525060208301356020820152604083013561411381613781565b6040820152606092830135928101929092525090565b5f6001600160401b0380831681810361414457614144613d5e565b6001019392505050565b5f6001820161415f5761415f613d5e565b5060010190565b8082028115828204841417613e1e57613e1e613d5e565b5f63ffffffff80891683526001600160a01b038089166020850152818816604085015280871660608501528086166080850152505060c060a08301526141c660c0830184613712565b98975050505050505050565b61ffff818116838216019080821115613d9257613d92613d5e565b5f6001600160f81b0319808a60f81b1683527fffff0000000000000000000000000000000000000000000000000000000000008960f01b166001840152875161423d816003860160208c016136f0565b606088901b6bffffffffffffffffffffffff1916600391850191820152865161426d816017840160208b016136f0565b808201915050818660f81b166017820152845191506142938260188301602088016136f0565b016018019998505050505050505050565b6001600160f81b03198860f81b1681525f7fffff000000000000000000000000000000000000000000000000000000000000808960f01b16600184015287516142f4816003860160208c016136f0565b606088901b6bffffffffffffffffffffffff19166003918501918201528651614324816017840160208b016136f0565b808201915050818660f01b1660178201528451915061434a8260198301602088016136f0565b016019019998505050505050505050565b5f865161436c818460208b016136f0565b9190910194855250602084019290925260f81b6001600160f81b03199081166040840152166041820152604201919050565b5f608082840312156143ae575f80fd5b6143b6613832565b823581526020830135602082015260408301356143d281613781565b60408201526060928301359281019290925250919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b838152604060208201525f613fe86040830184866143ea565b818382375f9101908152919050565b8481526001600160a01b0384166020820152606060408201525f6144626060830184866143ea565b9695505050505050565b5f82516140b58184602087016136f056fedf2a8080944d5cf5032b2a844602278b01199ed191a86c93ff8080821092808000000000000000000000000000000000000000000000000000000005ca1ab1e000000000000000000000000000000000000000000000000000000005ca1ab1e01bffa2646970667358221220c2435e04e64c752d7f46402e477fd48b2cad6971e47bd2475071ebfa145bd68c64736f6c63430008140033",
}

// PolygonvalidiumABI is the input ABI used to generate the binding from.
// Deprecated: Use PolygonvalidiumMetaData.ABI instead.
var PolygonvalidiumABI = PolygonvalidiumMetaData.ABI

// PolygonvalidiumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolygonvalidiumMetaData.Bin instead.
var PolygonvalidiumBin = PolygonvalidiumMetaData.Bin

// DeployPolygonvalidium deploys a new Ethereum contract, binding an instance of Polygonvalidium to it.
func DeployPolygonvalidium(auth *bind.TransactOpts, backend bind.ContractBackend, _globalExitRootManager common.Address, _pol common.Address, _bridgeAddress common.Address, _rollupManager common.Address) (common.Address, *types.Transaction, *Polygonvalidium, error) {
	parsed, err := PolygonvalidiumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolygonvalidiumBin), backend, _globalExitRootManager, _pol, _bridgeAddress, _rollupManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Polygonvalidium{PolygonvalidiumCaller: PolygonvalidiumCaller{contract: contract}, PolygonvalidiumTransactor: PolygonvalidiumTransactor{contract: contract}, PolygonvalidiumFilterer: PolygonvalidiumFilterer{contract: contract}}, nil
}

// Polygonvalidium is an auto generated Go binding around an Ethereum contract.
type Polygonvalidium struct {
	PolygonvalidiumCaller     // Read-only binding to the contract
	PolygonvalidiumTransactor // Write-only binding to the contract
	PolygonvalidiumFilterer   // Log filterer for contract events
}

// PolygonvalidiumCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolygonvalidiumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolygonvalidiumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolygonvalidiumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolygonvalidiumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolygonvalidiumSession struct {
	Contract     *Polygonvalidium  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolygonvalidiumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolygonvalidiumCallerSession struct {
	Contract *PolygonvalidiumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PolygonvalidiumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolygonvalidiumTransactorSession struct {
	Contract     *PolygonvalidiumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PolygonvalidiumRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolygonvalidiumRaw struct {
	Contract *Polygonvalidium // Generic contract binding to access the raw methods on
}

// PolygonvalidiumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolygonvalidiumCallerRaw struct {
	Contract *PolygonvalidiumCaller // Generic read-only contract binding to access the raw methods on
}

// PolygonvalidiumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolygonvalidiumTransactorRaw struct {
	Contract *PolygonvalidiumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolygonvalidium creates a new instance of Polygonvalidium, bound to a specific deployed contract.
func NewPolygonvalidium(address common.Address, backend bind.ContractBackend) (*Polygonvalidium, error) {
	contract, err := bindPolygonvalidium(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Polygonvalidium{PolygonvalidiumCaller: PolygonvalidiumCaller{contract: contract}, PolygonvalidiumTransactor: PolygonvalidiumTransactor{contract: contract}, PolygonvalidiumFilterer: PolygonvalidiumFilterer{contract: contract}}, nil
}

// NewPolygonvalidiumCaller creates a new read-only instance of Polygonvalidium, bound to a specific deployed contract.
func NewPolygonvalidiumCaller(address common.Address, caller bind.ContractCaller) (*PolygonvalidiumCaller, error) {
	contract, err := bindPolygonvalidium(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumCaller{contract: contract}, nil
}

// NewPolygonvalidiumTransactor creates a new write-only instance of Polygonvalidium, bound to a specific deployed contract.
func NewPolygonvalidiumTransactor(address common.Address, transactor bind.ContractTransactor) (*PolygonvalidiumTransactor, error) {
	contract, err := bindPolygonvalidium(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumTransactor{contract: contract}, nil
}

// NewPolygonvalidiumFilterer creates a new log filterer instance of Polygonvalidium, bound to a specific deployed contract.
func NewPolygonvalidiumFilterer(address common.Address, filterer bind.ContractFilterer) (*PolygonvalidiumFilterer, error) {
	contract, err := bindPolygonvalidium(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumFilterer{contract: contract}, nil
}

// bindPolygonvalidium binds a generic wrapper to an already deployed contract.
func bindPolygonvalidium(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolygonvalidiumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidium *PolygonvalidiumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidium.Contract.PolygonvalidiumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidium *PolygonvalidiumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.PolygonvalidiumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidium *PolygonvalidiumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.PolygonvalidiumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Polygonvalidium *PolygonvalidiumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Polygonvalidium.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Polygonvalidium *PolygonvalidiumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Polygonvalidium *PolygonvalidiumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.contract.Transact(opts, method, params...)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) GLOBALEXITROOTMANAGERL2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "GLOBAL_EXIT_ROOT_MANAGER_L2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidium.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidium.CallOpts)
}

// GLOBALEXITROOTMANAGERL2 is a free data retrieval call binding the contract method 0x9e001877.
//
// Solidity: function GLOBAL_EXIT_ROOT_MANAGER_L2() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) GLOBALEXITROOTMANAGERL2() (common.Address, error) {
	return _Polygonvalidium.Contract.GLOBALEXITROOTMANAGERL2(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXBRIDGELISTLENLEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_LIST_LEN_LEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGELISTLENLEN is a free data retrieval call binding the contract method 0x11e892d4.
//
// Solidity: function INITIALIZE_TX_BRIDGE_LIST_LEN_LEN() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXBRIDGELISTLENLEN() (uint8, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGELISTLENLEN(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXBRIDGEPARAMS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMS is a free data retrieval call binding the contract method 0x05835f37.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXBRIDGEPARAMS() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMS(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS is a free data retrieval call binding the contract method 0x7a5460c5.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESS(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA is a free data retrieval call binding the contract method 0x52bdeb6d.
//
// Solidity: function INITIALIZE_TX_BRIDGE_PARAMS_AFTER_BRIDGE_ADDRESS_EMPTY_METADATA() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA() ([]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXBRIDGEPARAMSAFTERBRIDGEADDRESSEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXCONSTANTBYTES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidium.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTES is a free data retrieval call binding the contract method 0x03508963.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXCONSTANTBYTES() (uint16, error) {
	return _Polygonvalidium.Contract.INITIALIZETXCONSTANTBYTES(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXCONSTANTBYTESEMPTYMETADATA(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidium.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXCONSTANTBYTESEMPTYMETADATA is a free data retrieval call binding the contract method 0x676870d2.
//
// Solidity: function INITIALIZE_TX_CONSTANT_BYTES_EMPTY_METADATA() view returns(uint16)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXCONSTANTBYTESEMPTYMETADATA() (uint16, error) {
	return _Polygonvalidium.Contract.INITIALIZETXCONSTANTBYTESEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXDATALENEMPTYMETADATA(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_DATA_LEN_EMPTY_METADATA")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidium.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXDATALENEMPTYMETADATA is a free data retrieval call binding the contract method 0xc7fffd4b.
//
// Solidity: function INITIALIZE_TX_DATA_LEN_EMPTY_METADATA() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXDATALENEMPTYMETADATA() (uint8, error) {
	return _Polygonvalidium.Contract.INITIALIZETXDATALENEMPTYMETADATA(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidium *PolygonvalidiumCaller) INITIALIZETXEFFECTIVEPERCENTAGE(opts *bind.CallOpts) ([1]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "INITIALIZE_TX_EFFECTIVE_PERCENTAGE")

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidium *PolygonvalidiumSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidium.CallOpts)
}

// INITIALIZETXEFFECTIVEPERCENTAGE is a free data retrieval call binding the contract method 0x40b5de6c.
//
// Solidity: function INITIALIZE_TX_EFFECTIVE_PERCENTAGE() view returns(bytes1)
func (_Polygonvalidium *PolygonvalidiumCallerSession) INITIALIZETXEFFECTIVEPERCENTAGE() ([1]byte, error) {
	return _Polygonvalidium.Contract.INITIALIZETXEFFECTIVEPERCENTAGE(&_Polygonvalidium.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCaller) SETUPETROGTX(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "SET_UP_ETROG_TX")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonvalidium.Contract.SETUPETROGTX(&_Polygonvalidium.CallOpts)
}

// SETUPETROGTX is a free data retrieval call binding the contract method 0xaf7f3e02.
//
// Solidity: function SET_UP_ETROG_TX() view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCallerSession) SETUPETROGTX() ([]byte, error) {
	return _Polygonvalidium.Contract.SETUPETROGTX(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCaller) SIGNATUREINITIALIZETXR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_R")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXR is a free data retrieval call binding the contract method 0xb0afe154.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_R() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCallerSession) SIGNATUREINITIALIZETXR() ([32]byte, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXR(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCaller) SIGNATUREINITIALIZETXS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_S")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXS is a free data retrieval call binding the contract method 0xd7bc90ff.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_S() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCallerSession) SIGNATUREINITIALIZETXS() ([32]byte, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXS(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCaller) SIGNATUREINITIALIZETXV(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "SIGNATURE_INITIALIZE_TX_V")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidium.CallOpts)
}

// SIGNATUREINITIALIZETXV is a free data retrieval call binding the contract method 0xf35dda47.
//
// Solidity: function SIGNATURE_INITIALIZE_TX_V() view returns(uint8)
func (_Polygonvalidium *PolygonvalidiumCallerSession) SIGNATUREINITIALIZETXV() (uint8, error) {
	return _Polygonvalidium.Contract.SIGNATUREINITIALIZETXV(&_Polygonvalidium.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumCaller) TIMESTAMPRANGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "TIMESTAMP_RANGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidium.Contract.TIMESTAMPRANGE(&_Polygonvalidium.CallOpts)
}

// TIMESTAMPRANGE is a free data retrieval call binding the contract method 0x42308fab.
//
// Solidity: function TIMESTAMP_RANGE() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumCallerSession) TIMESTAMPRANGE() (*big.Int, error) {
	return _Polygonvalidium.Contract.TIMESTAMPRANGE(&_Polygonvalidium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) Admin() (common.Address, error) {
	return _Polygonvalidium.Contract.Admin(&_Polygonvalidium.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) Admin() (common.Address, error) {
	return _Polygonvalidium.Contract.Admin(&_Polygonvalidium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.BridgeAddress(&_Polygonvalidium.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) BridgeAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.BridgeAddress(&_Polygonvalidium.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumCaller) CalculatePolPerForceBatch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "calculatePolPerForceBatch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidium.Contract.CalculatePolPerForceBatch(&_Polygonvalidium.CallOpts)
}

// CalculatePolPerForceBatch is a free data retrieval call binding the contract method 0x00d0295d.
//
// Solidity: function calculatePolPerForceBatch() view returns(uint256)
func (_Polygonvalidium *PolygonvalidiumCallerSession) CalculatePolPerForceBatch() (*big.Int, error) {
	return _Polygonvalidium.Contract.CalculatePolPerForceBatch(&_Polygonvalidium.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) DataAvailabilityProtocol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "dataAvailabilityProtocol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidium.Contract.DataAvailabilityProtocol(&_Polygonvalidium.CallOpts)
}

// DataAvailabilityProtocol is a free data retrieval call binding the contract method 0xe57a0b4c.
//
// Solidity: function dataAvailabilityProtocol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) DataAvailabilityProtocol() (common.Address, error) {
	return _Polygonvalidium.Contract.DataAvailabilityProtocol(&_Polygonvalidium.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) ForceBatchAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "forceBatchAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.ForceBatchAddress(&_Polygonvalidium.CallOpts)
}

// ForceBatchAddress is a free data retrieval call binding the contract method 0x2c111c06.
//
// Solidity: function forceBatchAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) ForceBatchAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.ForceBatchAddress(&_Polygonvalidium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCaller) ForceBatchTimeout(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "forceBatchTimeout")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidium.Contract.ForceBatchTimeout(&_Polygonvalidium.CallOpts)
}

// ForceBatchTimeout is a free data retrieval call binding the contract method 0xc754c7ed.
//
// Solidity: function forceBatchTimeout() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCallerSession) ForceBatchTimeout() (uint64, error) {
	return _Polygonvalidium.Contract.ForceBatchTimeout(&_Polygonvalidium.CallOpts)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCaller) ForcedBatches(opts *bind.CallOpts, arg0 uint64) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "forcedBatches", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidium.Contract.ForcedBatches(&_Polygonvalidium.CallOpts, arg0)
}

// ForcedBatches is a free data retrieval call binding the contract method 0x6b8616ce.
//
// Solidity: function forcedBatches(uint64 ) view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCallerSession) ForcedBatches(arg0 uint64) ([32]byte, error) {
	return _Polygonvalidium.Contract.ForcedBatches(&_Polygonvalidium.CallOpts, arg0)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) GasTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "gasTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.GasTokenAddress(&_Polygonvalidium.CallOpts)
}

// GasTokenAddress is a free data retrieval call binding the contract method 0x3c351e10.
//
// Solidity: function gasTokenAddress() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) GasTokenAddress() (common.Address, error) {
	return _Polygonvalidium.Contract.GasTokenAddress(&_Polygonvalidium.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidium *PolygonvalidiumCaller) GasTokenNetwork(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "gasTokenNetwork")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidium *PolygonvalidiumSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidium.Contract.GasTokenNetwork(&_Polygonvalidium.CallOpts)
}

// GasTokenNetwork is a free data retrieval call binding the contract method 0x3cbc795b.
//
// Solidity: function gasTokenNetwork() view returns(uint32)
func (_Polygonvalidium *PolygonvalidiumCallerSession) GasTokenNetwork() (uint32, error) {
	return _Polygonvalidium.Contract.GasTokenNetwork(&_Polygonvalidium.CallOpts)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCaller) GenerateInitializeTransaction(opts *bind.CallOpts, networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "generateInitializeTransaction", networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidium.Contract.GenerateInitializeTransaction(&_Polygonvalidium.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GenerateInitializeTransaction is a free data retrieval call binding the contract method 0xa652f26c.
//
// Solidity: function generateInitializeTransaction(uint32 networkID, address _gasTokenAddress, uint32 _gasTokenNetwork, bytes _gasTokenMetadata) view returns(bytes)
func (_Polygonvalidium *PolygonvalidiumCallerSession) GenerateInitializeTransaction(networkID uint32, _gasTokenAddress common.Address, _gasTokenNetwork uint32, _gasTokenMetadata []byte) ([]byte, error) {
	return _Polygonvalidium.Contract.GenerateInitializeTransaction(&_Polygonvalidium.CallOpts, networkID, _gasTokenAddress, _gasTokenNetwork, _gasTokenMetadata)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) GlobalExitRootManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "globalExitRootManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidium.Contract.GlobalExitRootManager(&_Polygonvalidium.CallOpts)
}

// GlobalExitRootManager is a free data retrieval call binding the contract method 0xd02103ca.
//
// Solidity: function globalExitRootManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) GlobalExitRootManager() (common.Address, error) {
	return _Polygonvalidium.Contract.GlobalExitRootManager(&_Polygonvalidium.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidium *PolygonvalidiumCaller) IsSequenceWithDataAvailabilityAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "isSequenceWithDataAvailabilityAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidium *PolygonvalidiumSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidium.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidium.CallOpts)
}

// IsSequenceWithDataAvailabilityAllowed is a free data retrieval call binding the contract method 0x4c21fef3.
//
// Solidity: function isSequenceWithDataAvailabilityAllowed() view returns(bool)
func (_Polygonvalidium *PolygonvalidiumCallerSession) IsSequenceWithDataAvailabilityAllowed() (bool, error) {
	return _Polygonvalidium.Contract.IsSequenceWithDataAvailabilityAllowed(&_Polygonvalidium.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCaller) LastAccInputHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "lastAccInputHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidium.Contract.LastAccInputHash(&_Polygonvalidium.CallOpts)
}

// LastAccInputHash is a free data retrieval call binding the contract method 0x6e05d2cd.
//
// Solidity: function lastAccInputHash() view returns(bytes32)
func (_Polygonvalidium *PolygonvalidiumCallerSession) LastAccInputHash() ([32]byte, error) {
	return _Polygonvalidium.Contract.LastAccInputHash(&_Polygonvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCaller) LastForceBatch(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "lastForceBatch")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidium.Contract.LastForceBatch(&_Polygonvalidium.CallOpts)
}

// LastForceBatch is a free data retrieval call binding the contract method 0xe7a7ed02.
//
// Solidity: function lastForceBatch() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCallerSession) LastForceBatch() (uint64, error) {
	return _Polygonvalidium.Contract.LastForceBatch(&_Polygonvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCaller) LastForceBatchSequenced(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "lastForceBatchSequenced")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidium.Contract.LastForceBatchSequenced(&_Polygonvalidium.CallOpts)
}

// LastForceBatchSequenced is a free data retrieval call binding the contract method 0x45605267.
//
// Solidity: function lastForceBatchSequenced() view returns(uint64)
func (_Polygonvalidium *PolygonvalidiumCallerSession) LastForceBatchSequenced() (uint64, error) {
	return _Polygonvalidium.Contract.LastForceBatchSequenced(&_Polygonvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidium *PolygonvalidiumCaller) NetworkName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "networkName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidium *PolygonvalidiumSession) NetworkName() (string, error) {
	return _Polygonvalidium.Contract.NetworkName(&_Polygonvalidium.CallOpts)
}

// NetworkName is a free data retrieval call binding the contract method 0x107bf28c.
//
// Solidity: function networkName() view returns(string)
func (_Polygonvalidium *PolygonvalidiumCallerSession) NetworkName() (string, error) {
	return _Polygonvalidium.Contract.NetworkName(&_Polygonvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidium.Contract.PendingAdmin(&_Polygonvalidium.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) PendingAdmin() (common.Address, error) {
	return _Polygonvalidium.Contract.PendingAdmin(&_Polygonvalidium.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) Pol(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "pol")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) Pol() (common.Address, error) {
	return _Polygonvalidium.Contract.Pol(&_Polygonvalidium.CallOpts)
}

// Pol is a free data retrieval call binding the contract method 0xe46761c4.
//
// Solidity: function pol() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) Pol() (common.Address, error) {
	return _Polygonvalidium.Contract.Pol(&_Polygonvalidium.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) RollupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "rollupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) RollupManager() (common.Address, error) {
	return _Polygonvalidium.Contract.RollupManager(&_Polygonvalidium.CallOpts)
}

// RollupManager is a free data retrieval call binding the contract method 0x49b7b802.
//
// Solidity: function rollupManager() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) RollupManager() (common.Address, error) {
	return _Polygonvalidium.Contract.RollupManager(&_Polygonvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCaller) TrustedSequencer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "trustedSequencer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidium *PolygonvalidiumSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidium.Contract.TrustedSequencer(&_Polygonvalidium.CallOpts)
}

// TrustedSequencer is a free data retrieval call binding the contract method 0xcfa8ed47.
//
// Solidity: function trustedSequencer() view returns(address)
func (_Polygonvalidium *PolygonvalidiumCallerSession) TrustedSequencer() (common.Address, error) {
	return _Polygonvalidium.Contract.TrustedSequencer(&_Polygonvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidium *PolygonvalidiumCaller) TrustedSequencerURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Polygonvalidium.contract.Call(opts, &out, "trustedSequencerURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidium *PolygonvalidiumSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidium.Contract.TrustedSequencerURL(&_Polygonvalidium.CallOpts)
}

// TrustedSequencerURL is a free data retrieval call binding the contract method 0x542028d5.
//
// Solidity: function trustedSequencerURL() view returns(string)
func (_Polygonvalidium *PolygonvalidiumCallerSession) TrustedSequencerURL() (string, error) {
	return _Polygonvalidium.Contract.TrustedSequencerURL(&_Polygonvalidium.CallOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) AcceptAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "acceptAdminRole")
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidium *PolygonvalidiumSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidium.Contract.AcceptAdminRole(&_Polygonvalidium.TransactOpts)
}

// AcceptAdminRole is a paid mutator transaction binding the contract method 0x8c3d7301.
//
// Solidity: function acceptAdminRole() returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) AcceptAdminRole() (*types.Transaction, error) {
	return _Polygonvalidium.Contract.AcceptAdminRole(&_Polygonvalidium.TransactOpts)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) ForceBatch(opts *bind.TransactOpts, transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "forceBatch", transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidium *PolygonvalidiumSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.ForceBatch(&_Polygonvalidium.TransactOpts, transactions, polAmount)
}

// ForceBatch is a paid mutator transaction binding the contract method 0xeaeb077b.
//
// Solidity: function forceBatch(bytes transactions, uint256 polAmount) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) ForceBatch(transactions []byte, polAmount *big.Int) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.ForceBatch(&_Polygonvalidium.TransactOpts, transactions, polAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "initialize", _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidium *PolygonvalidiumSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.Initialize(&_Polygonvalidium.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// Initialize is a paid mutator transaction binding the contract method 0x71257022.
//
// Solidity: function initialize(address _admin, address sequencer, uint32 networkID, address _gasTokenAddress, string sequencerURL, string _networkName) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) Initialize(_admin common.Address, sequencer common.Address, networkID uint32, _gasTokenAddress common.Address, sequencerURL string, _networkName string) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.Initialize(&_Polygonvalidium.TransactOpts, _admin, sequencer, networkID, _gasTokenAddress, sequencerURL, _networkName)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) InitializeUpgrade(opts *bind.TransactOpts, _admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "initializeUpgrade", _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonvalidium *PolygonvalidiumSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.InitializeUpgrade(&_Polygonvalidium.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// InitializeUpgrade is a paid mutator transaction binding the contract method 0x5d6717a5.
//
// Solidity: function initializeUpgrade(address _admin, address _trustedSequencer, string _trustedSequencerURL, string _networkName, bytes32 _lastAccInputHash) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) InitializeUpgrade(_admin common.Address, _trustedSequencer common.Address, _trustedSequencerURL string, _networkName string, _lastAccInputHash [32]byte) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.InitializeUpgrade(&_Polygonvalidium.TransactOpts, _admin, _trustedSequencer, _trustedSequencerURL, _networkName, _lastAccInputHash)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) OnVerifyBatches(opts *bind.TransactOpts, lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "onVerifyBatches", lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidium *PolygonvalidiumSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.OnVerifyBatches(&_Polygonvalidium.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// OnVerifyBatches is a paid mutator transaction binding the contract method 0x32c2d153.
//
// Solidity: function onVerifyBatches(uint64 lastVerifiedBatch, bytes32 newStateRoot, address aggregator) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) OnVerifyBatches(lastVerifiedBatch uint64, newStateRoot [32]byte, aggregator common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.OnVerifyBatches(&_Polygonvalidium.TransactOpts, lastVerifiedBatch, newStateRoot, aggregator)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SequenceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "sequenceBatches", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceBatches(&_Polygonvalidium.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatches is a paid mutator transaction binding the contract method 0xdef57e54.
//
// Solidity: function sequenceBatches((bytes,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SequenceBatches(batches []PolygonRollupBaseEtrogBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceBatches(&_Polygonvalidium.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SequenceBatchesValidium(opts *bind.TransactOpts, batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "sequenceBatchesValidium", batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceBatchesValidium(&_Polygonvalidium.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceBatchesValidium is a paid mutator transaction binding the contract method 0xdb5b0ed7.
//
// Solidity: function sequenceBatchesValidium((bytes32,bytes32,uint64,bytes32)[] batches, uint64 maxSequenceTimestamp, uint64 initSequencedBatch, address l2Coinbase, bytes dataAvailabilityMessage) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SequenceBatchesValidium(batches []PolygonValidiumEtrogValidiumBatchData, maxSequenceTimestamp uint64, initSequencedBatch uint64, l2Coinbase common.Address, dataAvailabilityMessage []byte) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceBatchesValidium(&_Polygonvalidium.TransactOpts, batches, maxSequenceTimestamp, initSequencedBatch, l2Coinbase, dataAvailabilityMessage)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SequenceForceBatches(opts *bind.TransactOpts, batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "sequenceForceBatches", batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceForceBatches(&_Polygonvalidium.TransactOpts, batches)
}

// SequenceForceBatches is a paid mutator transaction binding the contract method 0x9f26f840.
//
// Solidity: function sequenceForceBatches((bytes,bytes32,uint64,bytes32)[] batches) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SequenceForceBatches(batches []PolygonRollupBaseEtrogBatchData) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SequenceForceBatches(&_Polygonvalidium.TransactOpts, batches)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SetDataAvailabilityProtocol(opts *bind.TransactOpts, newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "setDataAvailabilityProtocol", newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetDataAvailabilityProtocol(&_Polygonvalidium.TransactOpts, newDataAvailabilityProtocol)
}

// SetDataAvailabilityProtocol is a paid mutator transaction binding the contract method 0x7cd76b8b.
//
// Solidity: function setDataAvailabilityProtocol(address newDataAvailabilityProtocol) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SetDataAvailabilityProtocol(newDataAvailabilityProtocol common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetDataAvailabilityProtocol(&_Polygonvalidium.TransactOpts, newDataAvailabilityProtocol)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SetForceBatchAddress(opts *bind.TransactOpts, newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "setForceBatchAddress", newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetForceBatchAddress(&_Polygonvalidium.TransactOpts, newForceBatchAddress)
}

// SetForceBatchAddress is a paid mutator transaction binding the contract method 0x91cafe32.
//
// Solidity: function setForceBatchAddress(address newForceBatchAddress) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SetForceBatchAddress(newForceBatchAddress common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetForceBatchAddress(&_Polygonvalidium.TransactOpts, newForceBatchAddress)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SetForceBatchTimeout(opts *bind.TransactOpts, newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "setForceBatchTimeout", newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetForceBatchTimeout(&_Polygonvalidium.TransactOpts, newforceBatchTimeout)
}

// SetForceBatchTimeout is a paid mutator transaction binding the contract method 0x4e487706.
//
// Solidity: function setForceBatchTimeout(uint64 newforceBatchTimeout) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SetForceBatchTimeout(newforceBatchTimeout uint64) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetForceBatchTimeout(&_Polygonvalidium.TransactOpts, newforceBatchTimeout)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SetTrustedSequencer(opts *bind.TransactOpts, newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "setTrustedSequencer", newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetTrustedSequencer(&_Polygonvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencer is a paid mutator transaction binding the contract method 0x6ff512cc.
//
// Solidity: function setTrustedSequencer(address newTrustedSequencer) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SetTrustedSequencer(newTrustedSequencer common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetTrustedSequencer(&_Polygonvalidium.TransactOpts, newTrustedSequencer)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SetTrustedSequencerURL(opts *bind.TransactOpts, newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "setTrustedSequencerURL", newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetTrustedSequencerURL(&_Polygonvalidium.TransactOpts, newTrustedSequencerURL)
}

// SetTrustedSequencerURL is a paid mutator transaction binding the contract method 0xc89e42df.
//
// Solidity: function setTrustedSequencerURL(string newTrustedSequencerURL) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SetTrustedSequencerURL(newTrustedSequencerURL string) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SetTrustedSequencerURL(&_Polygonvalidium.TransactOpts, newTrustedSequencerURL)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) SwitchSequenceWithDataAvailability(opts *bind.TransactOpts, newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "switchSequenceWithDataAvailability", newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidium *PolygonvalidiumSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidium.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// SwitchSequenceWithDataAvailability is a paid mutator transaction binding the contract method 0x2acdc2b6.
//
// Solidity: function switchSequenceWithDataAvailability(bool newIsSequenceWithDataAvailabilityAllowed) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) SwitchSequenceWithDataAvailability(newIsSequenceWithDataAvailabilityAllowed bool) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.SwitchSequenceWithDataAvailability(&_Polygonvalidium.TransactOpts, newIsSequenceWithDataAvailabilityAllowed)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidium *PolygonvalidiumTransactor) TransferAdminRole(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.contract.Transact(opts, "transferAdminRole", newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidium *PolygonvalidiumSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.TransferAdminRole(&_Polygonvalidium.TransactOpts, newPendingAdmin)
}

// TransferAdminRole is a paid mutator transaction binding the contract method 0xada8f919.
//
// Solidity: function transferAdminRole(address newPendingAdmin) returns()
func (_Polygonvalidium *PolygonvalidiumTransactorSession) TransferAdminRole(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Polygonvalidium.Contract.TransferAdminRole(&_Polygonvalidium.TransactOpts, newPendingAdmin)
}

// PolygonvalidiumAcceptAdminRoleIterator is returned from FilterAcceptAdminRole and is used to iterate over the raw logs and unpacked data for AcceptAdminRole events raised by the Polygonvalidium contract.
type PolygonvalidiumAcceptAdminRoleIterator struct {
	Event *PolygonvalidiumAcceptAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumAcceptAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumAcceptAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumAcceptAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumAcceptAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumAcceptAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumAcceptAdminRole represents a AcceptAdminRole event raised by the Polygonvalidium contract.
type PolygonvalidiumAcceptAdminRole struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAcceptAdminRole is a free log retrieval operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterAcceptAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumAcceptAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumAcceptAdminRoleIterator{contract: _Polygonvalidium.contract, event: "AcceptAdminRole", logs: logs, sub: sub}, nil
}

// WatchAcceptAdminRole is a free log subscription operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchAcceptAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumAcceptAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "AcceptAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumAcceptAdminRole)
				if err := _Polygonvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAcceptAdminRole is a log parse operation binding the contract event 0x056dc487bbf0795d0bbb1b4f0af523a855503cff740bfb4d5475f7a90c091e8e.
//
// Solidity: event AcceptAdminRole(address newAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseAcceptAdminRole(log types.Log) (*PolygonvalidiumAcceptAdminRole, error) {
	event := new(PolygonvalidiumAcceptAdminRole)
	if err := _Polygonvalidium.contract.UnpackLog(event, "AcceptAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumForceBatchIterator is returned from FilterForceBatch and is used to iterate over the raw logs and unpacked data for ForceBatch events raised by the Polygonvalidium contract.
type PolygonvalidiumForceBatchIterator struct {
	Event *PolygonvalidiumForceBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumForceBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumForceBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumForceBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumForceBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumForceBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumForceBatch represents a ForceBatch event raised by the Polygonvalidium contract.
type PolygonvalidiumForceBatch struct {
	ForceBatchNum      uint64
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Transactions       []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterForceBatch is a free log retrieval operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterForceBatch(opts *bind.FilterOpts, forceBatchNum []uint64) (*PolygonvalidiumForceBatchIterator, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumForceBatchIterator{contract: _Polygonvalidium.contract, event: "ForceBatch", logs: logs, sub: sub}, nil
}

// WatchForceBatch is a free log subscription operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchForceBatch(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumForceBatch, forceBatchNum []uint64) (event.Subscription, error) {

	var forceBatchNumRule []interface{}
	for _, forceBatchNumItem := range forceBatchNum {
		forceBatchNumRule = append(forceBatchNumRule, forceBatchNumItem)
	}

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "ForceBatch", forceBatchNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumForceBatch)
				if err := _Polygonvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForceBatch is a log parse operation binding the contract event 0xf94bb37db835f1ab585ee00041849a09b12cd081d77fa15ca070757619cbc931.
//
// Solidity: event ForceBatch(uint64 indexed forceBatchNum, bytes32 lastGlobalExitRoot, address sequencer, bytes transactions)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseForceBatch(log types.Log) (*PolygonvalidiumForceBatch, error) {
	event := new(PolygonvalidiumForceBatch)
	if err := _Polygonvalidium.contract.UnpackLog(event, "ForceBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumInitialSequenceBatchesIterator is returned from FilterInitialSequenceBatches and is used to iterate over the raw logs and unpacked data for InitialSequenceBatches events raised by the Polygonvalidium contract.
type PolygonvalidiumInitialSequenceBatchesIterator struct {
	Event *PolygonvalidiumInitialSequenceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumInitialSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumInitialSequenceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumInitialSequenceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumInitialSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumInitialSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumInitialSequenceBatches represents a InitialSequenceBatches event raised by the Polygonvalidium contract.
type PolygonvalidiumInitialSequenceBatches struct {
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitialSequenceBatches is a free log retrieval operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterInitialSequenceBatches(opts *bind.FilterOpts) (*PolygonvalidiumInitialSequenceBatchesIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumInitialSequenceBatchesIterator{contract: _Polygonvalidium.contract, event: "InitialSequenceBatches", logs: logs, sub: sub}, nil
}

// WatchInitialSequenceBatches is a free log subscription operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchInitialSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumInitialSequenceBatches) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "InitialSequenceBatches")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumInitialSequenceBatches)
				if err := _Polygonvalidium.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialSequenceBatches is a log parse operation binding the contract event 0x060116213bcbf54ca19fd649dc84b59ab2bbd200ab199770e4d923e222a28e7f.
//
// Solidity: event InitialSequenceBatches(bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseInitialSequenceBatches(log types.Log) (*PolygonvalidiumInitialSequenceBatches, error) {
	event := new(PolygonvalidiumInitialSequenceBatches)
	if err := _Polygonvalidium.contract.UnpackLog(event, "InitialSequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Polygonvalidium contract.
type PolygonvalidiumInitializedIterator struct {
	Event *PolygonvalidiumInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumInitialized represents a Initialized event raised by the Polygonvalidium contract.
type PolygonvalidiumInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterInitialized(opts *bind.FilterOpts) (*PolygonvalidiumInitializedIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumInitializedIterator{contract: _Polygonvalidium.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumInitialized) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumInitialized)
				if err := _Polygonvalidium.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseInitialized(log types.Log) (*PolygonvalidiumInitialized, error) {
	event := new(PolygonvalidiumInitialized)
	if err := _Polygonvalidium.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSequenceBatchesIterator is returned from FilterSequenceBatches and is used to iterate over the raw logs and unpacked data for SequenceBatches events raised by the Polygonvalidium contract.
type PolygonvalidiumSequenceBatchesIterator struct {
	Event *PolygonvalidiumSequenceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSequenceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSequenceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSequenceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSequenceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSequenceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSequenceBatches represents a SequenceBatches event raised by the Polygonvalidium contract.
type PolygonvalidiumSequenceBatches struct {
	NumBatch   uint64
	L1InfoRoot [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSequenceBatches is a free log retrieval operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSequenceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumSequenceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSequenceBatchesIterator{contract: _Polygonvalidium.contract, event: "SequenceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceBatches is a free log subscription operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSequenceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSequenceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SequenceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSequenceBatches)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceBatches is a log parse operation binding the contract event 0x3e54d0825ed78523037d00a81759237eb436ce774bd546993ee67a1b67b6e766.
//
// Solidity: event SequenceBatches(uint64 indexed numBatch, bytes32 l1InfoRoot)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSequenceBatches(log types.Log) (*PolygonvalidiumSequenceBatches, error) {
	event := new(PolygonvalidiumSequenceBatches)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SequenceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSequenceForceBatchesIterator is returned from FilterSequenceForceBatches and is used to iterate over the raw logs and unpacked data for SequenceForceBatches events raised by the Polygonvalidium contract.
type PolygonvalidiumSequenceForceBatchesIterator struct {
	Event *PolygonvalidiumSequenceForceBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSequenceForceBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSequenceForceBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSequenceForceBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSequenceForceBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSequenceForceBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSequenceForceBatches represents a SequenceForceBatches event raised by the Polygonvalidium contract.
type PolygonvalidiumSequenceForceBatches struct {
	NumBatch uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSequenceForceBatches is a free log retrieval operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSequenceForceBatches(opts *bind.FilterOpts, numBatch []uint64) (*PolygonvalidiumSequenceForceBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSequenceForceBatchesIterator{contract: _Polygonvalidium.contract, event: "SequenceForceBatches", logs: logs, sub: sub}, nil
}

// WatchSequenceForceBatches is a free log subscription operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSequenceForceBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSequenceForceBatches, numBatch []uint64) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SequenceForceBatches", numBatchRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSequenceForceBatches)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequenceForceBatches is a log parse operation binding the contract event 0x648a61dd2438f072f5a1960939abd30f37aea80d2e94c9792ad142d3e0a490a4.
//
// Solidity: event SequenceForceBatches(uint64 indexed numBatch)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSequenceForceBatches(log types.Log) (*PolygonvalidiumSequenceForceBatches, error) {
	event := new(PolygonvalidiumSequenceForceBatches)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SequenceForceBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSetDataAvailabilityProtocolIterator is returned from FilterSetDataAvailabilityProtocol and is used to iterate over the raw logs and unpacked data for SetDataAvailabilityProtocol events raised by the Polygonvalidium contract.
type PolygonvalidiumSetDataAvailabilityProtocolIterator struct {
	Event *PolygonvalidiumSetDataAvailabilityProtocol // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSetDataAvailabilityProtocolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSetDataAvailabilityProtocol)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSetDataAvailabilityProtocol)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSetDataAvailabilityProtocolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSetDataAvailabilityProtocolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSetDataAvailabilityProtocol represents a SetDataAvailabilityProtocol event raised by the Polygonvalidium contract.
type PolygonvalidiumSetDataAvailabilityProtocol struct {
	NewDataAvailabilityProtocol common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetDataAvailabilityProtocol is a free log retrieval operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSetDataAvailabilityProtocol(opts *bind.FilterOpts) (*PolygonvalidiumSetDataAvailabilityProtocolIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSetDataAvailabilityProtocolIterator{contract: _Polygonvalidium.contract, event: "SetDataAvailabilityProtocol", logs: logs, sub: sub}, nil
}

// WatchSetDataAvailabilityProtocol is a free log subscription operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSetDataAvailabilityProtocol(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSetDataAvailabilityProtocol) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SetDataAvailabilityProtocol")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSetDataAvailabilityProtocol)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetDataAvailabilityProtocol is a log parse operation binding the contract event 0xd331bd4c4cd1afecb94a225184bded161ff3213624ba4fb58c4f30c5a861144a.
//
// Solidity: event SetDataAvailabilityProtocol(address newDataAvailabilityProtocol)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSetDataAvailabilityProtocol(log types.Log) (*PolygonvalidiumSetDataAvailabilityProtocol, error) {
	event := new(PolygonvalidiumSetDataAvailabilityProtocol)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SetDataAvailabilityProtocol", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSetForceBatchAddressIterator is returned from FilterSetForceBatchAddress and is used to iterate over the raw logs and unpacked data for SetForceBatchAddress events raised by the Polygonvalidium contract.
type PolygonvalidiumSetForceBatchAddressIterator struct {
	Event *PolygonvalidiumSetForceBatchAddress // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSetForceBatchAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSetForceBatchAddress)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSetForceBatchAddress)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSetForceBatchAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSetForceBatchAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSetForceBatchAddress represents a SetForceBatchAddress event raised by the Polygonvalidium contract.
type PolygonvalidiumSetForceBatchAddress struct {
	NewForceBatchAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchAddress is a free log retrieval operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSetForceBatchAddress(opts *bind.FilterOpts) (*PolygonvalidiumSetForceBatchAddressIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSetForceBatchAddressIterator{contract: _Polygonvalidium.contract, event: "SetForceBatchAddress", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchAddress is a free log subscription operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSetForceBatchAddress(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSetForceBatchAddress) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SetForceBatchAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSetForceBatchAddress)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForceBatchAddress is a log parse operation binding the contract event 0x5fbd7dd171301c4a1611a84aac4ba86d119478560557755f7927595b082634fb.
//
// Solidity: event SetForceBatchAddress(address newForceBatchAddress)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSetForceBatchAddress(log types.Log) (*PolygonvalidiumSetForceBatchAddress, error) {
	event := new(PolygonvalidiumSetForceBatchAddress)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SetForceBatchAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSetForceBatchTimeoutIterator is returned from FilterSetForceBatchTimeout and is used to iterate over the raw logs and unpacked data for SetForceBatchTimeout events raised by the Polygonvalidium contract.
type PolygonvalidiumSetForceBatchTimeoutIterator struct {
	Event *PolygonvalidiumSetForceBatchTimeout // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSetForceBatchTimeoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSetForceBatchTimeout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSetForceBatchTimeout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSetForceBatchTimeoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSetForceBatchTimeoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSetForceBatchTimeout represents a SetForceBatchTimeout event raised by the Polygonvalidium contract.
type PolygonvalidiumSetForceBatchTimeout struct {
	NewforceBatchTimeout uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetForceBatchTimeout is a free log retrieval operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSetForceBatchTimeout(opts *bind.FilterOpts) (*PolygonvalidiumSetForceBatchTimeoutIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSetForceBatchTimeoutIterator{contract: _Polygonvalidium.contract, event: "SetForceBatchTimeout", logs: logs, sub: sub}, nil
}

// WatchSetForceBatchTimeout is a free log subscription operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSetForceBatchTimeout(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSetForceBatchTimeout) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SetForceBatchTimeout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSetForceBatchTimeout)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetForceBatchTimeout is a log parse operation binding the contract event 0xa7eb6cb8a613eb4e8bddc1ac3d61ec6cf10898760f0b187bcca794c6ca6fa40b.
//
// Solidity: event SetForceBatchTimeout(uint64 newforceBatchTimeout)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSetForceBatchTimeout(log types.Log) (*PolygonvalidiumSetForceBatchTimeout, error) {
	event := new(PolygonvalidiumSetForceBatchTimeout)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SetForceBatchTimeout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSetTrustedSequencerIterator is returned from FilterSetTrustedSequencer and is used to iterate over the raw logs and unpacked data for SetTrustedSequencer events raised by the Polygonvalidium contract.
type PolygonvalidiumSetTrustedSequencerIterator struct {
	Event *PolygonvalidiumSetTrustedSequencer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSetTrustedSequencerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSetTrustedSequencer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSetTrustedSequencer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSetTrustedSequencerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSetTrustedSequencerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSetTrustedSequencer represents a SetTrustedSequencer event raised by the Polygonvalidium contract.
type PolygonvalidiumSetTrustedSequencer struct {
	NewTrustedSequencer common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencer is a free log retrieval operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSetTrustedSequencer(opts *bind.FilterOpts) (*PolygonvalidiumSetTrustedSequencerIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSetTrustedSequencerIterator{contract: _Polygonvalidium.contract, event: "SetTrustedSequencer", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencer is a free log subscription operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSetTrustedSequencer(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSetTrustedSequencer) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SetTrustedSequencer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSetTrustedSequencer)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencer is a log parse operation binding the contract event 0xf54144f9611984021529f814a1cb6a41e22c58351510a0d9f7e822618abb9cc0.
//
// Solidity: event SetTrustedSequencer(address newTrustedSequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSetTrustedSequencer(log types.Log) (*PolygonvalidiumSetTrustedSequencer, error) {
	event := new(PolygonvalidiumSetTrustedSequencer)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SetTrustedSequencer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSetTrustedSequencerURLIterator is returned from FilterSetTrustedSequencerURL and is used to iterate over the raw logs and unpacked data for SetTrustedSequencerURL events raised by the Polygonvalidium contract.
type PolygonvalidiumSetTrustedSequencerURLIterator struct {
	Event *PolygonvalidiumSetTrustedSequencerURL // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSetTrustedSequencerURLIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSetTrustedSequencerURL)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSetTrustedSequencerURL)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSetTrustedSequencerURLIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSetTrustedSequencerURLIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSetTrustedSequencerURL represents a SetTrustedSequencerURL event raised by the Polygonvalidium contract.
type PolygonvalidiumSetTrustedSequencerURL struct {
	NewTrustedSequencerURL string
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedSequencerURL is a free log retrieval operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSetTrustedSequencerURL(opts *bind.FilterOpts) (*PolygonvalidiumSetTrustedSequencerURLIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSetTrustedSequencerURLIterator{contract: _Polygonvalidium.contract, event: "SetTrustedSequencerURL", logs: logs, sub: sub}, nil
}

// WatchSetTrustedSequencerURL is a free log subscription operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSetTrustedSequencerURL(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSetTrustedSequencerURL) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SetTrustedSequencerURL")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSetTrustedSequencerURL)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedSequencerURL is a log parse operation binding the contract event 0x6b8f723a4c7a5335cafae8a598a0aa0301be1387c037dccc085b62add6448b20.
//
// Solidity: event SetTrustedSequencerURL(string newTrustedSequencerURL)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSetTrustedSequencerURL(log types.Log) (*PolygonvalidiumSetTrustedSequencerURL, error) {
	event := new(PolygonvalidiumSetTrustedSequencerURL)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SetTrustedSequencerURL", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator is returned from FilterSwitchSequenceWithDataAvailability and is used to iterate over the raw logs and unpacked data for SwitchSequenceWithDataAvailability events raised by the Polygonvalidium contract.
type PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator struct {
	Event *PolygonvalidiumSwitchSequenceWithDataAvailability // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumSwitchSequenceWithDataAvailability)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumSwitchSequenceWithDataAvailability)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumSwitchSequenceWithDataAvailability represents a SwitchSequenceWithDataAvailability event raised by the Polygonvalidium contract.
type PolygonvalidiumSwitchSequenceWithDataAvailability struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSwitchSequenceWithDataAvailability is a free log retrieval operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterSwitchSequenceWithDataAvailability(opts *bind.FilterOpts) (*PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumSwitchSequenceWithDataAvailabilityIterator{contract: _Polygonvalidium.contract, event: "SwitchSequenceWithDataAvailability", logs: logs, sub: sub}, nil
}

// WatchSwitchSequenceWithDataAvailability is a free log subscription operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchSwitchSequenceWithDataAvailability(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumSwitchSequenceWithDataAvailability) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "SwitchSequenceWithDataAvailability")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumSwitchSequenceWithDataAvailability)
				if err := _Polygonvalidium.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwitchSequenceWithDataAvailability is a log parse operation binding the contract event 0xf32a0473f809a720a4f8af1e50d353f1caf7452030626fdaac4273f5e6587f41.
//
// Solidity: event SwitchSequenceWithDataAvailability()
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseSwitchSequenceWithDataAvailability(log types.Log) (*PolygonvalidiumSwitchSequenceWithDataAvailability, error) {
	event := new(PolygonvalidiumSwitchSequenceWithDataAvailability)
	if err := _Polygonvalidium.contract.UnpackLog(event, "SwitchSequenceWithDataAvailability", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumTransferAdminRoleIterator is returned from FilterTransferAdminRole and is used to iterate over the raw logs and unpacked data for TransferAdminRole events raised by the Polygonvalidium contract.
type PolygonvalidiumTransferAdminRoleIterator struct {
	Event *PolygonvalidiumTransferAdminRole // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumTransferAdminRoleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumTransferAdminRole)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumTransferAdminRole)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumTransferAdminRoleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumTransferAdminRoleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumTransferAdminRole represents a TransferAdminRole event raised by the Polygonvalidium contract.
type PolygonvalidiumTransferAdminRole struct {
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTransferAdminRole is a free log retrieval operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterTransferAdminRole(opts *bind.FilterOpts) (*PolygonvalidiumTransferAdminRoleIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumTransferAdminRoleIterator{contract: _Polygonvalidium.contract, event: "TransferAdminRole", logs: logs, sub: sub}, nil
}

// WatchTransferAdminRole is a free log subscription operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchTransferAdminRole(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumTransferAdminRole) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "TransferAdminRole")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumTransferAdminRole)
				if err := _Polygonvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferAdminRole is a log parse operation binding the contract event 0xa5b56b7906fd0a20e3f35120dd8343db1e12e037a6c90111c7e42885e82a1ce6.
//
// Solidity: event TransferAdminRole(address newPendingAdmin)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseTransferAdminRole(log types.Log) (*PolygonvalidiumTransferAdminRole, error) {
	event := new(PolygonvalidiumTransferAdminRole)
	if err := _Polygonvalidium.contract.UnpackLog(event, "TransferAdminRole", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumUpdateEtrogSequenceIterator is returned from FilterUpdateEtrogSequence and is used to iterate over the raw logs and unpacked data for UpdateEtrogSequence events raised by the Polygonvalidium contract.
type PolygonvalidiumUpdateEtrogSequenceIterator struct {
	Event *PolygonvalidiumUpdateEtrogSequence // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumUpdateEtrogSequenceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumUpdateEtrogSequence)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumUpdateEtrogSequence)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumUpdateEtrogSequenceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumUpdateEtrogSequenceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumUpdateEtrogSequence represents a UpdateEtrogSequence event raised by the Polygonvalidium contract.
type PolygonvalidiumUpdateEtrogSequence struct {
	NumBatch           uint64
	Transactions       []byte
	LastGlobalExitRoot [32]byte
	Sequencer          common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateEtrogSequence is a free log retrieval operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterUpdateEtrogSequence(opts *bind.FilterOpts) (*PolygonvalidiumUpdateEtrogSequenceIterator, error) {

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumUpdateEtrogSequenceIterator{contract: _Polygonvalidium.contract, event: "UpdateEtrogSequence", logs: logs, sub: sub}, nil
}

// WatchUpdateEtrogSequence is a free log subscription operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchUpdateEtrogSequence(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumUpdateEtrogSequence) (event.Subscription, error) {

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "UpdateEtrogSequence")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumUpdateEtrogSequence)
				if err := _Polygonvalidium.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateEtrogSequence is a log parse operation binding the contract event 0xd2c80353fc15ef62c6affc7cd6b7ab5b42c43290c50be3372e55ae552cecd19c.
//
// Solidity: event UpdateEtrogSequence(uint64 numBatch, bytes transactions, bytes32 lastGlobalExitRoot, address sequencer)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseUpdateEtrogSequence(log types.Log) (*PolygonvalidiumUpdateEtrogSequence, error) {
	event := new(PolygonvalidiumUpdateEtrogSequence)
	if err := _Polygonvalidium.contract.UnpackLog(event, "UpdateEtrogSequence", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PolygonvalidiumVerifyBatchesIterator is returned from FilterVerifyBatches and is used to iterate over the raw logs and unpacked data for VerifyBatches events raised by the Polygonvalidium contract.
type PolygonvalidiumVerifyBatchesIterator struct {
	Event *PolygonvalidiumVerifyBatches // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PolygonvalidiumVerifyBatchesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolygonvalidiumVerifyBatches)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PolygonvalidiumVerifyBatches)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PolygonvalidiumVerifyBatchesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolygonvalidiumVerifyBatchesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolygonvalidiumVerifyBatches represents a VerifyBatches event raised by the Polygonvalidium contract.
type PolygonvalidiumVerifyBatches struct {
	NumBatch   uint64
	StateRoot  [32]byte
	Aggregator common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVerifyBatches is a free log retrieval operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidium *PolygonvalidiumFilterer) FilterVerifyBatches(opts *bind.FilterOpts, numBatch []uint64, aggregator []common.Address) (*PolygonvalidiumVerifyBatchesIterator, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidium.contract.FilterLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return &PolygonvalidiumVerifyBatchesIterator{contract: _Polygonvalidium.contract, event: "VerifyBatches", logs: logs, sub: sub}, nil
}

// WatchVerifyBatches is a free log subscription operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidium *PolygonvalidiumFilterer) WatchVerifyBatches(opts *bind.WatchOpts, sink chan<- *PolygonvalidiumVerifyBatches, numBatch []uint64, aggregator []common.Address) (event.Subscription, error) {

	var numBatchRule []interface{}
	for _, numBatchItem := range numBatch {
		numBatchRule = append(numBatchRule, numBatchItem)
	}

	var aggregatorRule []interface{}
	for _, aggregatorItem := range aggregator {
		aggregatorRule = append(aggregatorRule, aggregatorItem)
	}

	logs, sub, err := _Polygonvalidium.contract.WatchLogs(opts, "VerifyBatches", numBatchRule, aggregatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolygonvalidiumVerifyBatches)
				if err := _Polygonvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVerifyBatches is a log parse operation binding the contract event 0x9c72852172521097ba7e1482e6b44b351323df0155f97f4ea18fcec28e1f5966.
//
// Solidity: event VerifyBatches(uint64 indexed numBatch, bytes32 stateRoot, address indexed aggregator)
func (_Polygonvalidium *PolygonvalidiumFilterer) ParseVerifyBatches(log types.Log) (*PolygonvalidiumVerifyBatches, error) {
	event := new(PolygonvalidiumVerifyBatches)
	if err := _Polygonvalidium.contract.UnpackLog(event, "VerifyBatches", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
