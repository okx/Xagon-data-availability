package main

import (
	"context"
	"github.com/0xPolygon/cdk-data-availability/config"
	"github.com/0xPolygon/cdk-data-availability/db"
	"github.com/0xPolygon/cdk-data-availability/log"
	"github.com/0xPolygon/cdk-data-availability/rpc"
	"github.com/0xPolygon/cdk-data-availability/rpcsyncer"
	"github.com/0xPolygon/cdk-data-availability/services/sync"
	"github.com/urfave/cli/v2"
)

func startSyncRpc(cliCtx *cli.Context) error {
	// Load config
	c, err := config.Load(cliCtx)
	if err != nil {
		panic(err)
	}
	setupLog(c.Log)

	// Prepare DB
	pg, err := db.InitContext(cliCtx.Context, c.DB)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.RunMigrationsUp(pg); err != nil {
		log.Fatal(err)
	}

	storage := db.New(pg)

	var cancelFuncs []context.CancelFunc

	syncer := rpcsyncer.NewRPCSyncer(c.L2RpcURL, c.MaxBatchSize, c.IntervalTime.Duration, storage)
	go syncer.Start(cliCtx.Context)
	cancelFuncs = append(cancelFuncs, syncer.Stop)
	// Register services
	server := rpc.NewServer(
		c.RPC,
		[]rpc.Service{
			{
				Name:    sync.APISYNC,
				Service: sync.NewSyncEndpoints(storage),
			},
		},
	)

	// Run!
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	waitSignal(cancelFuncs)
	return nil
}
