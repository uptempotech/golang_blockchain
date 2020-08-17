package network

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/uptempotech/golang_blockchain/core"
)

func mineblock(chain *core.BlockChain) {
	log.Println("Beginning Miner routine")
	var txs []*core.Transaction

	for id := range memoryPool {
		fmt.Printf("ts: %x\n", memoryPool[id].ID)
		tx := memoryPool[id]
		if chain.VerifyTransaction(&tx) {
			txs = append(txs, &tx)
		} else {
			log.Printf("ts: %x - failed VerifyTransaction", memoryPool[id].ID)
		}
	}

	if len(txs) == 0 {
		fmt.Println("All Transactions are invalid")
		return
	}

	cbTx := core.CoinbaseTx(mineAddress, "")
	txs = append(txs, cbTx)

	newBlock := chain.MineBlock(txs)
	UTXOSet := core.UTXOSet{Blockchain: chain}
	UTXOSet.Reindex()

	fmt.Println("New Block mined")

	for _, tx := range txs {
		txID := hex.EncodeToString(tx.ID)
		delete(memoryPool, txID)
	}

	for _, node := range KnownNodes {
		if node != nodeAddress {
			SendInv(node, "block", [][]byte{newBlock.Hash})
		}
	}

	if len(memoryPool) > 0 {
		mineblock(chain)
	}
}
