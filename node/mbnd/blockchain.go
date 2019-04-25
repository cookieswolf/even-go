// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package mbnd

import (
	"os"
	"path/filepath"
	"time"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/blockchain/indexers"
	"github.com/btcsuite/btcd/database"
	"github.com/btcsuite/btcd/mempool"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

const (
	defaultSigCacheMaxSize       = 100000
	defaultFreeTxRelayLimit      = 15.0
	defaultMaxOrphanTransactions = 100
	defaultMaxOrphanTxSize       = 100000
)

var (
	chainConfig *blockchain.Config

	addrIndex *indexers.AddrIndex
	cfIndex   *indexers.CfIndex
	txIndex   *indexers.TxIndex

	// txMemPool defines the transaction memory pool to interact with.
	txMemPool *mempool.TxPool
)

// initDb opens the block database and returns it.
func initDb() error {

	// The database name is based on the database type.
	dbName := blockDbNamePrefix + "_" + cfg.DbType

	path := []string{
		cfg.DataDir, defaultExternalDirname, "btcd", defaultDataDirname, "testnet", dbName,
	}

	dbPath := filepath.Join(path...)

	logger.Infof("Loading block database from '%s'", dbPath)

	db, err = database.Open(cfg.DbType, dbPath, chainParams.Net)
	if err != nil {
		// Return the error if it's not because the database doesn't exist.
		if dbErr, ok := err.(database.Error); !ok || dbErr.ErrorCode != database.ErrDbDoesNotExist {
			return err
		}

		// Create the db if it does not exist.
		err = os.MkdirAll(cfg.DataDir, 0700)
		if err != nil {
			return err
		}

		db, err = database.Create(cfg.DbType, dbPath, chainParams.Net)
		if err != nil {
			return err
		}
	}

	logger.Info("Block database loaded")

	return nil
}

// initBlockchain create the blockchain instance and returns it.
func initBlockchain() error {

	// defines indexers for the database to interact with.
	addrIndex = indexers.NewAddrIndex(db, chainParams)
	cfIndex = indexers.NewCfIndex(db, chainParams)
	txIndex = indexers.NewTxIndex(db)

	indexes := []indexers.Indexer{addrIndex, cfIndex, txIndex}

	// Init new Blockchain
	chainConfig = &blockchain.Config{
		DB:           db,
		Interrupt:    nil,
		ChainParams:  chainParams,
		Checkpoints:  chainParams.Checkpoints,
		TimeSource:   blockchain.NewMedianTime(),
		SigCache:     txscript.NewSigCache(defaultSigCacheMaxSize),
		IndexManager: indexers.NewManager(db, indexes),
		HashCache:    txscript.NewHashCache(defaultSigCacheMaxSize),
	}

	bchain, err = blockchain.New(chainConfig)
	if bchain != nil {
		txMemPool = initMempool(bchain)
	}

	return err
}

// initMempool create the mempool instance and returns it.
func initMempool(bc *blockchain.BlockChain) *mempool.TxPool {

	// Init the minrelaytxfee
	minRelayTxFee, _ := btcutil.NewAmount(mempool.DefaultMinRelayTxFee.ToBTC())

	// Init Mempool transactions config
	conf := &mempool.Config{
		Policy: mempool.Policy{
			DisableRelayPriority: true,
			AcceptNonStd:         true,
			FreeTxRelayLimit:     defaultFreeTxRelayLimit,
			MaxOrphanTxs:         defaultMaxOrphanTransactions,
			MaxOrphanTxSize:      defaultMaxOrphanTxSize,
			MaxSigOpCostPerTx:    blockchain.MaxBlockSigOpsCost / 4,
			MinRelayTxFee:        minRelayTxFee,
			MaxTxVersion:         2,
		},
		ChainParams:    chainParams,
		FetchUtxoView:  bc.FetchUtxoView,
		BestHeight:     func() int32 { return bc.BestSnapshot().Height },
		MedianTimePast: func() time.Time { return bc.BestSnapshot().MedianTime },
		CalcSequenceLock: func(tx *btcutil.Tx, view *blockchain.UtxoViewpoint) (*blockchain.SequenceLock, error) {
			return bc.CalcSequenceLock(tx, view, true)
		},
		IsDeploymentActive: bc.IsDeploymentActive,
		SigCache:           chainConfig.SigCache,
		HashCache:          chainConfig.HashCache,
		AddrIndex:          addrIndex,
		FeeEstimator: mempool.NewFeeEstimator(
			mempool.DefaultEstimateFeeMaxRollback,
			mempool.DefaultEstimateFeeMinRegisteredBlocks),
	}

	return mempool.New(conf)
}
