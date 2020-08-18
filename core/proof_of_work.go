package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const difficulty = 12

// ProofOfWork ...
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof ...
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	pow := &ProofOfWork{b, target}
	return pow
}

// PrepareData ...
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	txAsBytes := []byte{}
	for _, tx := range pow.Block.Transactions {
		txAsBytes = append(txAsBytes, tx.Serialize()...)
	}

	data := bytes.Join(
		[][]byte{
			txAsBytes,
			pow.Block.PrevHash,
			intToBytes(int(pow.Block.Timestamp)),
			intToBytes(pow.Block.Height),
			intToBytes(nonce),
		},
		[]byte{},
	)

	return data
}

// Run ...
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte
	nonce := 0

	for nonce < maxNonce {
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)

		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("\n\n")
	return nonce, hash[:]
}

// Validate ...
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.PrepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	isValid := intHash.Cmp(pow.Target) == -1

	return isValid
}

// ToHex ...
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
