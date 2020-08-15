package wallet

import (
	"github.com/mr-tron/base58"
	"github.com/uptempotech/golang_blockchain/blockchain"
)

// Base58Encode ...
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Base58Decode ...
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	blockchain.Handle(err)

	return decode
}
