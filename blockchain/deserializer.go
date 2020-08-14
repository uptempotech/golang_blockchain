package blockchain

import (
	"bytes"
	"encoding/gob"
)

// Deserialize takes a byte array and returns a block
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	Handle(err)

	return &block
}
