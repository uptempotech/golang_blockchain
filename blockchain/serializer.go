package blockchain

import (
	"bytes"
	"encoding/gob"
)

// Serialize creates a byte array of a block
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()
}
