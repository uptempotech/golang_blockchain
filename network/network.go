package network

import "github.com/uptempotech/golang_blockchain/core"

const (
	protocol      = "tcp"
	version       = 1
	commandLength = 12
)

var (
	nodeAddress     string
	mineAddress     string
	blocksInTransit = [][]byte{}
	memoryPool      = make(map[string]core.Transaction)

	// KnownNodes ...
	KnownNodes = []string{"localhost:3000"}
)

// Addr ...
type Addr struct {
	AddrList []string
}

// Block struct
type Block struct {
	AddrFrom string
	Block    []byte
}

// GetBlocks ...
type GetBlocks struct {
	AddrFrom string
}

// GetData ...
type GetData struct {
	AddrFrom string
	Type     string
	ID       []byte
}

// Inv ...
type Inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

// Tx ...
type Tx struct {
	AddrFrom    string
	Transaction []byte
}

// Version ...
type Version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}
