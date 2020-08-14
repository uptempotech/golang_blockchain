package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath      = "./tmp/blocks"
	genesisData = "First Transaction from Genesis"
)

// BlockChain represents the structure of the BlockChain itself.
type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}

// InitBlockChain creates a new BlockChain with a Genesis Block added.
func InitBlockChain(address string) *BlockChain {
	var lastHash []byte

	opts := badger.DefaultOptions("")
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			cbtx := CoinbaseTx(address, genesisData)
			genesis := Genesis(cbtx)
			fmt.Println("Genesis created")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("ln"), genesis.Hash)

			lastHash = genesis.Hash

			return &BlockChain{[]*Block{Genesis()}}
		}
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
		})
		return err
	})

	Handle(err)

	blockchain := BlockChain{lastHash, db}
	return &blockchain

	// err = txn.Set([]byte("1h"), genesis.Hash)

	// lastHash = genesis.Hash

	//
}

// AddBlock will add a new Block to the BlockChain.
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// ContinueBlockChain ...
func ContinueBlockChain(address string) *BlockChain {
	var lastHash []byte

	item, err := txn.Get([]byte("1h"))
	Handle(err)
	lastHash, err = item.Value()

	return err

	chain := BlockChain{lastHash, db}

	return &chain
}
