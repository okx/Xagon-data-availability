package rpcsyncer

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"time"

	"github.com/0xPolygon/cdk-data-availability/db"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/0xPolygon/cdk-data-availability/types"
)

type RPCSyncer struct {
	l2RPCUrl     string
	maxBatchSize uint64
	intervalTime time.Duration
	db           db.DB

	stop chan struct{}
}

func NewRPCSyncer(l2RPCUrl string, maxBatchSize uint64, intervalTime time.Duration, db db.DB) *RPCSyncer {
	rpcSyncer := &RPCSyncer{
		l2RPCUrl:     l2RPCUrl,
		maxBatchSize: maxBatchSize,
		intervalTime: intervalTime,
		db:           db,
		stop:         make(chan struct{}),
	}

	return rpcSyncer
}

// Start starts the SequencerTracker
func (syncer *RPCSyncer) Start(ctx context.Context) {
	log.Infof("starting rpc syncer")

	start, _ := getStartBlock(syncer.db)

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() != nil && ctx.Err() != context.DeadlineExceeded {
				log.Warnf("context cancelled: %v", ctx.Err())
			}
		default:
			time.Sleep(syncer.intervalTime)
			l2MaxBatch, err := BatchNumber(syncer.l2RPCUrl)
			if err != nil {
				log.Fatal("error getting max batch: %v", err)
			}
			log.Infof("starting from block %v, max block %v", start, l2MaxBatch)
			if start > l2MaxBatch {
				log.Infof("no new blocks to sync")
				time.Sleep(10 * time.Second)
				continue
			}
			to := start + syncer.maxBatchSize - 1
			if to > l2MaxBatch {
				to = l2MaxBatch
			}
			seqBatches, err := BatchesByNumbers(syncer.l2RPCUrl, start, to)
			if err != nil {
				log.Fatal("error getting batch data: %v", err)
			}
			offChainData := []types.OffChainData{}
			for _, seqBatch := range seqBatches {
				key := crypto.Keccak256Hash(seqBatch.BatchL2Data)

				offChainData = append(offChainData, types.OffChainData{
					Key:   key,
					Value: seqBatch.BatchL2Data,
				})
			}
			err = setStoreOffChainData(syncer.db, offChainData)
			if err != nil {
				log.Fatal("error storing off chain data: %v", err)
			}
			log.Infof("stored off chain data for blocks from %v to %v, store size:%v", start, to, len(offChainData))
			if setStartBlock(syncer.db, to) != nil {
				log.Fatal("error setting start block: %v", err)
			}
			start = to + 1
		}
	}
}

// Stop stops the SequencerTracker
func (st *RPCSyncer) Stop() {
	close(st.stop)
}
