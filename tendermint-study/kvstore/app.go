package main

import (
	"bytes"
	"log"

	"github.com/dgraph-io/badger/v3"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

var _ abcitypes.Application = (*KVStoreApplication)(nil)

type KVStoreApplication struct {
	db           *badger.DB
	pendingBlock *badger.Txn
}

func NewKVStoreApplication(db *badger.DB) *KVStoreApplication {
	return &KVStoreApplication{db: db}
}

func (app *KVStoreApplication) Info(req abcitypes.RequestInfo) abcitypes.ResponseInfo {
	return abcitypes.ResponseInfo{}
}

func (app *KVStoreApplication) InitChain(req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	return abcitypes.ResponseInitChain{}
}

func (app *KVStoreApplication) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *KVStoreApplication) ListSnapshots(abcitypes.RequestListSnapshots) abcitypes.ResponseListSnapshots {
	return abcitypes.ResponseListSnapshots{}
}

func (app *KVStoreApplication) OfferSnapshot(abcitypes.RequestOfferSnapshot) abcitypes.ResponseOfferSnapshot {
	return abcitypes.ResponseOfferSnapshot{}
}

func (app *KVStoreApplication) LoadSnapshotChunk(abcitypes.RequestLoadSnapshotChunk) abcitypes.ResponseLoadSnapshotChunk {
	return abcitypes.ResponseLoadSnapshotChunk{}
}

func (app *KVStoreApplication) ApplySnapshotChunk(abcitypes.RequestApplySnapshotChunk) abcitypes.ResponseApplySnapshotChunk {
	return abcitypes.ResponseApplySnapshotChunk{}
}

// ================ CheckTx ================

func (app *KVStoreApplication) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	code := app.validateTx(req.Tx)
	return abcitypes.ResponseCheckTx{Code: code}
}

func (app *KVStoreApplication) validateTx(tx []byte) uint32 {
	parts := bytes.SplitN(tx, []byte("="), 2)

	// check that the transaction is not malformed
	if len(parts) != 2 || len(parts[0]) == 0 {
		return 1
	}
	return 0
}

// ================ BeginBlock ================

func (app *KVStoreApplication) BeginBlock(req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	app.pendingBlock = app.db.NewTransaction(true)
	return abcitypes.ResponseBeginBlock{}
}

// ================ DeliverTx ================
// DeliverTx 내에 커밋이 없는 이유는 블록이 찍히기도 전에 여기서 처리되면 해당 데이터가 처리되기전에 조회 가능해져버림.

func (app *KVStoreApplication) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	if code := app.validateTx(req.Tx); code != 0 {
		return abcitypes.ResponseDeliverTx{Code: code}
	}

	parts := bytes.SplitN(req.Tx, []byte("="), 2)
	key, value := parts[0], parts[1]

	if err := app.pendingBlock.Set(key, value); err != nil {
		log.Panicf("Error reading database, unable to verify tx: %v", err)
	}

	return abcitypes.ResponseDeliverTx{Code: 0}
}

// ================ Commit ================

func (app *KVStoreApplication) Commit() abcitypes.ResponseCommit {
	app.pendingBlock.Commit()
	return abcitypes.ResponseCommit{Data: []byte{}}
}

// ================ Queries ================
// 텐더민트 코어를 프록시로 사용함으로써, light client can leverage the unified API
// what is the light client?
// It describes a watcher of distributed consensus process that only validates the consensus algorithm and not the state machine transactions within.

func (app *KVStoreApplication) Query(req abcitypes.RequestQuery) abcitypes.ResponseQuery {
	resp := abcitypes.ResponseQuery{Key: req.Data}

	dbErr := app.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(req.Data)
		if err != nil {
			if err != badger.ErrKeyNotFound {
				return err
			}
			resp.Log = "key does not exist"
			return nil
		}

		return item.Value(func(val []byte) error {
			resp.Log = "exists"
			resp.Value = val
			return nil
		})
	})
	if dbErr != nil {
		log.Panicf("Error reading database, unable to verify tx: %v", dbErr)
	}
	return resp
}
